// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlockQuota(v int32) *DescribeCdnUserQuotaResponseBody
	GetBlockQuota() *int32
	SetBlockRemain(v int32) *DescribeCdnUserQuotaResponseBody
	GetBlockRemain() *int32
	SetDomainQuota(v int32) *DescribeCdnUserQuotaResponseBody
	GetDomainQuota() *int32
	SetIgnoreParamsQuota(v int32) *DescribeCdnUserQuotaResponseBody
	GetIgnoreParamsQuota() *int32
	SetIgnoreParamsRemain(v int32) *DescribeCdnUserQuotaResponseBody
	GetIgnoreParamsRemain() *int32
	SetPreloadQuota(v int32) *DescribeCdnUserQuotaResponseBody
	GetPreloadQuota() *int32
	SetPreloadRemain(v int32) *DescribeCdnUserQuotaResponseBody
	GetPreloadRemain() *int32
	SetRefreshDirQuota(v int32) *DescribeCdnUserQuotaResponseBody
	GetRefreshDirQuota() *int32
	SetRefreshDirRemain(v int32) *DescribeCdnUserQuotaResponseBody
	GetRefreshDirRemain() *int32
	SetRefreshUrlQuota(v int32) *DescribeCdnUserQuotaResponseBody
	GetRefreshUrlQuota() *int32
	SetRefreshUrlRemain(v int32) *DescribeCdnUserQuotaResponseBody
	GetRefreshUrlRemain() *int32
	SetRequestId(v string) *DescribeCdnUserQuotaResponseBody
	GetRequestId() *string
}

type DescribeCdnUserQuotaResponseBody struct {
	// The maximum number of URLs and directories that can be blocked.
	//
	// example:
	//
	// 100
	BlockQuota *int32 `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	// The remaining number of URLs and directories that can be blocked.
	//
	// example:
	//
	// 100
	BlockRemain *int32 `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	// The maximum number of accelerated domain names.
	//
	// example:
	//
	// 50
	DomainQuota *int32 `json:"DomainQuota,omitempty" xml:"DomainQuota,omitempty"`
	// The maximum number of ignore params that can be refreshed.
	//
	// example:
	//
	// 100
	IgnoreParamsQuota *int32 `json:"IgnoreParamsQuota,omitempty" xml:"IgnoreParamsQuota,omitempty"`
	// The remaining number of ignore params that can be refreshed.
	//
	// example:
	//
	// 10
	IgnoreParamsRemain *int32 `json:"IgnoreParamsRemain,omitempty" xml:"IgnoreParamsRemain,omitempty"`
	// The maximum number of URLs that can be prefetched.
	//
	// example:
	//
	// 500
	PreloadQuota *int32 `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	// The remaining number of URLs that can be prefetched.
	//
	// example:
	//
	// 100
	PreloadRemain *int32 `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	// The maximum number of directories that can be refreshed.
	//
	// example:
	//
	// 100
	RefreshDirQuota *int32 `json:"RefreshDirQuota,omitempty" xml:"RefreshDirQuota,omitempty"`
	// The remaining number of directories that can be refreshed.
	//
	// example:
	//
	// 500
	RefreshDirRemain *int32 `json:"RefreshDirRemain,omitempty" xml:"RefreshDirRemain,omitempty"`
	// The maximum number of URLs that can be refreshed.
	//
	// example:
	//
	// 2000
	RefreshUrlQuota *int32 `json:"RefreshUrlQuota,omitempty" xml:"RefreshUrlQuota,omitempty"`
	// The remaining number of URLs that can be refreshed.
	//
	// example:
	//
	// 2000
	RefreshUrlRemain *int32 `json:"RefreshUrlRemain,omitempty" xml:"RefreshUrlRemain,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EF4F084A-2F49-4E1C-9CA0-DC85BCE7F391
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnUserQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserQuotaResponseBody) GetBlockQuota() *int32 {
	return s.BlockQuota
}

func (s *DescribeCdnUserQuotaResponseBody) GetBlockRemain() *int32 {
	return s.BlockRemain
}

func (s *DescribeCdnUserQuotaResponseBody) GetDomainQuota() *int32 {
	return s.DomainQuota
}

func (s *DescribeCdnUserQuotaResponseBody) GetIgnoreParamsQuota() *int32 {
	return s.IgnoreParamsQuota
}

func (s *DescribeCdnUserQuotaResponseBody) GetIgnoreParamsRemain() *int32 {
	return s.IgnoreParamsRemain
}

func (s *DescribeCdnUserQuotaResponseBody) GetPreloadQuota() *int32 {
	return s.PreloadQuota
}

func (s *DescribeCdnUserQuotaResponseBody) GetPreloadRemain() *int32 {
	return s.PreloadRemain
}

func (s *DescribeCdnUserQuotaResponseBody) GetRefreshDirQuota() *int32 {
	return s.RefreshDirQuota
}

func (s *DescribeCdnUserQuotaResponseBody) GetRefreshDirRemain() *int32 {
	return s.RefreshDirRemain
}

func (s *DescribeCdnUserQuotaResponseBody) GetRefreshUrlQuota() *int32 {
	return s.RefreshUrlQuota
}

func (s *DescribeCdnUserQuotaResponseBody) GetRefreshUrlRemain() *int32 {
	return s.RefreshUrlRemain
}

func (s *DescribeCdnUserQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnUserQuotaResponseBody) SetBlockQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetBlockRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetDomainQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.DomainQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetIgnoreParamsQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.IgnoreParamsQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetIgnoreParamsRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.IgnoreParamsRemain = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetPreloadQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetPreloadRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRefreshDirQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.RefreshDirQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRefreshDirRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.RefreshDirRemain = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRefreshUrlQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.RefreshUrlQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRefreshUrlRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.RefreshUrlRemain = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRequestId(v string) *DescribeCdnUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
