// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccountsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccounts(v []*ListApplicationAccountsForUserResponseBodyApplicationAccounts) *ListApplicationAccountsForUserResponseBody
	GetApplicationAccounts() []*ListApplicationAccountsForUserResponseBodyApplicationAccounts
	SetRequestId(v string) *ListApplicationAccountsForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationAccountsForUserResponseBody
	GetTotalCount() *int64
}

type ListApplicationAccountsForUserResponseBody struct {
	// The information about the member accounts.
	ApplicationAccounts []*ListApplicationAccountsForUserResponseBodyApplicationAccounts `json:"ApplicationAccounts,omitempty" xml:"ApplicationAccounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationAccountsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccountsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationAccountsForUserResponseBody) GetApplicationAccounts() []*ListApplicationAccountsForUserResponseBodyApplicationAccounts {
	return s.ApplicationAccounts
}

func (s *ListApplicationAccountsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationAccountsForUserResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationAccountsForUserResponseBody) SetApplicationAccounts(v []*ListApplicationAccountsForUserResponseBodyApplicationAccounts) *ListApplicationAccountsForUserResponseBody {
	s.ApplicationAccounts = v
	return s
}

func (s *ListApplicationAccountsForUserResponseBody) SetRequestId(v string) *ListApplicationAccountsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationAccountsForUserResponseBody) SetTotalCount(v int64) *ListApplicationAccountsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationAccountsForUserResponseBody) Validate() error {
	if s.ApplicationAccounts != nil {
		for _, item := range s.ApplicationAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationAccountsForUserResponseBodyApplicationAccounts struct {
	// The application account ID.
	//
	// example:
	//
	// aac_m6e3ukegwvbcb2fne7j32xxxxxx
	ApplicationAccountId *string `json:"ApplicationAccountId,omitempty" xml:"ApplicationAccountId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// app_na2r76irswrwfgpkz7xvcj7xxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the application account.
	//
	// example:
	//
	// test
	ApplicationUsername *string `json:"ApplicationUsername,omitempty" xml:"ApplicationUsername,omitempty"`
	// The time when the account was created.
	//
	// example:
	//
	// 1754359439000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// idaas_ki6hd7ihir4ybawogqk6xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// user_tkmboufpnvpbitdpzrlng6mxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListApplicationAccountsForUserResponseBodyApplicationAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccountsForUserResponseBodyApplicationAccounts) GoString() string {
	return s.String()
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) GetApplicationAccountId() *string {
	return s.ApplicationAccountId
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) GetApplicationUsername() *string {
	return s.ApplicationUsername
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) GetUserId() *string {
	return s.UserId
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) SetApplicationAccountId(v string) *ListApplicationAccountsForUserResponseBodyApplicationAccounts {
	s.ApplicationAccountId = &v
	return s
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) SetApplicationId(v string) *ListApplicationAccountsForUserResponseBodyApplicationAccounts {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) SetApplicationUsername(v string) *ListApplicationAccountsForUserResponseBodyApplicationAccounts {
	s.ApplicationUsername = &v
	return s
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) SetCreateTime(v int64) *ListApplicationAccountsForUserResponseBodyApplicationAccounts {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) SetInstanceId(v string) *ListApplicationAccountsForUserResponseBodyApplicationAccounts {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) SetUserId(v string) *ListApplicationAccountsForUserResponseBodyApplicationAccounts {
	s.UserId = &v
	return s
}

func (s *ListApplicationAccountsForUserResponseBodyApplicationAccounts) Validate() error {
	return dara.Validate(s)
}
