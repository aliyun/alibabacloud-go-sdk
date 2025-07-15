// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveUpstreamQosDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdnDomainsShrink(v string) *LiveUpstreamQosDataShrinkRequest
	GetCdnDomainsShrink() *string
	SetCdnIspsShrink(v string) *LiveUpstreamQosDataShrinkRequest
	GetCdnIspsShrink() *string
	SetCdnProvincesShrink(v string) *LiveUpstreamQosDataShrinkRequest
	GetCdnProvincesShrink() *string
	SetEndTime(v string) *LiveUpstreamQosDataShrinkRequest
	GetEndTime() *string
	SetKwaiSidcsShrink(v string) *LiveUpstreamQosDataShrinkRequest
	GetKwaiSidcsShrink() *string
	SetKwaiTscShrink(v string) *LiveUpstreamQosDataShrinkRequest
	GetKwaiTscShrink() *string
	SetOwnerId(v int64) *LiveUpstreamQosDataShrinkRequest
	GetOwnerId() *int64
	SetRegion(v string) *LiveUpstreamQosDataShrinkRequest
	GetRegion() *string
	SetRegionId(v string) *LiveUpstreamQosDataShrinkRequest
	GetRegionId() *string
	SetStartTime(v string) *LiveUpstreamQosDataShrinkRequest
	GetStartTime() *string
	SetUpstreamDomainsShrink(v string) *LiveUpstreamQosDataShrinkRequest
	GetUpstreamDomainsShrink() *string
}

type LiveUpstreamQosDataShrinkRequest struct {
	CdnDomainsShrink   *string `json:"CdnDomains,omitempty" xml:"CdnDomains,omitempty"`
	CdnIspsShrink      *string `json:"CdnIsps,omitempty" xml:"CdnIsps,omitempty"`
	CdnProvincesShrink *string `json:"CdnProvinces,omitempty" xml:"CdnProvinces,omitempty"`
	// This parameter is required.
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	KwaiSidcsShrink *string `json:"KwaiSidcs,omitempty" xml:"KwaiSidcs,omitempty"`
	KwaiTscShrink   *string `json:"KwaiTsc,omitempty" xml:"KwaiTsc,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UpstreamDomainsShrink *string `json:"UpstreamDomains,omitempty" xml:"UpstreamDomains,omitempty"`
}

func (s LiveUpstreamQosDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s LiveUpstreamQosDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *LiveUpstreamQosDataShrinkRequest) GetCdnDomainsShrink() *string {
	return s.CdnDomainsShrink
}

func (s *LiveUpstreamQosDataShrinkRequest) GetCdnIspsShrink() *string {
	return s.CdnIspsShrink
}

func (s *LiveUpstreamQosDataShrinkRequest) GetCdnProvincesShrink() *string {
	return s.CdnProvincesShrink
}

func (s *LiveUpstreamQosDataShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *LiveUpstreamQosDataShrinkRequest) GetKwaiSidcsShrink() *string {
	return s.KwaiSidcsShrink
}

func (s *LiveUpstreamQosDataShrinkRequest) GetKwaiTscShrink() *string {
	return s.KwaiTscShrink
}

func (s *LiveUpstreamQosDataShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LiveUpstreamQosDataShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *LiveUpstreamQosDataShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LiveUpstreamQosDataShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *LiveUpstreamQosDataShrinkRequest) GetUpstreamDomainsShrink() *string {
	return s.UpstreamDomainsShrink
}

func (s *LiveUpstreamQosDataShrinkRequest) SetCdnDomainsShrink(v string) *LiveUpstreamQosDataShrinkRequest {
	s.CdnDomainsShrink = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetCdnIspsShrink(v string) *LiveUpstreamQosDataShrinkRequest {
	s.CdnIspsShrink = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetCdnProvincesShrink(v string) *LiveUpstreamQosDataShrinkRequest {
	s.CdnProvincesShrink = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetEndTime(v string) *LiveUpstreamQosDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetKwaiSidcsShrink(v string) *LiveUpstreamQosDataShrinkRequest {
	s.KwaiSidcsShrink = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetKwaiTscShrink(v string) *LiveUpstreamQosDataShrinkRequest {
	s.KwaiTscShrink = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetOwnerId(v int64) *LiveUpstreamQosDataShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetRegion(v string) *LiveUpstreamQosDataShrinkRequest {
	s.Region = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetRegionId(v string) *LiveUpstreamQosDataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetStartTime(v string) *LiveUpstreamQosDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) SetUpstreamDomainsShrink(v string) *LiveUpstreamQosDataShrinkRequest {
	s.UpstreamDomainsShrink = &v
	return s
}

func (s *LiveUpstreamQosDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
