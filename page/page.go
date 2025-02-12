package page

// ReloadPage 重载 page 分页参数
// 返回的分别为 page、pageSize
func ReloadPage(page, pageSize int64) (int64, int64) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = minPageSize
	}

	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	return page, pageSize
}
