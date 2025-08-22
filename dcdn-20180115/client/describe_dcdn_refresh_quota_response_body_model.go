// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlockQuota(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetBlockQuota() *string
	SetBlockRemain(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetBlockRemain() *string
	SetDirQuota(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetDirQuota() *string
	SetDirRemain(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetDirRemain() *string
	SetIgnoreParamsQuota(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetIgnoreParamsQuota() *string
	SetIgnoreParamsRemain(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetIgnoreParamsRemain() *string
	SetPreloadQuota(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetPreloadQuota() *string
	SetPreloadRemain(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetPreloadRemain() *string
	SetRegexQuota(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetRegexQuota() *string
	SetRegexRemain(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetRegexRemain() *string
	SetRequestId(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetRequestId() *string
	SetUrlQuota(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetUrlQuota() *string
	SetUrlRemain(v string) *DescribeDcdnRefreshQuotaResponseBody
	GetUrlRemain() *string
}

type DescribeDcdnRefreshQuotaResponseBody struct {
	// The maximum number of URLs that can be blocked.
	//
	// example:
	//
	// 100
	BlockQuota *string `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	// The remaining number of URLs that can be blocked on the current day.
	//
	// example:
	//
	// 100
	BlockRemain *string `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	// The maximum number of directories that can be refreshed on the current day.
	//
	// example:
	//
	// 100
	DirQuota *string `json:"DirQuota,omitempty" xml:"DirQuota,omitempty"`
	// The remaining number of directories that can be refreshed on the current day.
	//
	// example:
	//
	// 100
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
	// The maximum number of URLs that can be prefetched on the current day.
	//
	// example:
	//
	// 500
	PreloadQuota *string `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	// The remaining number of URLs that can be prefetched on the current day.
	//
	// example:
	//
	// 500
	PreloadRemain *string `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	// The maximum number of URLs or directories that can be refreshed by using regular expressions on the current day.
	//
	// example:
	//
	// 100
	RegexQuota *string `json:"RegexQuota,omitempty" xml:"RegexQuota,omitempty"`
	// The remaining number of URLs or directories that can be refreshed by using regular expressions on the current day.
	//
	// example:
	//
	// 100
	RegexRemain *string `json:"RegexRemain,omitempty" xml:"RegexRemain,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 42E0554B-80F4-4921-AED6-ACFB22CAAAD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The maximum number of URLs that can be refreshed on the current day.
	//
	// example:
	//
	// 2000
	UrlQuota *string `json:"UrlQuota,omitempty" xml:"UrlQuota,omitempty"`
	// The remaining number of URLs that can be refreshed on the current day.
	//
	// example:
	//
	// 2000
	UrlRemain *string `json:"UrlRemain,omitempty" xml:"UrlRemain,omitempty"`
}

func (s DescribeDcdnRefreshQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetBlockQuota() *string {
	return s.BlockQuota
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetBlockRemain() *string {
	return s.BlockRemain
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetDirQuota() *string {
	return s.DirQuota
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetDirRemain() *string {
	return s.DirRemain
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetIgnoreParamsQuota() *string {
	return s.IgnoreParamsQuota
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetIgnoreParamsRemain() *string {
	return s.IgnoreParamsRemain
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetPreloadQuota() *string {
	return s.PreloadQuota
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetPreloadRemain() *string {
	return s.PreloadRemain
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetRegexQuota() *string {
	return s.RegexQuota
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetRegexRemain() *string {
	return s.RegexRemain
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetUrlQuota() *string {
	return s.UrlQuota
}

func (s *DescribeDcdnRefreshQuotaResponseBody) GetUrlRemain() *string {
	return s.UrlRemain
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetBlockQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetBlockRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetDirQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.DirQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetDirRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.DirRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetIgnoreParamsQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.IgnoreParamsQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetIgnoreParamsRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.IgnoreParamsRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetPreloadQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetPreloadRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetRegexQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.RegexQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetRegexRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.RegexRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetRequestId(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetUrlQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.UrlQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetUrlRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.UrlRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
