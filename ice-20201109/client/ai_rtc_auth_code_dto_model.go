// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiRtcAuthCodeDTO interface {
	dara.Model
	String() string
	GoString() string
	SetActivatedTime(v string) *AiRtcAuthCodeDTO
	GetActivatedTime() *string
	SetAuthCode(v string) *AiRtcAuthCodeDTO
	GetAuthCode() *string
	SetCreationTime(v string) *AiRtcAuthCodeDTO
	GetCreationTime() *string
	SetDeviceId(v string) *AiRtcAuthCodeDTO
	GetDeviceId() *string
	SetLicense(v string) *AiRtcAuthCodeDTO
	GetLicense() *string
	SetLicenseItemId(v string) *AiRtcAuthCodeDTO
	GetLicenseItemId() *string
	SetModificationTime(v string) *AiRtcAuthCodeDTO
	GetModificationTime() *string
	SetStatus(v int32) *AiRtcAuthCodeDTO
	GetStatus() *int32
	SetType(v int32) *AiRtcAuthCodeDTO
	GetType() *int32
}

type AiRtcAuthCodeDTO struct {
	ActivatedTime    *string `json:"ActivatedTime,omitempty" xml:"ActivatedTime,omitempty"`
	AuthCode         *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	License          *string `json:"License,omitempty" xml:"License,omitempty"`
	LicenseItemId    *string `json:"LicenseItemId,omitempty" xml:"LicenseItemId,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AiRtcAuthCodeDTO) String() string {
	return dara.Prettify(s)
}

func (s AiRtcAuthCodeDTO) GoString() string {
	return s.String()
}

func (s *AiRtcAuthCodeDTO) GetActivatedTime() *string {
	return s.ActivatedTime
}

func (s *AiRtcAuthCodeDTO) GetAuthCode() *string {
	return s.AuthCode
}

func (s *AiRtcAuthCodeDTO) GetCreationTime() *string {
	return s.CreationTime
}

func (s *AiRtcAuthCodeDTO) GetDeviceId() *string {
	return s.DeviceId
}

func (s *AiRtcAuthCodeDTO) GetLicense() *string {
	return s.License
}

func (s *AiRtcAuthCodeDTO) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *AiRtcAuthCodeDTO) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *AiRtcAuthCodeDTO) GetStatus() *int32 {
	return s.Status
}

func (s *AiRtcAuthCodeDTO) GetType() *int32 {
	return s.Type
}

func (s *AiRtcAuthCodeDTO) SetActivatedTime(v string) *AiRtcAuthCodeDTO {
	s.ActivatedTime = &v
	return s
}

func (s *AiRtcAuthCodeDTO) SetAuthCode(v string) *AiRtcAuthCodeDTO {
	s.AuthCode = &v
	return s
}

func (s *AiRtcAuthCodeDTO) SetCreationTime(v string) *AiRtcAuthCodeDTO {
	s.CreationTime = &v
	return s
}

func (s *AiRtcAuthCodeDTO) SetDeviceId(v string) *AiRtcAuthCodeDTO {
	s.DeviceId = &v
	return s
}

func (s *AiRtcAuthCodeDTO) SetLicense(v string) *AiRtcAuthCodeDTO {
	s.License = &v
	return s
}

func (s *AiRtcAuthCodeDTO) SetLicenseItemId(v string) *AiRtcAuthCodeDTO {
	s.LicenseItemId = &v
	return s
}

func (s *AiRtcAuthCodeDTO) SetModificationTime(v string) *AiRtcAuthCodeDTO {
	s.ModificationTime = &v
	return s
}

func (s *AiRtcAuthCodeDTO) SetStatus(v int32) *AiRtcAuthCodeDTO {
	s.Status = &v
	return s
}

func (s *AiRtcAuthCodeDTO) SetType(v int32) *AiRtcAuthCodeDTO {
	s.Type = &v
	return s
}

func (s *AiRtcAuthCodeDTO) Validate() error {
	return dara.Validate(s)
}
