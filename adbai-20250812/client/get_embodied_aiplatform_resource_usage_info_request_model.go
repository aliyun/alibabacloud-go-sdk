// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmbodiedAIPlatformResourceUsageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest
	GetDBClusterId() *string
	SetEndTime(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest
	GetEndTime() *string
	SetPlatformName(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest
	GetPlatformName() *string
	SetRegionId(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest
	GetRegionId() *string
	SetStartTime(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest
	GetStartTime() *string
}

type GetEmbodiedAIPlatformResourceUsageInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-01-20T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// platform1
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-01-10T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetEmbodiedAIPlatformResourceUsageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmbodiedAIPlatformResourceUsageInfoRequest) GoString() string {
	return s.String()
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) SetDBClusterId(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) SetEndTime(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest {
	s.EndTime = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) SetPlatformName(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest {
	s.PlatformName = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) SetRegionId(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) SetStartTime(v string) *GetEmbodiedAIPlatformResourceUsageInfoRequest {
	s.StartTime = &v
	return s
}

func (s *GetEmbodiedAIPlatformResourceUsageInfoRequest) Validate() error {
	return dara.Validate(s)
}
