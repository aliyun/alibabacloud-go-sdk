// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFundAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListFundAccountResponseBodyData) *ListFundAccountResponseBody
	GetData() []*ListFundAccountResponseBodyData
	SetMetadata(v interface{}) *ListFundAccountResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *ListFundAccountResponseBody
	GetRequestId() *string
}

type ListFundAccountResponseBody struct {
	Data []*ListFundAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFundAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFundAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListFundAccountResponseBody) GetData() []*ListFundAccountResponseBodyData {
	return s.Data
}

func (s *ListFundAccountResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ListFundAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFundAccountResponseBody) SetData(v []*ListFundAccountResponseBodyData) *ListFundAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListFundAccountResponseBody) SetMetadata(v interface{}) *ListFundAccountResponseBody {
	s.Metadata = v
	return s
}

func (s *ListFundAccountResponseBody) SetRequestId(v string) *ListFundAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFundAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFundAccountResponseBodyData struct {
	// example:
	//
	// 2024-12-30 12:00:00
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// example:
	//
	// 1232121
	FundAccountAdminAccountId *string `json:"FundAccountAdminAccountId,omitempty" xml:"FundAccountAdminAccountId,omitempty"`
	// example:
	//
	// 云某的账户
	FundAccountAdminAccountName *string `json:"FundAccountAdminAccountName,omitempty" xml:"FundAccountAdminAccountName,omitempty"`
	// example:
	//
	// 1022231
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 云某的账户
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// example:
	//
	// 132123211
	FundAccountOwnerAccountId *string `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// VALID
	FundAccountStatus *string `json:"FundAccountStatus,omitempty" xml:"FundAccountStatus,omitempty"`
	// example:
	//
	// DIRECT_USER
	FundAccountType *string `json:"FundAccountType,omitempty" xml:"FundAccountType,omitempty"`
	// example:
	//
	// 2684210001
	Nbid        *string   `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	Permissions []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s ListFundAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFundAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFundAccountResponseBodyData) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListFundAccountResponseBodyData) GetFundAccountAdminAccountId() *string {
	return s.FundAccountAdminAccountId
}

func (s *ListFundAccountResponseBodyData) GetFundAccountAdminAccountName() *string {
	return s.FundAccountAdminAccountName
}

func (s *ListFundAccountResponseBodyData) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *ListFundAccountResponseBodyData) GetFundAccountName() *string {
	return s.FundAccountName
}

func (s *ListFundAccountResponseBodyData) GetFundAccountOwnerAccountId() *string {
	return s.FundAccountOwnerAccountId
}

func (s *ListFundAccountResponseBodyData) GetFundAccountStatus() *string {
	return s.FundAccountStatus
}

func (s *ListFundAccountResponseBodyData) GetFundAccountType() *string {
	return s.FundAccountType
}

func (s *ListFundAccountResponseBodyData) GetNbid() *string {
	return s.Nbid
}

func (s *ListFundAccountResponseBodyData) GetPermissions() []*string {
	return s.Permissions
}

func (s *ListFundAccountResponseBodyData) GetSite() *string {
	return s.Site
}

func (s *ListFundAccountResponseBodyData) SetCreateDate(v string) *ListFundAccountResponseBodyData {
	s.CreateDate = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountAdminAccountId(v string) *ListFundAccountResponseBodyData {
	s.FundAccountAdminAccountId = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountAdminAccountName(v string) *ListFundAccountResponseBodyData {
	s.FundAccountAdminAccountName = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountId(v string) *ListFundAccountResponseBodyData {
	s.FundAccountId = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountName(v string) *ListFundAccountResponseBodyData {
	s.FundAccountName = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountOwnerAccountId(v string) *ListFundAccountResponseBodyData {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountStatus(v string) *ListFundAccountResponseBodyData {
	s.FundAccountStatus = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountType(v string) *ListFundAccountResponseBodyData {
	s.FundAccountType = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetNbid(v string) *ListFundAccountResponseBodyData {
	s.Nbid = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetPermissions(v []*string) *ListFundAccountResponseBodyData {
	s.Permissions = v
	return s
}

func (s *ListFundAccountResponseBodyData) SetSite(v string) *ListFundAccountResponseBodyData {
	s.Site = &v
	return s
}

func (s *ListFundAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
