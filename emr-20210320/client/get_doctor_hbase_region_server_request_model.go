// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseRegionServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHBaseRegionServerRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorHBaseRegionServerRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorHBaseRegionServerRequest
	GetRegionId() *string
	SetRegionServerHost(v string) *GetDoctorHBaseRegionServerRequest
	GetRegionServerHost() *string
}

type GetDoctorHBaseRegionServerRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The host of the region server.
	//
	// This parameter is required.
	//
	// example:
	//
	// emr-worker-4.cluster-20****
	RegionServerHost *string `json:"RegionServerHost,omitempty" xml:"RegionServerHost,omitempty"`
}

func (s GetDoctorHBaseRegionServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHBaseRegionServerRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHBaseRegionServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHBaseRegionServerRequest) GetRegionServerHost() *string {
	return s.RegionServerHost
}

func (s *GetDoctorHBaseRegionServerRequest) SetClusterId(v string) *GetDoctorHBaseRegionServerRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHBaseRegionServerRequest) SetDateTime(v string) *GetDoctorHBaseRegionServerRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHBaseRegionServerRequest) SetRegionId(v string) *GetDoctorHBaseRegionServerRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHBaseRegionServerRequest) SetRegionServerHost(v string) *GetDoctorHBaseRegionServerRequest {
	s.RegionServerHost = &v
	return s
}

func (s *GetDoctorHBaseRegionServerRequest) Validate() error {
	return dara.Validate(s)
}
