// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHiveClusterRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorHiveClusterRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorHiveClusterRequest
	GetRegionId() *string
}

type GetDoctorHiveClusterRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specify the date in the ISO 8601 standard. For example, 2023-01-01 represents January 1, 2023.
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

func (s GetDoctorHiveClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHiveClusterRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHiveClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHiveClusterRequest) SetClusterId(v string) *GetDoctorHiveClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHiveClusterRequest) SetDateTime(v string) *GetDoctorHiveClusterRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHiveClusterRequest) SetRegionId(v string) *GetDoctorHiveClusterRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHiveClusterRequest) Validate() error {
	return dara.Validate(s)
}
