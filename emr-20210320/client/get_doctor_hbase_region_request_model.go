// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHBaseRegionRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorHBaseRegionRequest
	GetDateTime() *string
	SetHbaseRegionId(v string) *GetDoctorHBaseRegionRequest
	GetHbaseRegionId() *string
	SetRegionId(v string) *GetDoctorHBaseRegionRequest
	GetRegionId() *string
}

type GetDoctorHBaseRegionRequest struct {
	// Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 67f6808f60a8c357103a3a95fe00610e
	HbaseRegionId *string `json:"HbaseRegionId,omitempty" xml:"HbaseRegionId,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDoctorHBaseRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHBaseRegionRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHBaseRegionRequest) GetHbaseRegionId() *string {
	return s.HbaseRegionId
}

func (s *GetDoctorHBaseRegionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHBaseRegionRequest) SetClusterId(v string) *GetDoctorHBaseRegionRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHBaseRegionRequest) SetDateTime(v string) *GetDoctorHBaseRegionRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHBaseRegionRequest) SetHbaseRegionId(v string) *GetDoctorHBaseRegionRequest {
	s.HbaseRegionId = &v
	return s
}

func (s *GetDoctorHBaseRegionRequest) SetRegionId(v string) *GetDoctorHBaseRegionRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHBaseRegionRequest) Validate() error {
	return dara.Validate(s)
}
