// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlockQuota(v string) *DescribeRefreshQuotaResponseBody
	GetBlockQuota() *string
	SetBlockRemain(v string) *DescribeRefreshQuotaResponseBody
	GetBlockRemain() *string
	SetDirQuota(v string) *DescribeRefreshQuotaResponseBody
	GetDirQuota() *string
	SetDirRemain(v string) *DescribeRefreshQuotaResponseBody
	GetDirRemain() *string
	SetIgnoreParamsQuota(v string) *DescribeRefreshQuotaResponseBody
	GetIgnoreParamsQuota() *string
	SetIgnoreParamsRemain(v string) *DescribeRefreshQuotaResponseBody
	GetIgnoreParamsRemain() *string
	SetPreloadEdgeQuota(v string) *DescribeRefreshQuotaResponseBody
	GetPreloadEdgeQuota() *string
	SetPreloadEdgeRemain(v string) *DescribeRefreshQuotaResponseBody
	GetPreloadEdgeRemain() *string
	SetPreloadQuota(v string) *DescribeRefreshQuotaResponseBody
	GetPreloadQuota() *string
	SetPreloadRemain(v string) *DescribeRefreshQuotaResponseBody
	GetPreloadRemain() *string
	SetRegexQuota(v string) *DescribeRefreshQuotaResponseBody
	GetRegexQuota() *string
	SetRegexRemain(v string) *DescribeRefreshQuotaResponseBody
	GetRegexRemain() *string
	SetRequestId(v string) *DescribeRefreshQuotaResponseBody
	GetRequestId() *string
	SetUrlQuota(v string) *DescribeRefreshQuotaResponseBody
	GetUrlQuota() *string
	SetUrlRemain(v string) *DescribeRefreshQuotaResponseBody
	GetUrlRemain() *string
}

type DescribeRefreshQuotaResponseBody struct {
	// The maximum number of URLs that can be refreshed on the current day.
	//
	// example:
	//
	// 300
	BlockQuota *string `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	// The remaining number of times that you can prefetch content to L2 points of presence (POPs) on the current day.
	//
	// example:
	//
	// 100
	BlockRemain *string `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 100
	DirQuota *string `json:"DirQuota,omitempty" xml:"DirQuota,omitempty"`
	// The remaining number of URLs that can be refreshed on the current day.
	//
	// example:
	//
	// 99
	DirRemain *string `json:"DirRemain,omitempty" xml:"DirRemain,omitempty"`
	// The maximum number of URLs or directories with parameters ignored that can be refreshed on the current day.
	//
	// example:
	//
	// 100
	IgnoreParamsQuota *string `json:"IgnoreParamsQuota,omitempty" xml:"IgnoreParamsQuota,omitempty"`
	// The number of remaining URLs or directories that can be refreshed with parameters ignored on the current day.
	//
	// example:
	//
	// 10
	IgnoreParamsRemain *string `json:"IgnoreParamsRemain,omitempty" xml:"IgnoreParamsRemain,omitempty"`
	// The maximum number of directories that can be refreshed on the current day.
	//
	// example:
	//
	// 20
	PreloadEdgeQuota *string `json:"PreloadEdgeQuota,omitempty" xml:"PreloadEdgeQuota,omitempty"`
	// The maximum number of times that you can prefetch content to L1 POPs on the current day.
	//
	// example:
	//
	// 20
	PreloadEdgeRemain *string `json:"PreloadEdgeRemain,omitempty" xml:"PreloadEdgeRemain,omitempty"`
	// The remaining number of times that you can prefetch content to L1 POPs on the current day.
	//
	// example:
	//
	// 500
	PreloadQuota *string `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	// The maximum number of times that you can prefetch content to L1 nodes on the current day.
	//
	// example:
	//
	// 400
	PreloadRemain *string `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	// The maximum number of times that you can prefetch content to L2 POPs on the current day.
	//
	// example:
	//
	// 20
	RegexQuota *string `json:"RegexQuota,omitempty" xml:"RegexQuota,omitempty"`
	// The remaining number of URLs that can be blocked on the current day.
	//
	// example:
	//
	// 10
	RegexRemain *string `json:"RegexRemain,omitempty" xml:"RegexRemain,omitempty"`
	// The maximum number of URLs and directories that can be blocked on the current day.
	//
	// example:
	//
	// 42E0554B-80F4-4921-AED6-ACFB22CAAAD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The remaining number of directories that can be refreshed on the current day.
	//
	// example:
	//
	// 2000
	UrlQuota *string `json:"UrlQuota,omitempty" xml:"UrlQuota,omitempty"`
	// The remaining number of URLs or directories that can be refreshed by using regular expressions on the current day.
	//
	// example:
	//
	// 1996
	UrlRemain *string `json:"UrlRemain,omitempty" xml:"UrlRemain,omitempty"`
}

func (s DescribeRefreshQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaResponseBody) GetBlockQuota() *string {
	return s.BlockQuota
}

func (s *DescribeRefreshQuotaResponseBody) GetBlockRemain() *string {
	return s.BlockRemain
}

func (s *DescribeRefreshQuotaResponseBody) GetDirQuota() *string {
	return s.DirQuota
}

func (s *DescribeRefreshQuotaResponseBody) GetDirRemain() *string {
	return s.DirRemain
}

func (s *DescribeRefreshQuotaResponseBody) GetIgnoreParamsQuota() *string {
	return s.IgnoreParamsQuota
}

func (s *DescribeRefreshQuotaResponseBody) GetIgnoreParamsRemain() *string {
	return s.IgnoreParamsRemain
}

func (s *DescribeRefreshQuotaResponseBody) GetPreloadEdgeQuota() *string {
	return s.PreloadEdgeQuota
}

func (s *DescribeRefreshQuotaResponseBody) GetPreloadEdgeRemain() *string {
	return s.PreloadEdgeRemain
}

func (s *DescribeRefreshQuotaResponseBody) GetPreloadQuota() *string {
	return s.PreloadQuota
}

func (s *DescribeRefreshQuotaResponseBody) GetPreloadRemain() *string {
	return s.PreloadRemain
}

func (s *DescribeRefreshQuotaResponseBody) GetRegexQuota() *string {
	return s.RegexQuota
}

func (s *DescribeRefreshQuotaResponseBody) GetRegexRemain() *string {
	return s.RegexRemain
}

func (s *DescribeRefreshQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRefreshQuotaResponseBody) GetUrlQuota() *string {
	return s.UrlQuota
}

func (s *DescribeRefreshQuotaResponseBody) GetUrlRemain() *string {
	return s.UrlRemain
}

func (s *DescribeRefreshQuotaResponseBody) SetBlockQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetBlockRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetDirQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.DirQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetDirRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.DirRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetIgnoreParamsQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.IgnoreParamsQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetIgnoreParamsRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.IgnoreParamsRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadEdgeQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadEdgeQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadEdgeRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadEdgeRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetRegexQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.RegexQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetRegexRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.RegexRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetRequestId(v string) *DescribeRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetUrlQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.UrlQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetUrlRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.UrlRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
