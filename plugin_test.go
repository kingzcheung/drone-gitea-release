package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var changelogStr = `
# Changelog

### [1.4.3] 2020-04-14
#### Fixed:
- fixed bugs

---

### [1.4.2] 2020-04-14
#### Fixed:
- fixed bugs

---

### [1.4.1] 2020-04-14
#### Fixed:
- fixed bugs
#### Changed
- changed files

`

func Test_splitContent(t *testing.T) {
	as := assert.New(t)
	res := `# Changelog

### [1.4.3] 2020-04-14
#### Fixed:
- fixed bugs`
	as.Equal(splitFileContent(changelogStr), res)
}
