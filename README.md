url64 is a small drop-in replacement for the base64 utility that handles
URL-safe base64. (With - and _ in place of + and /, and with no padding).

### Installation

```
go get github.com/superhuman/url64
```

### Usage

To encode:

```
cat file | url64
```

To decode:

```
cat file | url64 -d
```

### Meta-fu

Bug-reports and pull-requests welcome. Everything is available under the MIT
license (see LICENSE.MIT for details)
