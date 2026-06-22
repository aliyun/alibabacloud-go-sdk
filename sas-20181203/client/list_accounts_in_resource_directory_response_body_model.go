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
	// The list of member accounts in the resource directory.
	Accounts []*ListAccountsInResourceDirectoryResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
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
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccountsInResourceDirectoryResponseBodyAccounts struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 180803538814****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated and does not need to be used.
	//
	// example:
	//
	// true
	BuySas *int32 `json:"BuySas,omitempty" xml:"BuySas,omitempty"`
	// Indicates whether a Security Center subscription instance is purchased. Valid values:
	//
	// - **true**
	//
	// - **false**.
	//
	// example:
	//
	// true
	BuySasNew *bool `json:"BuySasNew,omitempty" xml:"BuySasNew,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **0**: upfront
	//
	// 	- **1**: pay-as-you-go.
	//
	// example:
	//
	// 1
	ChargeType *int32 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The account name.
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
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The instance purchase type. Valid values:
	//
	// - **0**: self-purchased
	//
	// - **1**: allocated through multi-account management.
	//
	// example:
	//
	// 0
	InstanceBuyType *int32 `json:"InstanceBuyType,omitempty" xml:"InstanceBuyType,omitempty"`
	// Indicates whether the account is the management account of the resource directory. Valid values:
	//
	// - **yes**
	//
	// - **no**.
	//
	// example:
	//
	// no
	IsMaAccount *string `json:"IsMaAccount,omitempty" xml:"IsMaAccount,omitempty"`
	// Indicates whether the account is marked as followed.
	//
	// example:
	//
	// no
	IsMarked *string `json:"IsMarked,omitempty" xml:"IsMarked,omitempty"`
	// Indicates whether the account is a delegated administrator account of Security Center. Valid values:
	//
	// - **yes**
	//
	// - **no**.
	//
	// example:
	//
	// no
	IsSasDaAccount *string `json:"IsSasDaAccount,omitempty" xml:"IsSasDaAccount,omitempty"`
	// Indicates whether the account is a management account of Cloud Threat Detection and Response (CTDR). Valid values:
	//
	// - **yes**
	//
	// - **no**.
	//
	// example:
	//
	// no
	IsSiemControlAccount *string `json:"IsSiemControlAccount,omitempty" xml:"IsSiemControlAccount,omitempty"`
	// Indicates whether the account is a delegated administrator account of Cloud Threat Detection and Response (CTDR). Valid values:
	//
	// - **yes**
	//
	// - **no**.
	//
	// example:
	//
	// no
	IsSiemDaAccount *string `json:"IsSiemDaAccount,omitempty" xml:"IsSiemDaAccount,omitempty"`
	// The pay-as-you-go module switch. Valid values:
	//
	// - **0**: disabled
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 1
	PostBasicService *int32 `json:"PostBasicService,omitempty" xml:"PostBasicService,omitempty"`
	// The status of pay-as-you-go module switches, in JsonString format. Valid values:
	//
	// - Key:
	//
	//   - **VUL**: vulnerability management module
	//
	//   - **CSPM**: Cloud Security Posture Management (CSPM) module
	//
	//   - **AGENTLESS**: agentless detection module
	//
	//   - **SERVERLESS**: serverless asset module
	//
	//   - **CTDR**: Cloud Threat Detection and Response (CTDR) module
	//
	//   - **RASP**: Runtime Application Self-Protection (RASP) module
	//
	//   - **SDK**: malicious file detection SDK module
	//
	//   - **POST_HOST**: host and container security module
	//
	// - Value: 0 indicates disabled. 1 indicates enabled.
	//
	// example:
	//
	// {\\"VUL\\":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// The Security Center instance ID.
	//
	// example:
	//
	// postpay-sas-x5x3hy1ly***
	SaleInstance *string `json:"SaleInstance,omitempty" xml:"SaleInstance,omitempty"`
	// The purchased edition of Security Center. Valid values:
	//
	// - **0*	- or **1**: Free Edition
	//
	// - **2*	- or **3**: Enterprise Edition
	//
	// - **5**: Premium Edition
	//
	// - **6**: Anti-virus Edition
	//
	// - **7**: Ultimate Edition.
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
