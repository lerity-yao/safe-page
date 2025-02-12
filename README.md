# safe-page

对分页数据进行标准化，防止滥用

``` go
func (l *GSiteLogic) GSite(req *types.GSiteOneReq) (resp *types.GSiteOneRes, err error) {
    ......
    page,pageSzie := ReloadPage(req.page, req.pageSize)
    ......
}
```