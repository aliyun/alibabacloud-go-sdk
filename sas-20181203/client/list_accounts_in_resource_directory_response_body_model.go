// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsInResourceDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v []*ListAccountsInResourceDirectoryResponseBodyAccounts) *ListAccountsInResourceDirectoryResponseBody
	GetAccounts() []*ListAccountsInResourceDirectoryResponseBodyAccounts
	SetRequestId(v string) *ListAccountsInResourceDirectoryResponseBody
	GetRequestId() *string
}

type ListAccountsInResourceDirectoryResponseBody struct {
	// The members in the resource directory.
	Accounts []*ListAccountsInResourceDirectoryResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B7A839E8-70AE-591D-8D9E-C5419A2240DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccountsInResourceDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsInResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsInResourceDirectoryResponseBody) GetAccounts() []*ListAccountsInResourceDirectoryResponseBodyAccounts {
	return s.Accounts
}

func (s *ListAccountsInResourceDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountsInResourceDirectoryResponseBody) SetAccounts(v []*ListAccountsInResourceDirectoryResponseBodyAccounts) *ListAccountsInResourceDirectoryResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBody) SetRequestId(v string) *ListAccountsInResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAccountsInResourceDirectoryResponseBodyAccounts struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 180803538814****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Deprecated
	BuySas     *int32 `json:"BuySas,omitempty" xml:"BuySas,omitempty"`
	BuySasNew  *bool  `json:"BuySasNew,omitempty" xml:"BuySasNew,omitempty"`
	ChargeType *int32 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The name of the member.
	//
	// example:
	//
	// abc**
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the folder in the resource directory.
	//
	// example:
	//
	// fd-cE2SQP****
	FolderId        *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	InstanceBuyType *int32  `json:"InstanceBuyType,omitempty" xml:"InstanceBuyType,omitempty"`
	// Indicates whether the member is an administrator account of the resource directory. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// no
	IsMaAccount *string `json:"IsMaAccount,omitempty" xml:"IsMaAccount,omitempty"`
	// Indicates whether the member is marked as followed.
	//
	// example:
	//
	// no
	IsMarked *string `json:"IsMarked,omitempty" xml:"IsMarked,omitempty"`
	// Indicates whether the member is a delegated administrator account of Security Center. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// no
	IsSasDaAccount *string `json:"IsSasDaAccount,omitempty" xml:"IsSasDaAccount,omitempty"`
	// Indicates whether the member is an account of the threat analysis and response feature. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// no
	IsSiemControlAccount *string `json:"IsSiemControlAccount,omitempty" xml:"IsSiemControlAccount,omitempty"`
	// Indicates whether the member is a delegated administrator account of the threat analysis and response feature. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// no
	IsSiemDaAccount     *string `json:"IsSiemDaAccount,omitempty" xml:"IsSiemDaAccount,omitempty"`
	PostBasicService    *int32  `json:"PostBasicService,omitempty" xml:"PostBasicService,omitempty"`
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	SaleInstance        *string `json:"SaleInstance,omitempty" xml:"SaleInstance,omitempty"`
	// The edition of Security Center that you use. Valid value:
	//
	// 	- **0*	- or **1**: Basic
	//
	// 	- **2*	- or **3**: Enterprise
	//
	// 	- **5**: Advanced
	//
	// 	- **6**: Anti-virus
	//
	// 	- **7**: Ultimate
	//
	// example:
	//
	// 0
	SasVersion *string `json:"SasVersion,omitempty" xml:"SasVersion,omitempty"`
}

func (s ListAccountsInResourceDirectoryResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsInResourceDirectoryResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetBuySas() *int32 {
	return s.BuySas
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetBuySasNew() *bool {
	return s.BuySasNew
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetChargeType() *int32 {
	return s.ChargeType
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetFolderId() *string {
	return s.FolderId
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetInstanceBuyType() *int32 {
	return s.InstanceBuyType
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsMaAccount() *string {
	return s.IsMaAccount
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsMarked() *string {
	return s.IsMarked
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsSasDaAccount() *string {
	return s.IsSasDaAccount
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsSiemControlAccount() *string {
	return s.IsSiemControlAccount
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsSiemDaAccount() *string {
	return s.IsSiemDaAccount
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetPostBasicService() *int32 {
	return s.PostBasicService
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetPostPayModuleSwitch() *string {
	return s.PostPayModuleSwitch
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetSaleInstance() *string {
	return s.SaleInstance
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetSasVersion() *string {
	return s.SasVersion
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetAccountId(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.AccountId = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetBuySas(v int32) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.BuySas = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetBuySasNew(v bool) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.BuySasNew = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetChargeType(v int32) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.ChargeType = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetDisplayName(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetFolderId(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.FolderId = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetInstanceBuyType(v int32) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.InstanceBuyType = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsMaAccount(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsMaAccount = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsMarked(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsMarked = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsSasDaAccount(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsSasDaAccount = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsSiemControlAccount(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsSiemControlAccount = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsSiemDaAccount(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsSiemDaAccount = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetPostBasicService(v int32) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.PostBasicService = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetPostPayModuleSwitch(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.PostPayModuleSwitch = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetSaleInstance(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.SaleInstance = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetSasVersion(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.SasVersion = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) Validate() error {
	return dara.Validate(s)
}
