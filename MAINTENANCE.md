# Maintenance

## Background

Maintained fork: `0xble/tailscale` of `tailscale/tailscale`; maintained and
upstream branch `main`. This temporary isolated checkout is remote-only, not a
runtime canonical checkout. Accepted baseline:
`48e0334aaca92682f1ec59962de93afd21c49ac8`. Publish only to `origin`; never
push upstream.

## Preserve

- Mac GUI builds retain the intended `drive` command behavior.
- Fork-branded version output is source/build behavior; installation and running
the Homebrew package are separate and not evidence of fork consumption.

## Active patches

### TAILSCALE-001: `cmd/tailscale/cli,version: enable drive in ts_mac_gui and brand version output`

- **Status:** Active; `df10e2df`, following superseded experiment/revert `cb10f95f`, `1ff1375b`.
- **Behavior:** `drive` is available in Mac GUI builds and version output identifies the fork.
- **Surfaces:** `cmd/tailscale/cli/drive.go`, `version/print.go`.
- **Upstream issue:** None after checked 2026-09-09.
- **Upstream PR:** None after checked 2026-09-09.
- **Regression:** `./tool/go test ./cmd/tailscale/cli ./version` and `./tool/go build ./cmd/tailscale` pass.
- **Rollback:** revert `df10e2df`; do not reapply the superseded experiment.
- **Retire when:** released upstream supports the Mac-GUI behavior and neutral version identity is authorized and verified.

### TAILSCALE-002: `chore: bump fork version to 0xble.1.1.0`

- **Status:** Active; `eb7e538c`.
- **Behavior:** fork build version remains distinct across upstream merges.
- **Surfaces:** `version/print.go`.
- **Upstream issue:** None after checked 2026-09-09.
- **Upstream PR:** None after checked 2026-09-09.
- **Regression:** `./tool/go test ./version` passes and build emits the expected fork identity.
- **Rollback:** revert `eb7e538c` only after selecting the replacement version policy.
- **Retire when:** an approved unbranded build replaces the fork.

## Update

Every maintenance run fetches `origin` and latest `upstream/main`, reconciles
`main`, preserves only these active records, and runs declared regression, test,
and build proof before authorized publication. Immediately before `Updated` or
`Already current`, fetch upstream again and prove no upstream-only commits;
otherwise report `Blocked` with stage, refs, and evidence. Update this contract
with each patch addition, change, or retirement; missing coverage blocks
publication.

## Verify

```text
./tool/go test ./cmd/tailscale/cli ./version
./tool/go build ./cmd/tailscale
git rev-list --left-right --count upstream/main...main
```

Require a fresh final fetch with zero upstream-only commits and local/`origin`
SHA parity after authorized publication. Installation and runtime SHA proof are
only required in separately authorized stages.