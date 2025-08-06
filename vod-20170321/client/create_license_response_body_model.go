// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseInfo(v *CreateLicenseResponseBodyLicenseInfo) *CreateLicenseResponseBody
	GetLicenseInfo() *CreateLicenseResponseBodyLicenseInfo
	SetRequestId(v string) *CreateLicenseResponseBody
	GetRequestId() *string
}

type CreateLicenseResponseBody struct {
	LicenseInfo *CreateLicenseResponseBodyLicenseInfo `json:"LicenseInfo,omitempty" xml:"LicenseInfo,omitempty" type:"Struct"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLicenseResponseBody) GetLicenseInfo() *CreateLicenseResponseBodyLicenseInfo {
	return s.LicenseInfo
}

func (s *CreateLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLicenseResponseBody) SetLicenseInfo(v *CreateLicenseResponseBodyLicenseInfo) *CreateLicenseResponseBody {
	s.LicenseInfo = v
	return s
}

func (s *CreateLicenseResponseBody) SetRequestId(v string) *CreateLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateLicenseResponseBodyLicenseInfo struct {
	AccountId    *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	BeginTime    *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ContractNo   *string `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredOn    *string `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
	ExtraInfo    *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	LicenseId    *string `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
	ModifyTime   *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateLicenseResponseBodyLicenseInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateLicenseResponseBodyLicenseInfo) GoString() string {
	return s.String()
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetAccountId() *int64 {
	return s.AccountId
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetBeginTime() *string {
	return s.BeginTime
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetContractNo() *string {
	return s.ContractNo
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetLicenseId() *string {
	return s.LicenseId
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetStatus() *string {
	return s.Status
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetType() *string {
	return s.Type
}

func (s *CreateLicenseResponseBodyLicenseInfo) GetUserId() *int64 {
	return s.UserId
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetAccountId(v int64) *CreateLicenseResponseBodyLicenseInfo {
	s.AccountId = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetBeginTime(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.BeginTime = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetBusinessType(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.BusinessType = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetContractNo(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.ContractNo = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetCreateTime(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.CreateTime = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetExpiredOn(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.ExpiredOn = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetExtraInfo(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.ExtraInfo = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetLicenseId(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.LicenseId = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetModifyTime(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.ModifyTime = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetStatus(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.Status = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetType(v string) *CreateLicenseResponseBodyLicenseInfo {
	s.Type = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) SetUserId(v int64) *CreateLicenseResponseBodyLicenseInfo {
	s.UserId = &v
	return s
}

func (s *CreateLicenseResponseBodyLicenseInfo) Validate() error {
	return dara.Validate(s)
}
