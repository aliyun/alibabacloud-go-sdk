// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLicenseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseInfo(v *ModifyLicenseInfoResponseBodyLicenseInfo) *ModifyLicenseInfoResponseBody
	GetLicenseInfo() *ModifyLicenseInfoResponseBodyLicenseInfo
	SetRequestId(v string) *ModifyLicenseInfoResponseBody
	GetRequestId() *string
}

type ModifyLicenseInfoResponseBody struct {
	LicenseInfo *ModifyLicenseInfoResponseBodyLicenseInfo `json:"LicenseInfo,omitempty" xml:"LicenseInfo,omitempty" type:"Struct"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLicenseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLicenseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLicenseInfoResponseBody) GetLicenseInfo() *ModifyLicenseInfoResponseBodyLicenseInfo {
	return s.LicenseInfo
}

func (s *ModifyLicenseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLicenseInfoResponseBody) SetLicenseInfo(v *ModifyLicenseInfoResponseBodyLicenseInfo) *ModifyLicenseInfoResponseBody {
	s.LicenseInfo = v
	return s
}

func (s *ModifyLicenseInfoResponseBody) SetRequestId(v string) *ModifyLicenseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLicenseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyLicenseInfoResponseBodyLicenseInfo struct {
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

func (s ModifyLicenseInfoResponseBodyLicenseInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyLicenseInfoResponseBodyLicenseInfo) GoString() string {
	return s.String()
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetContractNo() *string {
	return s.ContractNo
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetLicenseId() *string {
	return s.LicenseId
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetStatus() *string {
	return s.Status
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetType() *string {
	return s.Type
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) GetUserId() *int64 {
	return s.UserId
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetAccountId(v int64) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.AccountId = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetBeginTime(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.BeginTime = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetBusinessType(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.BusinessType = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetContractNo(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.ContractNo = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetCreateTime(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.CreateTime = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetExpiredOn(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.ExpiredOn = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetExtraInfo(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.ExtraInfo = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetLicenseId(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.LicenseId = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetModifyTime(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.ModifyTime = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetStatus(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.Status = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetType(v string) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.Type = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) SetUserId(v int64) *ModifyLicenseInfoResponseBodyLicenseInfo {
	s.UserId = &v
	return s
}

func (s *ModifyLicenseInfoResponseBodyLicenseInfo) Validate() error {
	return dara.Validate(s)
}
