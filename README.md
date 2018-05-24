# Performance problem when generating random numbers

When running this code on an AWS m4.large, it turns out that generating random uuids is slower without the seccomp=unconfined switch.

Repeating the tests on Docker for mac doesn't give different timings for the two settings.

docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app node:8 ./linux -m 100000

Ran 100000 loops in 1.947512694s

docker run --security-opt seccomp=unconfined --rm -v "$PWD":/usr/src/app -w /usr/src/app node:8 ./linux -m 100000

Ran 100000 loops in 440.173787ms