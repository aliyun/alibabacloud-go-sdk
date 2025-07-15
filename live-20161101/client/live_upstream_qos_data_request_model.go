// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveUpstreamQosDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdnDomains(v []*string) *LiveUpstreamQosDataRequest
	GetCdnDomains() []*string
	SetCdnIsps(v []*string) *LiveUpstreamQosDataRequest
	GetCdnIsps() []*string
	SetCdnProvinces(v []*string) *LiveUpstreamQosDataRequest
	GetCdnProvinces() []*string
	SetEndTime(v string) *LiveUpstreamQosDataRequest
	GetEndTime() *string
	SetKwaiSidcs(v []*string) *LiveUpstreamQosDataRequest
	GetKwaiSidcs() []*string
	SetKwaiTsc(v []*int32) *LiveUpstreamQosDataRequest
	GetKwaiTsc() []*int32
	SetOwnerId(v int64) *LiveUpstreamQosDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *LiveUpstreamQosDataRequest
	GetRegion() *string
	SetRegionId(v string) *LiveUpstreamQosDataRequest
	GetRegionId() *string
	SetStartTime(v string) *LiveUpstreamQosDataRequest
	GetStartTime() *string
	SetUpstreamDomains(v []*string) *LiveUpstreamQosDataRequest
	GetUpstreamDomains() []*string
}

type LiveUpstreamQosDataRequest struct {
	CdnDomains   []*string `json:"CdnDomains,omitempty" xml:"CdnDomains,omitempty" type:"Repeated"`
	CdnIsps      []*string `json:"CdnIsps,omitempty" xml:"CdnIsps,omitempty" type:"Repeated"`
	CdnProvinces []*string `json:"CdnProvinces,omitempty" xml:"CdnProvinces,omitempty" type:"Repeated"`
	// This parameter is required.
	EndTime   *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	KwaiSidcs []*string `json:"KwaiSidcs,omitempty" xml:"KwaiSidcs,omitempty" type:"Repeated"`
	KwaiTsc   []*int32  `json:"KwaiTsc,omitempty" xml:"KwaiTsc,omitempty" type:"Repeated"`
	OwnerId   *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region    *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	StartTime       *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UpstreamDomains []*string `json:"UpstreamDomains,omitempty" xml:"UpstreamDomains,omitempty" type:"Repeated"`
}

func (s LiveUpstreamQosDataRequest) String() string {
	return dara.Prettify(s)
}

func (s LiveUpstreamQosDataRequest) GoString() string {
	return s.String()
}

func (s *LiveUpstreamQosDataRequest) GetCdnDomains() []*string {
	return s.CdnDomains
}

func (s *LiveUpstreamQosDataRequest) GetCdnIsps() []*string {
	return s.CdnIsps
}

func (s *LiveUpstreamQosDataRequest) GetCdnProvinces() []*string {
	return s.CdnProvinces
}

func (s *LiveUpstreamQosDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *LiveUpstreamQosDataRequest) GetKwaiSidcs() []*string {
	return s.KwaiSidcs
}

func (s *LiveUpstreamQosDataRequest) GetKwaiTsc() []*int32 {
	return s.KwaiTsc
}

func (s *LiveUpstreamQosDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LiveUpstreamQosDataRequest) GetRegion() *string {
	return s.Region
}

func (s *LiveUpstreamQosDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LiveUpstreamQosDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *LiveUpstreamQosDataRequest) GetUpstreamDomains() []*string {
	return s.UpstreamDomains
}

func (s *LiveUpstreamQosDataRequest) SetCdnDomains(v []*string) *LiveUpstreamQosDataRequest {
	s.CdnDomains = v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetCdnIsps(v []*string) *LiveUpstreamQosDataRequest {
	s.CdnIsps = v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetCdnProvinces(v []*string) *LiveUpstreamQosDataRequest {
	s.CdnProvinces = v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetEndTime(v string) *LiveUpstreamQosDataRequest {
	s.EndTime = &v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetKwaiSidcs(v []*string) *LiveUpstreamQosDataRequest {
	s.KwaiSidcs = v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetKwaiTsc(v []*int32) *LiveUpstreamQosDataRequest {
	s.KwaiTsc = v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetOwnerId(v int64) *LiveUpstreamQosDataRequest {
	s.OwnerId = &v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetRegion(v string) *LiveUpstreamQosDataRequest {
	s.Region = &v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetRegionId(v string) *LiveUpstreamQosDataRequest {
	s.RegionId = &v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetStartTime(v string) *LiveUpstreamQosDataRequest {
	s.StartTime = &v
	return s
}

func (s *LiveUpstreamQosDataRequest) SetUpstreamDomains(v []*string) *LiveUpstreamQosDataRequest {
	s.UpstreamDomains = v
	return s
}

func (s *LiveUpstreamQosDataRequest) Validate() error {
	return dara.Validate(s)
}
