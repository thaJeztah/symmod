// Package symmod is a module to test how go modules handles symlinked
// go.mod files on a repository which uses a vendor.mod.
//
// We use a "vendor.mod" on the Moby repository (https://github.com/moby/moby),
// because the project is at > v2.0.0, but does not yet have versioned
// import paths, so uses "+incompatible" versions.
//
// While go modules allows for "+incompatible" versions to exist on repositories
// that are not a module (so don't have a "go.mod"), it does NOT allow this
// for repositories that DO have a go.mod (adding a "go.mod" means that the
// project is required to use SemVer AND to use versioned import paths for
// versions > 1).
//
// This repository is created to test what happens if the modules uses a
// symlink for the go.mod. If go modules ignores symlinks, this means that
// the repository can be used as a "module" when running locally, but act
// as a "non-module" (thus allowing "+incompatible") when used as a module
// in other projects.
package symmod

// HelloWorld is a message to send to the world.
const HelloWorld = "hello world"

// Version the module version.
const Version = "1.0.1"
