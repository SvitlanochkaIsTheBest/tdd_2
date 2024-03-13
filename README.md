# Друга лаба ПКПЗ

## How to use

```shell
go test
```

```shell
# For more verbose output
go test -v
```

## What to expect

Rules are mostly the same.

`TestToArabic_TwentyOne` should fail because of the incorrect expected value.

`TestToArabic_Forty` should fail because the number is out of the range (>3999)

Also the `TestToArabic_TwentyTwo` should fail as a test with an invalid Roman number

All other tests should show a `PASS` state.
