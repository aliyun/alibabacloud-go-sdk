// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetDoctorJobRequest
	GetAppId() *string
	SetClusterId(v string) *GetDoctorJobRequest
	GetClusterId() *string
	SetRegionId(v string) *GetDoctorJobRequest
	GetRegionId() *string
}

type GetDoctorJobRequest struct {
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDoctorJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorJobRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetDoctorJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorJobRequest) SetAppId(v string) *GetDoctorJobRequest {
	s.AppId = &v
	return s
}

func (s *GetDoctorJobRequest) SetClusterId(v string) *GetDoctorJobRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorJobRequest) SetRegionId(v string) *GetDoctorJobRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorJobRequest) Validate() error {
	return dara.Validate(s)
}
