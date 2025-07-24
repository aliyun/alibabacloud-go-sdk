// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetDoctorApplicationRequest
	GetAppId() *string
	SetClusterId(v string) *GetDoctorApplicationRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorApplicationRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorApplicationRequest
	GetRegionId() *string
}

type GetDoctorApplicationRequest struct {
	// The ID of the job that is submitted to YARN.
	//
	// This parameter is required.
	//
	// example:
	//
	// application_1542620905989_****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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

func (s GetDoctorApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetDoctorApplicationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorApplicationRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorApplicationRequest) SetAppId(v string) *GetDoctorApplicationRequest {
	s.AppId = &v
	return s
}

func (s *GetDoctorApplicationRequest) SetClusterId(v string) *GetDoctorApplicationRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorApplicationRequest) SetDateTime(v string) *GetDoctorApplicationRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorApplicationRequest) SetRegionId(v string) *GetDoctorApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorApplicationRequest) Validate() error {
	return dara.Validate(s)
}
