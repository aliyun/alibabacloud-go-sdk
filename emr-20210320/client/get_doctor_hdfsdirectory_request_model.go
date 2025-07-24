// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHDFSDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHDFSDirectoryRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorHDFSDirectoryRequest
	GetDateTime() *string
	SetDirPath(v string) *GetDoctorHDFSDirectoryRequest
	GetDirPath() *string
	SetRegionId(v string) *GetDoctorHDFSDirectoryRequest
	GetRegionId() *string
}

type GetDoctorHDFSDirectoryRequest struct {
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
	// The directory name. The depth of the directory is not greater than five.
	//
	// This parameter is required.
	//
	// example:
	//
	// /tmp/test
	DirPath *string `json:"DirPath,omitempty" xml:"DirPath,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDoctorHDFSDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHDFSDirectoryRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHDFSDirectoryRequest) GetDirPath() *string {
	return s.DirPath
}

func (s *GetDoctorHDFSDirectoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHDFSDirectoryRequest) SetClusterId(v string) *GetDoctorHDFSDirectoryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHDFSDirectoryRequest) SetDateTime(v string) *GetDoctorHDFSDirectoryRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHDFSDirectoryRequest) SetDirPath(v string) *GetDoctorHDFSDirectoryRequest {
	s.DirPath = &v
	return s
}

func (s *GetDoctorHDFSDirectoryRequest) SetRegionId(v string) *GetDoctorHDFSDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHDFSDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
