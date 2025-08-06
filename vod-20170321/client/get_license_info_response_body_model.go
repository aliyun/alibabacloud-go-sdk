// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseInfo(v *GetLicenseInfoResponseBodyLicenseInfo) *GetLicenseInfoResponseBody
	GetLicenseInfo() *GetLicenseInfoResponseBodyLicenseInfo
	SetRequestId(v string) *GetLicenseInfoResponseBody
	GetRequestId() *string
}

type GetLicenseInfoResponseBody struct {
	LicenseInfo *GetLicenseInfoResponseBodyLicenseInfo `json:"LicenseInfo,omitempty" xml:"LicenseInfo,omitempty" type:"Struct"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLicenseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicenseInfoResponseBody) GetLicenseInfo() *GetLicenseInfoResponseBodyLicenseInfo {
	return s.LicenseInfo
}

func (s *GetLicenseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLicenseInfoResponseBody) SetLicenseInfo(v *GetLicenseInfoResponseBodyLicenseInfo) *GetLicenseInfoResponseBody {
	s.LicenseInfo = v
	return s
}

func (s *GetLicenseInfoResponseBody) SetRequestId(v string) *GetLicenseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLicenseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLicenseInfoResponseBodyLicenseInfo struct {
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

func (s GetLicenseInfoResponseBodyLicenseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseInfoResponseBodyLicenseInfo) GoString() string {
	return s.String()
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetContractNo() *string {
	return s.ContractNo
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetLicenseId() *string {
	return s.LicenseId
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetStatus() *string {
	return s.Status
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetType() *string {
	return s.Type
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) GetUserId() *int64 {
	return s.UserId
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetAccountId(v int64) *GetLicenseInfoResponseBodyLicenseInfo {
	s.AccountId = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetBeginTime(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.BeginTime = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetBusinessType(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.BusinessType = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetContractNo(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.ContractNo = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetCreateTime(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetExpiredOn(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.ExpiredOn = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetExtraInfo(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.ExtraInfo = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetLicenseId(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.LicenseId = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetModifyTime(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetStatus(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.Status = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetType(v string) *GetLicenseInfoResponseBodyLicenseInfo {
	s.Type = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) SetUserId(v int64) *GetLicenseInfoResponseBodyLicenseInfo {
	s.UserId = &v
	return s
}

func (s *GetLicenseInfoResponseBodyLicenseInfo) Validate() error {
	return dara.Validate(s)
}
