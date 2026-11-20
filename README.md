# [Personal Website](https://mrtspeedy.dev)

This repository holds the files for my personal website, including the generated static files in the [public](/public/) directory.
The site can be accessed [here](https://mrtspeedy.dev).

*There is more to come...*

## Structure

```text
kebabspeedy/
├── cmd/
│   └── kebabspeedy/
│       └── main.go
├── internal/
│   ├── assets/
│   │   └── css/
│   │       └── input.css
│   └── views/
│       ├── components/
│       │   └── icons.templ
│       ├── layouts/
│       │   └── layout.templ
│       └── pages/
│           ├── 404.templ
│           ├── about.templ
│           └── home.templ
├── public/
│   ├── about/
│   └── favicon.ico
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

The `main.go` simply renders the pages into html which is in the `public` directory. This process is simplified by using:
```bash
make build
```
This is defined in the `Makefile`, where the templ components are first generated; then tailwind generates an output css from the input, found at `internal/assets/css`; and finally, `main.go` is ran to generate the html files from the components.

**Why the name "kebabspeedy"?**

Well... uhm... I just... like the name..? It's funny, I think...
