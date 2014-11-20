# Slimbox

Busybox-like project; a pared-down version of the GNU tools in one big binary.
Since golang is so good at making tiny binaries, I decided to call it slimbox.

Really trying to do this TDD - and that has resulted in a couple rewrites as I
learn better ways of doing the problem and TDD in its domain.

## Implemented

Right now, "pared-down" means "only got `cat`", working through some low-hanging
fruit first, like `wc` and some of the other GNU text-utils.
