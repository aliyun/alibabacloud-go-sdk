// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterKubeconfigStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v *ListClusterKubeconfigStatesResponseBodyPage) *ListClusterKubeconfigStatesResponseBody
	GetPage() *ListClusterKubeconfigStatesResponseBodyPage
	SetStates(v []*ListClusterKubeconfigStatesResponseBodyStates) *ListClusterKubeconfigStatesResponseBody
	GetStates() []*ListClusterKubeconfigStatesResponseBodyStates
}

type ListClusterKubeconfigStatesResponseBody struct {
	// The pagination information.
	Page *ListClusterKubeconfigStatesResponseBodyPage `json:"page,omitempty" xml:"page,omitempty" type:"Struct"`
	// The status list of the kubeconfig files associated with the cluster.
	States []*ListClusterKubeconfigStatesResponseBodyStates `json:"states,omitempty" xml:"states,omitempty" type:"Repeated"`
}

func (s ListClusterKubeconfigStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesResponseBody) GetPage() *ListClusterKubeconfigStatesResponseBodyPage {
	return s.Page
}

func (s *ListClusterKubeconfigStatesResponseBody) GetStates() []*ListClusterKubeconfigStatesResponseBodyStates {
	return s.States
}

func (s *ListClusterKubeconfigStatesResponseBody) SetPage(v *ListClusterKubeconfigStatesResponseBodyPage) *ListClusterKubeconfigStatesResponseBody {
	s.Page = v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBody) SetStates(v []*ListClusterKubeconfigStatesResponseBodyStates) *ListClusterKubeconfigStatesResponseBody {
	s.States = v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.States != nil {
		for _, item := range s.States {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterKubeconfigStatesResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s ListClusterKubeconfigStatesResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) SetPageNumber(v int32) *ListClusterKubeconfigStatesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) SetPageSize(v int32) *ListClusterKubeconfigStatesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) SetTotalCount(v int32) *ListClusterKubeconfigStatesResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type ListClusterKubeconfigStatesResponseBodyStates struct {
	// The displayed name or role name of the RAM user.
	//
	// example:
	//
	// tom
	AccountDisplayName *string `json:"account_display_name,omitempty" xml:"account_display_name,omitempty"`
	// The ID of an Alibaba Cloud account, RAM user, or RAM role.
	//
	// example:
	//
	// 22855*****************
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// The logon name or role name of the RAM user.
	//
	// example:
	//
	// tom
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// The status of the account.
	//
	// 	- Active: The account is active.
	//
	// 	- InActive: The account is locked.
	//
	// 	- Deleted: The account is deleted.
	//
	// example:
	//
	// Active
	AccountState *string `json:"account_state,omitempty" xml:"account_state,omitempty"`
	// The type of the account.
	//
	// 	- RootAccount: Alibaba Cloud account.
	//
	// 	- RamUser: RAM user.
	//
	// 	- RamRole: RAM role.
	//
	// example:
	//
	// RamUser
	AccountType *string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// The expiration time of the client certificate for the kubeconfig file.
	//
	// example:
	//
	// 2027-07-15T01:32:20Z
	CertExpireTime *string `json:"cert_expire_time,omitempty" xml:"cert_expire_time,omitempty"`
	// The status of the client certificate for the kubeconfig file.
	//
	// 	- Unexpired: The certificate is not expired.
	//
	// 	- Expired: The certificate is expired.
	//
	// 	- Unknown: The status of the certificate is unknown.
	//
	// example:
	//
	// Expired
	CertState *string `json:"cert_state,omitempty" xml:"cert_state,omitempty"`
	// Indicates whether the client certificate for the kubeconfig file can be revoked.
	//
	// example:
	//
	// true
	Revokable *bool `json:"revokable,omitempty" xml:"revokable,omitempty"`
}

func (s ListClusterKubeconfigStatesResponseBodyStates) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesResponseBodyStates) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountDisplayName() *string {
	return s.AccountDisplayName
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountId() *string {
	return s.AccountId
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountName() *string {
	return s.AccountName
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountState() *string {
	return s.AccountState
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountType() *string {
	return s.AccountType
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetCertState() *string {
	return s.CertState
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetRevokable() *bool {
	return s.Revokable
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountDisplayName(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountDisplayName = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountId(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountId = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountName(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountName = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountState(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountState = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountType(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountType = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetCertExpireTime(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.CertExpireTime = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetCertState(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.CertState = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetRevokable(v bool) *ListClusterKubeconfigStatesResponseBodyStates {
	s.Revokable = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) Validate() error {
	return dara.Validate(s)
}
