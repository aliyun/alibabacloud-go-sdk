// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHBaseClusterRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorHBaseClusterRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorHBaseClusterRequest
	GetRegionId() *string
}

type GetDoctorHBaseClusterRequest struct {
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
}

func (s GetDoctorHBaseClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHBaseClusterRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHBaseClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHBaseClusterRequest) SetClusterId(v string) *GetDoctorHBaseClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHBaseClusterRequest) SetDateTime(v string) *GetDoctorHBaseClusterRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHBaseClusterRequest) SetRegionId(v string) *GetDoctorHBaseClusterRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHBaseClusterRequest) Validate() error {
	return dara.Validate(s)
}
