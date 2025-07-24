// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLocale(v string) *GetDoctorApplicationRequest
	GetLocale() *string
	SetQueryTime(v string) *GetDoctorApplicationRequest
	GetQueryTime() *string
	SetRegionId(v string) *GetDoctorApplicationRequest
	GetRegionId() *string
}

type GetDoctorApplicationRequest struct {
	// The language of diagnostic information.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The query time.
	//
	// example:
	//
	// 2024-01-01
	QueryTime *string `json:"queryTime,omitempty" xml:"queryTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetDoctorApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetDoctorApplicationRequest) GetQueryTime() *string {
	return s.QueryTime
}

func (s *GetDoctorApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorApplicationRequest) SetLocale(v string) *GetDoctorApplicationRequest {
	s.Locale = &v
	return s
}

func (s *GetDoctorApplicationRequest) SetQueryTime(v string) *GetDoctorApplicationRequest {
	s.QueryTime = &v
	return s
}

func (s *GetDoctorApplicationRequest) SetRegionId(v string) *GetDoctorApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorApplicationRequest) Validate() error {
	return dara.Validate(s)
}
