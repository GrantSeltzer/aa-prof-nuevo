# aa-prof-nuevo

This project is 2 parts:

- JSON specification for apparmor profiles
- Translator from JSON spec to traditional apparmor profiles

### Why?

The [OCI runtime-spec](https://github.com/opencontainers/runtime-spec) does not currently have an embedded apparmor configuration even though having one would make writing profiles much more accesible. In the same way that container runtimes like [docker](https://github.com/moby/moby) or [runc](https://github.com/opencontainers/runc) can [parse seccomp JSON configurations](https://github.com/opencontainers/runc/blob/master/libcontainer/seccomp/seccomp_linux.go), it would be great to have them understand this apparmor spec as well. 

### Goals
1) Have the specification be fully expressive of apparmor configurations
2) Get the specification merged into the OCI runtime-spec
3) Contribute the parsing of the spec to container runtimes