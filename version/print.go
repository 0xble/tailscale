// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package version

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

const forkVersionSuffix = "-0xble.1.1.0"

func withForkVersionSuffix(v string) string {
	if strings.HasSuffix(v, forkVersionSuffix) {
		return v
	}
	return v + forkVersionSuffix
}

var stringLazy = sync.OnceValue(func() string {
	var ret strings.Builder
	ret.WriteString(withForkVersionSuffix(Short()))
	ret.WriteByte('\n')
	if IsUnstableBuild() {
		fmt.Fprintf(&ret, "  track: unstable (dev); frequent updates and bugs are likely\n")
	}
	if gitCommit() != "" {
		fmt.Fprintf(&ret, "  tailscale commit: %s%s\n", gitCommit(), dirtyString())
	}
	fmt.Fprintf(&ret, "  long version: %s\n", withForkVersionSuffix(Long()))
	if extraGitCommitStamp != "" {
		fmt.Fprintf(&ret, "  other commit: %s\n", extraGitCommitStamp)
	}
	fmt.Fprintf(&ret, "  go version: %s\n", runtime.Version())
	return strings.TrimSpace(ret.String())
})

func String() string {
	return stringLazy()
}
