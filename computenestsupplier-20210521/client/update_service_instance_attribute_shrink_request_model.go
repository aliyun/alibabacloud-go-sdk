// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *UpdateServiceInstanceAttributeShrinkRequest
	GetEndTime() *string
	SetLicenseDataShrink(v string) *UpdateServiceInstanceAttributeShrinkRequest
	GetLicenseDataShrink() *string
	SetReason(v string) *UpdateServiceInstanceAttributeShrinkRequest
	GetReason() *string
	SetRegionId(v string) *UpdateServiceInstanceAttributeShrinkRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *UpdateServiceInstanceAttributeShrinkRequest
	GetServiceInstanceId() *string
}

type UpdateServiceInstanceAttributeShrinkRequest struct {
	// The time when the service instance expires.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-25T02:28:40Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The License Data
	LicenseDataShrink *string `json:"LicenseData,omitempty" xml:"LicenseData,omitempty"`
	// Application reason, currently used for trial application extension.
	//
	// example:
	//
	// \\"\\"
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-3df88e962cdexxxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) GetLicenseDataShrink() *string {
	return s.LicenseDataShrink
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetEndTime(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetLicenseDataShrink(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.LicenseDataShrink = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetReason(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.Reason = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetRegionId(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
