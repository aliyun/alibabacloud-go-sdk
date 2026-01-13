// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepaidRecord interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PrepaidRecord
	GetAccountId() *string
	SetAccountName(v string) *PrepaidRecord
	GetAccountName() *string
	SetCodeType(v string) *PrepaidRecord
	GetCodeType() *string
	SetGmtCreate(v string) *PrepaidRecord
	GetGmtCreate() *string
	SetGmtModified(v string) *PrepaidRecord
	GetGmtModified() *string
	SetId(v int64) *PrepaidRecord
	GetId() *int64
	SetRemainQuota(v int64) *PrepaidRecord
	GetRemainQuota() *int64
	SetStatus(v string) *PrepaidRecord
	GetStatus() *string
	SetTotalQuota(v int64) *PrepaidRecord
	GetTotalQuota() *int64
	SetUsedQuota(v int64) *PrepaidRecord
	GetUsedQuota() *int64
}

type PrepaidRecord struct {
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	CodeType    *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RemainQuota *int64  `json:"remainQuota,omitempty" xml:"remainQuota,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	TotalQuota  *int64  `json:"totalQuota,omitempty" xml:"totalQuota,omitempty"`
	UsedQuota   *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s PrepaidRecord) String() string {
	return dara.Prettify(s)
}

func (s PrepaidRecord) GoString() string {
	return s.String()
}

func (s *PrepaidRecord) GetAccountId() *string {
	return s.AccountId
}

func (s *PrepaidRecord) GetAccountName() *string {
	return s.AccountName
}

func (s *PrepaidRecord) GetCodeType() *string {
	return s.CodeType
}

func (s *PrepaidRecord) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PrepaidRecord) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PrepaidRecord) GetId() *int64 {
	return s.Id
}

func (s *PrepaidRecord) GetRemainQuota() *int64 {
	return s.RemainQuota
}

func (s *PrepaidRecord) GetStatus() *string {
	return s.Status
}

func (s *PrepaidRecord) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *PrepaidRecord) GetUsedQuota() *int64 {
	return s.UsedQuota
}

func (s *PrepaidRecord) SetAccountId(v string) *PrepaidRecord {
	s.AccountId = &v
	return s
}

func (s *PrepaidRecord) SetAccountName(v string) *PrepaidRecord {
	s.AccountName = &v
	return s
}

func (s *PrepaidRecord) SetCodeType(v string) *PrepaidRecord {
	s.CodeType = &v
	return s
}

func (s *PrepaidRecord) SetGmtCreate(v string) *PrepaidRecord {
	s.GmtCreate = &v
	return s
}

func (s *PrepaidRecord) SetGmtModified(v string) *PrepaidRecord {
	s.GmtModified = &v
	return s
}

func (s *PrepaidRecord) SetId(v int64) *PrepaidRecord {
	s.Id = &v
	return s
}

func (s *PrepaidRecord) SetRemainQuota(v int64) *PrepaidRecord {
	s.RemainQuota = &v
	return s
}

func (s *PrepaidRecord) SetStatus(v string) *PrepaidRecord {
	s.Status = &v
	return s
}

func (s *PrepaidRecord) SetTotalQuota(v int64) *PrepaidRecord {
	s.TotalQuota = &v
	return s
}

func (s *PrepaidRecord) SetUsedQuota(v int64) *PrepaidRecord {
	s.UsedQuota = &v
	return s
}

func (s *PrepaidRecord) Validate() error {
	return dara.Validate(s)
}
