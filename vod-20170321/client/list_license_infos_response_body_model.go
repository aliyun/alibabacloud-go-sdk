// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLicenseInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseList(v []*ListLicenseInfosResponseBodyLicenseList) *ListLicenseInfosResponseBody
	GetLicenseList() []*ListLicenseInfosResponseBodyLicenseList
	SetRequestId(v string) *ListLicenseInfosResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListLicenseInfosResponseBody
	GetTotal() *int64
}

type ListLicenseInfosResponseBody struct {
	LicenseList []*ListLicenseInfosResponseBodyLicenseList `json:"LicenseList,omitempty" xml:"LicenseList,omitempty" type:"Repeated"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total       *int64                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLicenseInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLicenseInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListLicenseInfosResponseBody) GetLicenseList() []*ListLicenseInfosResponseBodyLicenseList {
	return s.LicenseList
}

func (s *ListLicenseInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLicenseInfosResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListLicenseInfosResponseBody) SetLicenseList(v []*ListLicenseInfosResponseBodyLicenseList) *ListLicenseInfosResponseBody {
	s.LicenseList = v
	return s
}

func (s *ListLicenseInfosResponseBody) SetRequestId(v string) *ListLicenseInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLicenseInfosResponseBody) SetTotal(v int64) *ListLicenseInfosResponseBody {
	s.Total = &v
	return s
}

func (s *ListLicenseInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLicenseInfosResponseBodyLicenseList struct {
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

func (s ListLicenseInfosResponseBodyLicenseList) String() string {
	return dara.Prettify(s)
}

func (s ListLicenseInfosResponseBodyLicenseList) GoString() string {
	return s.String()
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetContractNo() *string {
	return s.ContractNo
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetLicenseId() *string {
	return s.LicenseId
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetStatus() *string {
	return s.Status
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetType() *string {
	return s.Type
}

func (s *ListLicenseInfosResponseBodyLicenseList) GetUserId() *int64 {
	return s.UserId
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetAccountId(v int64) *ListLicenseInfosResponseBodyLicenseList {
	s.AccountId = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetBeginTime(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.BeginTime = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetBusinessType(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.BusinessType = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetContractNo(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.ContractNo = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetCreateTime(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.CreateTime = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetExpiredOn(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.ExpiredOn = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetExtraInfo(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.ExtraInfo = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetLicenseId(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.LicenseId = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetModifyTime(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.ModifyTime = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetStatus(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.Status = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetType(v string) *ListLicenseInfosResponseBodyLicenseList {
	s.Type = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) SetUserId(v int64) *ListLicenseInfosResponseBodyLicenseList {
	s.UserId = &v
	return s
}

func (s *ListLicenseInfosResponseBodyLicenseList) Validate() error {
	return dara.Validate(s)
}
