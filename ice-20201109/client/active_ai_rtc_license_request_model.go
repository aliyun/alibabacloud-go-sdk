// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveAiRtcLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *ActiveAiRtcLicenseRequest
	GetAuthCode() *string
	SetDeviceId(v string) *ActiveAiRtcLicenseRequest
	GetDeviceId() *string
	SetLicenseItemId(v string) *ActiveAiRtcLicenseRequest
	GetLicenseItemId() *string
}

type ActiveAiRtcLicenseRequest struct {
	// example:
	//
	// iU1IeJech7***
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// device-***
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 17712***
	LicenseItemId *string `json:"LicenseItemId,omitempty" xml:"LicenseItemId,omitempty"`
}

func (s ActiveAiRtcLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveAiRtcLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActiveAiRtcLicenseRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *ActiveAiRtcLicenseRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ActiveAiRtcLicenseRequest) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *ActiveAiRtcLicenseRequest) SetAuthCode(v string) *ActiveAiRtcLicenseRequest {
	s.AuthCode = &v
	return s
}

func (s *ActiveAiRtcLicenseRequest) SetDeviceId(v string) *ActiveAiRtcLicenseRequest {
	s.DeviceId = &v
	return s
}

func (s *ActiveAiRtcLicenseRequest) SetLicenseItemId(v string) *ActiveAiRtcLicenseRequest {
	s.LicenseItemId = &v
	return s
}

func (s *ActiveAiRtcLicenseRequest) Validate() error {
	return dara.Validate(s)
}
