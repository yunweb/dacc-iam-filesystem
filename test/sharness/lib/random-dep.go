// package randomdep is here to introduce a dependency in random for godep to
// function properly. this way we can keep go-random vendored and not
// accidentally break our tests when we change it.
package randomdep

import (
	_ "github.com/Harold-the-Axeman/dacc-iam-filesystem/Godeps/_workspace/src/github.com/jbenet/go-random"
	_ "github.com/Harold-the-Axeman/dacc-iam-filesystem/Godeps/_workspace/src/github.com/jbenet/go-random-files"
)
