// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRefreshQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlockQuota(v string) *DescribeVodRefreshQuotaResponseBody
	GetBlockQuota() *string
	SetDirQuota(v string) *DescribeVodRefreshQuotaResponseBody
	GetDirQuota() *string
	SetDirRemain(v string) *DescribeVodRefreshQuotaResponseBody
	GetDirRemain() *string
	SetPreloadQuota(v string) *DescribeVodRefreshQuotaResponseBody
	GetPreloadQuota() *string
	SetPreloadRemain(v string) *DescribeVodRefreshQuotaResponseBody
	GetPreloadRemain() *string
	SetRequestId(v string) *DescribeVodRefreshQuotaResponseBody
	GetRequestId() *string
	SetUrlQuota(v string) *DescribeVodRefreshQuotaResponseBody
	GetUrlQuota() *string
	SetUrlRemain(v string) *DescribeVodRefreshQuotaResponseBody
	GetUrlRemain() *string
	SetBlockRemain(v string) *DescribeVodRefreshQuotaResponseBody
	GetBlockRemain() *string
}

type DescribeVodRefreshQuotaResponseBody struct {
	// The maximum number of Object Storage Service (OSS) buckets that can be refreshed each day.
	//
	// example:
	//
	// 500
	BlockQuota *string `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	// The maximum number of directories of files that can be refreshed each day.
	//
	// example:
	//
	// 100
	DirQuota *string `json:"DirQuota,omitempty" xml:"DirQuota,omitempty"`
	// The remaining number of directories of files that can be refreshed on the current day.
	//
	// example:
	//
	// 99
	DirRemain *string `json:"DirRemain,omitempty" xml:"DirRemain,omitempty"`
	// The maximum number of URLs of files that can be prefetched each day.
	//
	// example:
	//
	// 500
	PreloadQuota *string `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	// The remaining number of URLs of files that can be prefetched on the current day.
	//
	// example:
	//
	// 500
	PreloadRemain *string `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 42E0554B-80F4-4921-****-ACFB22CAAAD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The maximum number of URLs of files that can be refreshed each day.
	//
	// example:
	//
	// 2000
	UrlQuota *string `json:"UrlQuota,omitempty" xml:"UrlQuota,omitempty"`
	// The remaining number of URLs of files that can be refreshed on the current day.
	//
	// example:
	//
	// 1996
	UrlRemain *string `json:"UrlRemain,omitempty" xml:"UrlRemain,omitempty"`
	// The remaining number of OSS buckets that can be refreshed on the current day.
	//
	// example:
	//
	// 500
	BlockRemain *string `json:"blockRemain,omitempty" xml:"blockRemain,omitempty"`
}

func (s DescribeVodRefreshQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshQuotaResponseBody) GetBlockQuota() *string {
	return s.BlockQuota
}

func (s *DescribeVodRefreshQuotaResponseBody) GetDirQuota() *string {
	return s.DirQuota
}

func (s *DescribeVodRefreshQuotaResponseBody) GetDirRemain() *string {
	return s.DirRemain
}

func (s *DescribeVodRefreshQuotaResponseBody) GetPreloadQuota() *string {
	return s.PreloadQuota
}

func (s *DescribeVodRefreshQuotaResponseBody) GetPreloadRemain() *string {
	return s.PreloadRemain
}

func (s *DescribeVodRefreshQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodRefreshQuotaResponseBody) GetUrlQuota() *string {
	return s.UrlQuota
}

func (s *DescribeVodRefreshQuotaResponseBody) GetUrlRemain() *string {
	return s.UrlRemain
}

func (s *DescribeVodRefreshQuotaResponseBody) GetBlockRemain() *string {
	return s.BlockRemain
}

func (s *DescribeVodRefreshQuotaResponseBody) SetBlockQuota(v string) *DescribeVodRefreshQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetDirQuota(v string) *DescribeVodRefreshQuotaResponseBody {
	s.DirQuota = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetDirRemain(v string) *DescribeVodRefreshQuotaResponseBody {
	s.DirRemain = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetPreloadQuota(v string) *DescribeVodRefreshQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetPreloadRemain(v string) *DescribeVodRefreshQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetRequestId(v string) *DescribeVodRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetUrlQuota(v string) *DescribeVodRefreshQuotaResponseBody {
	s.UrlQuota = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetUrlRemain(v string) *DescribeVodRefreshQuotaResponseBody {
	s.UrlRemain = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetBlockRemain(v string) *DescribeVodRefreshQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
