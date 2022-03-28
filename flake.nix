{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/31ffc50c";
    flake-utils.url = "github:numtide/flake-utils/7e5bf3925";
  };

  outputs = { self, nixpkgs, flake-utils }:
  flake-utils.lib.eachDefaultSystem (system:
  let
    pkgs = nixpkgs.legacyPackages.${system};
  in
  rec {
    packages = flake-utils.lib.flattenTree {
      secret =  let lib = pkgs.lib; in
      pkgs.buildGoModule rec {
        pname = "secret";
        name = "secret";
        version = "development";

        vendorSha256 = "sha256-H4xw+8vvCxQG/ZQBSxE62O1wSpICpeXqjXPzNBzZU1Y=";

        src = ./.;

      };
    };

    defaultPackage = packages.secret;
    defaultApp = packages.secret;


  }
  );
}
