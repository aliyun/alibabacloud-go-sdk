// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiRtcLicenseInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableCapacity(v int64) *AiRtcLicenseInfoDTO
	GetAvailableCapacity() *int64
	SetBeginOn(v string) *AiRtcLicenseInfoDTO
	GetBeginOn() *string
	SetContractNo(v string) *AiRtcLicenseInfoDTO
	GetContractNo() *string
	SetCreationTime(v string) *AiRtcLicenseInfoDTO
	GetCreationTime() *string
	SetExpiredOn(v string) *AiRtcLicenseInfoDTO
	GetExpiredOn() *string
	SetInstanceId(v string) *AiRtcLicenseInfoDTO
	GetInstanceId() *string
	SetLicenseCount(v int64) *AiRtcLicenseInfoDTO
	GetLicenseCount() *int64
	SetLicenseItemId(v string) *AiRtcLicenseInfoDTO
	GetLicenseItemId() *string
	SetModificationTime(v string) *AiRtcLicenseInfoDTO
	GetModificationTime() *string
	SetStatus(v int32) *AiRtcLicenseInfoDTO
	GetStatus() *int32
	SetType(v int32) *AiRtcLicenseInfoDTO
	GetType() *int32
	SetValidDays(v int64) *AiRtcLicenseInfoDTO
	GetValidDays() *int64
}

type AiRtcLicenseInfoDTO struct {
	AvailableCapacity *int64  `json:"AvailableCapacity,omitempty" xml:"AvailableCapacity,omitempty"`
	BeginOn           *string `json:"BeginOn,omitempty" xml:"BeginOn,omitempty"`
	ContractNo        *string `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	CreationTime      *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ExpiredOn         *string `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LicenseCount      *int64  `json:"LicenseCount,omitempty" xml:"LicenseCount,omitempty"`
	LicenseItemId     *string `json:"LicenseItemId,omitempty" xml:"LicenseItemId,omitempty"`
	ModificationTime  *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type              *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	ValidDays         *int64  `json:"ValidDays,omitempty" xml:"ValidDays,omitempty"`
}

func (s AiRtcLicenseInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s AiRtcLicenseInfoDTO) GoString() string {
	return s.String()
}

func (s *AiRtcLicenseInfoDTO) GetAvailableCapacity() *int64 {
	return s.AvailableCapacity
}

func (s *AiRtcLicenseInfoDTO) GetBeginOn() *string {
	return s.BeginOn
}

func (s *AiRtcLicenseInfoDTO) GetContractNo() *string {
	return s.ContractNo
}

func (s *AiRtcLicenseInfoDTO) GetCreationTime() *string {
	return s.CreationTime
}

func (s *AiRtcLicenseInfoDTO) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *AiRtcLicenseInfoDTO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AiRtcLicenseInfoDTO) GetLicenseCount() *int64 {
	return s.LicenseCount
}

func (s *AiRtcLicenseInfoDTO) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *AiRtcLicenseInfoDTO) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *AiRtcLicenseInfoDTO) GetStatus() *int32 {
	return s.Status
}

func (s *AiRtcLicenseInfoDTO) GetType() *int32 {
	return s.Type
}

func (s *AiRtcLicenseInfoDTO) GetValidDays() *int64 {
	return s.ValidDays
}

func (s *AiRtcLicenseInfoDTO) SetAvailableCapacity(v int64) *AiRtcLicenseInfoDTO {
	s.AvailableCapacity = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetBeginOn(v string) *AiRtcLicenseInfoDTO {
	s.BeginOn = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetContractNo(v string) *AiRtcLicenseInfoDTO {
	s.ContractNo = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetCreationTime(v string) *AiRtcLicenseInfoDTO {
	s.CreationTime = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetExpiredOn(v string) *AiRtcLicenseInfoDTO {
	s.ExpiredOn = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetInstanceId(v string) *AiRtcLicenseInfoDTO {
	s.InstanceId = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetLicenseCount(v int64) *AiRtcLicenseInfoDTO {
	s.LicenseCount = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetLicenseItemId(v string) *AiRtcLicenseInfoDTO {
	s.LicenseItemId = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetModificationTime(v string) *AiRtcLicenseInfoDTO {
	s.ModificationTime = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetStatus(v int32) *AiRtcLicenseInfoDTO {
	s.Status = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetType(v int32) *AiRtcLicenseInfoDTO {
	s.Type = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) SetValidDays(v int64) *AiRtcLicenseInfoDTO {
	s.ValidDays = &v
	return s
}

func (s *AiRtcLicenseInfoDTO) Validate() error {
	return dara.Validate(s)
}
