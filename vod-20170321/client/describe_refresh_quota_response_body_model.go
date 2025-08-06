// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshCacheQuota(v *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) *DescribeRefreshQuotaResponseBody
	GetRefreshCacheQuota() *DescribeRefreshQuotaResponseBodyRefreshCacheQuota
	SetRequestId(v string) *DescribeRefreshQuotaResponseBody
	GetRequestId() *string
}

type DescribeRefreshQuotaResponseBody struct {
	RefreshCacheQuota *DescribeRefreshQuotaResponseBodyRefreshCacheQuota `json:"RefreshCacheQuota,omitempty" xml:"RefreshCacheQuota,omitempty" type:"Struct"`
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRefreshQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaResponseBody) GetRefreshCacheQuota() *DescribeRefreshQuotaResponseBodyRefreshCacheQuota {
	return s.RefreshCacheQuota
}

func (s *DescribeRefreshQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRefreshQuotaResponseBody) SetRefreshCacheQuota(v *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) *DescribeRefreshQuotaResponseBody {
	s.RefreshCacheQuota = v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetRequestId(v string) *DescribeRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRefreshQuotaResponseBodyRefreshCacheQuota struct {
	DirQuota  *string `json:"DirQuota,omitempty" xml:"DirQuota,omitempty"`
	DirRemain *string `json:"DirRemain,omitempty" xml:"DirRemain,omitempty"`
	UrlQuota  *string `json:"UrlQuota,omitempty" xml:"UrlQuota,omitempty"`
	UrlRemain *string `json:"UrlRemain,omitempty" xml:"UrlRemain,omitempty"`
}

func (s DescribeRefreshQuotaResponseBodyRefreshCacheQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshQuotaResponseBodyRefreshCacheQuota) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) GetDirQuota() *string {
	return s.DirQuota
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) GetDirRemain() *string {
	return s.DirRemain
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) GetUrlQuota() *string {
	return s.UrlQuota
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) GetUrlRemain() *string {
	return s.UrlRemain
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) SetDirQuota(v string) *DescribeRefreshQuotaResponseBodyRefreshCacheQuota {
	s.DirQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) SetDirRemain(v string) *DescribeRefreshQuotaResponseBodyRefreshCacheQuota {
	s.DirRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) SetUrlQuota(v string) *DescribeRefreshQuotaResponseBodyRefreshCacheQuota {
	s.UrlQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) SetUrlRemain(v string) *DescribeRefreshQuotaResponseBodyRefreshCacheQuota {
	s.UrlRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBodyRefreshCacheQuota) Validate() error {
	return dara.Validate(s)
}
