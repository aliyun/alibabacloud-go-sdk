// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHDFSClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHDFSClusterRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorHDFSClusterRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorHDFSClusterRequest
	GetRegionId() *string
}

type GetDoctorHDFSClusterRequest struct {
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

func (s GetDoctorHDFSClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHDFSClusterRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHDFSClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHDFSClusterRequest) SetClusterId(v string) *GetDoctorHDFSClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHDFSClusterRequest) SetDateTime(v string) *GetDoctorHDFSClusterRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHDFSClusterRequest) SetRegionId(v string) *GetDoctorHDFSClusterRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHDFSClusterRequest) Validate() error {
	return dara.Validate(s)
}
