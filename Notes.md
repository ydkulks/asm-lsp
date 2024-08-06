# asm-lsp Notes

## Neovim side
In `.config/after/plugin`, create a lua script for running the lsp:
```lua
local client = vim.lsp.start_client {
  name = "asm-lsp",
  cmd = {"path/to/go executable"},
}

if not client then
  vim.notify "Client side error"
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "asm",
  callback = function ()
    vim.lsp.buf_attach_client(0,client)
  end
})
```

## Things to learn
- [ ] LSP Specification: [docs](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/)
- [ ] Writing tests in golang: [docs](https://pkg.go.dev/testing)

- [ ] Testing: [docs](https://pkg.go.dev/testing)
    - Test file has to have `_test.go` in it's file name
    - Here is how to run all the tests in the current directory and its sub directories
```sh
go test ./...
```
- [ ] Logging: [docs](https://pkg.go.dev/log)
