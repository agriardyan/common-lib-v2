package commonreq

import (
	"regexp"
	"strconv"

	"github.com/emicklei/go-restful"
)

const (
	PathParamNamespace  = "namespace"
	PathParamUserID     = "userId"
	QueryParamLimit     = "limit"
	QueryParamSort      = "sort"
	QueryParamPage      = "page"
	DefaultLimitPerPage = 30
	DefaultPage         = 0
)

// allow only alphanumeric, hypen, underscore
// https://stackoverflow.com/a/26368676
var regexSortAllowedChars, _ = regexp.Compile("^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$")

//ExtractPagingSortingRequest extract limit and sort query param and convert to int
func ExtractPagingSortingRequest(req *restful.Request) PagingSortingRequest {

	limitRaw := req.QueryParameter(QueryParamLimit)
	sortRaw := req.QueryParameter(QueryParamSort)
	pageRaw := req.QueryParameter(QueryParamPage)

	limit, err := strconv.ParseInt(limitRaw, 10, 64)
	if err != nil {
		limit = DefaultLimitPerPage
	}
	isSortParamClean := regexSortAllowedChars.MatchString(sortRaw)
	sort := sortRaw
	if !isSortParamClean {
		sort = ""
	}
	page, err := strconv.ParseInt(pageRaw, 10, 64)
	if err != nil {
		page = DefaultPage
	}
	return PagingSortingRequest{
		Limit: limit,
		Sort:  sort,
		Page:  page,
	}
}

//PagingSortingRequest dto for paging-sorting request
type PagingSortingRequest struct {
	Limit int64
	Sort  string
	Page  int64
}

type DecodedPagingSortingRequest struct {
	Limit      int64
	Sort       string
	IsSortDesc bool
	Page       int64
}
