// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlockQuota(v int32) *DescribeVodUserQuotaResponseBody
	GetBlockQuota() *int32
	SetBlockRemain(v int32) *DescribeVodUserQuotaResponseBody
	GetBlockRemain() *int32
	SetDomainQuota(v int32) *DescribeVodUserQuotaResponseBody
	GetDomainQuota() *int32
	SetPreloadQuota(v int32) *DescribeVodUserQuotaResponseBody
	GetPreloadQuota() *int32
	SetPreloadRemain(v int32) *DescribeVodUserQuotaResponseBody
	GetPreloadRemain() *int32
	SetRefreshDirQuota(v int32) *DescribeVodUserQuotaResponseBody
	GetRefreshDirQuota() *int32
	SetRefreshDirRemain(v int32) *DescribeVodUserQuotaResponseBody
	GetRefreshDirRemain() *int32
	SetRefreshUrlQuota(v int32) *DescribeVodUserQuotaResponseBody
	GetRefreshUrlQuota() *int32
	SetRefreshUrlRemain(v int32) *DescribeVodUserQuotaResponseBody
	GetRefreshUrlRemain() *int32
	SetRequestId(v string) *DescribeVodUserQuotaResponseBody
	GetRequestId() *string
}

type DescribeVodUserQuotaResponseBody struct {
	BlockQuota       *int32  `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	BlockRemain      *int32  `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	DomainQuota      *int32  `json:"DomainQuota,omitempty" xml:"DomainQuota,omitempty"`
	PreloadQuota     *int32  `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	PreloadRemain    *int32  `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	RefreshDirQuota  *int32  `json:"RefreshDirQuota,omitempty" xml:"RefreshDirQuota,omitempty"`
	RefreshDirRemain *int32  `json:"RefreshDirRemain,omitempty" xml:"RefreshDirRemain,omitempty"`
	RefreshUrlQuota  *int32  `json:"RefreshUrlQuota,omitempty" xml:"RefreshUrlQuota,omitempty"`
	RefreshUrlRemain *int32  `json:"RefreshUrlRemain,omitempty" xml:"RefreshUrlRemain,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodUserQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserQuotaResponseBody) GetBlockQuota() *int32 {
	return s.BlockQuota
}

func (s *DescribeVodUserQuotaResponseBody) GetBlockRemain() *int32 {
	return s.BlockRemain
}

func (s *DescribeVodUserQuotaResponseBody) GetDomainQuota() *int32 {
	return s.DomainQuota
}

func (s *DescribeVodUserQuotaResponseBody) GetPreloadQuota() *int32 {
	return s.PreloadQuota
}

func (s *DescribeVodUserQuotaResponseBody) GetPreloadRemain() *int32 {
	return s.PreloadRemain
}

func (s *DescribeVodUserQuotaResponseBody) GetRefreshDirQuota() *int32 {
	return s.RefreshDirQuota
}

func (s *DescribeVodUserQuotaResponseBody) GetRefreshDirRemain() *int32 {
	return s.RefreshDirRemain
}

func (s *DescribeVodUserQuotaResponseBody) GetRefreshUrlQuota() *int32 {
	return s.RefreshUrlQuota
}

func (s *DescribeVodUserQuotaResponseBody) GetRefreshUrlRemain() *int32 {
	return s.RefreshUrlRemain
}

func (s *DescribeVodUserQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodUserQuotaResponseBody) SetBlockQuota(v int32) *DescribeVodUserQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetBlockRemain(v int32) *DescribeVodUserQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetDomainQuota(v int32) *DescribeVodUserQuotaResponseBody {
	s.DomainQuota = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetPreloadQuota(v int32) *DescribeVodUserQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetPreloadRemain(v int32) *DescribeVodUserQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetRefreshDirQuota(v int32) *DescribeVodUserQuotaResponseBody {
	s.RefreshDirQuota = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetRefreshDirRemain(v int32) *DescribeVodUserQuotaResponseBody {
	s.RefreshDirRemain = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetRefreshUrlQuota(v int32) *DescribeVodUserQuotaResponseBody {
	s.RefreshUrlQuota = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetRefreshUrlRemain(v int32) *DescribeVodUserQuotaResponseBody {
	s.RefreshUrlRemain = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) SetRequestId(v string) *DescribeVodUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
