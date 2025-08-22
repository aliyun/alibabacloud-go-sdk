// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlockQuota(v int32) *DescribeDcdnUserQuotaResponseBody
	GetBlockQuota() *int32
	SetBlockRemain(v int32) *DescribeDcdnUserQuotaResponseBody
	GetBlockRemain() *int32
	SetDomainQuota(v int32) *DescribeDcdnUserQuotaResponseBody
	GetDomainQuota() *int32
	SetIgnoreParamsQuota(v int32) *DescribeDcdnUserQuotaResponseBody
	GetIgnoreParamsQuota() *int32
	SetIgnoreParamsRemain(v int32) *DescribeDcdnUserQuotaResponseBody
	GetIgnoreParamsRemain() *int32
	SetPreloadQuota(v int32) *DescribeDcdnUserQuotaResponseBody
	GetPreloadQuota() *int32
	SetPreloadRemain(v int32) *DescribeDcdnUserQuotaResponseBody
	GetPreloadRemain() *int32
	SetRefreshDirQuota(v int32) *DescribeDcdnUserQuotaResponseBody
	GetRefreshDirQuota() *int32
	SetRefreshDirRemain(v int32) *DescribeDcdnUserQuotaResponseBody
	GetRefreshDirRemain() *int32
	SetRefreshUrlQuota(v int32) *DescribeDcdnUserQuotaResponseBody
	GetRefreshUrlQuota() *int32
	SetRefreshUrlRemain(v int32) *DescribeDcdnUserQuotaResponseBody
	GetRefreshUrlRemain() *int32
	SetRequestId(v string) *DescribeDcdnUserQuotaResponseBody
	GetRequestId() *string
}

type DescribeDcdnUserQuotaResponseBody struct {
	// The maximum number of URLs that can be blocked.
	//
	// example:
	//
	// 20
	BlockQuota *int32 `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	// The remaining number of URLs that can be blocked.
	//
	// example:
	//
	// 500
	BlockRemain *int32 `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	// The maximum number of accelerated domains.
	//
	// example:
	//
	// 50
	DomainQuota *int32 `json:"DomainQuota,omitempty" xml:"DomainQuota,omitempty"`
	// The maximum number of URLs or directories with parameters ignored that can be refreshed.
	//
	// example:
	//
	// 100
	IgnoreParamsQuota *int32 `json:"IgnoreParamsQuota,omitempty" xml:"IgnoreParamsQuota,omitempty"`
	// The number of remaining URLs or directories with parameters ignored that can be refreshed.
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
	// 300
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
	// 100
	RefreshDirRemain *int32 `json:"RefreshDirRemain,omitempty" xml:"RefreshDirRemain,omitempty"`
	// The maximum number of URLs that can be refreshed.
	//
	// example:
	//
	// 100
	RefreshUrlQuota *int32 `json:"RefreshUrlQuota,omitempty" xml:"RefreshUrlQuota,omitempty"`
	// The remaining number of URLs that can be refreshed.
	//
	// example:
	//
	// 100
	RefreshUrlRemain *int32 `json:"RefreshUrlRemain,omitempty" xml:"RefreshUrlRemain,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BFFCDFAD-DACC-484E-9BE6-0AF3B3A0DD23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserQuotaResponseBody) GetBlockQuota() *int32 {
	return s.BlockQuota
}

func (s *DescribeDcdnUserQuotaResponseBody) GetBlockRemain() *int32 {
	return s.BlockRemain
}

func (s *DescribeDcdnUserQuotaResponseBody) GetDomainQuota() *int32 {
	return s.DomainQuota
}

func (s *DescribeDcdnUserQuotaResponseBody) GetIgnoreParamsQuota() *int32 {
	return s.IgnoreParamsQuota
}

func (s *DescribeDcdnUserQuotaResponseBody) GetIgnoreParamsRemain() *int32 {
	return s.IgnoreParamsRemain
}

func (s *DescribeDcdnUserQuotaResponseBody) GetPreloadQuota() *int32 {
	return s.PreloadQuota
}

func (s *DescribeDcdnUserQuotaResponseBody) GetPreloadRemain() *int32 {
	return s.PreloadRemain
}

func (s *DescribeDcdnUserQuotaResponseBody) GetRefreshDirQuota() *int32 {
	return s.RefreshDirQuota
}

func (s *DescribeDcdnUserQuotaResponseBody) GetRefreshDirRemain() *int32 {
	return s.RefreshDirRemain
}

func (s *DescribeDcdnUserQuotaResponseBody) GetRefreshUrlQuota() *int32 {
	return s.RefreshUrlQuota
}

func (s *DescribeDcdnUserQuotaResponseBody) GetRefreshUrlRemain() *int32 {
	return s.RefreshUrlRemain
}

func (s *DescribeDcdnUserQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserQuotaResponseBody) SetBlockQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetBlockRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetDomainQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.DomainQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetIgnoreParamsQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.IgnoreParamsQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetIgnoreParamsRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.IgnoreParamsRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetPreloadQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetPreloadRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRefreshDirQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.RefreshDirQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRefreshDirRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.RefreshDirRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRefreshUrlQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.RefreshUrlQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRefreshUrlRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.RefreshUrlRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRequestId(v string) *DescribeDcdnUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
