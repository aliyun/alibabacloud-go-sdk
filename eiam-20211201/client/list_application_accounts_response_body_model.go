// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccounts(v []*ListApplicationAccountsResponseBodyApplicationAccounts) *ListApplicationAccountsResponseBody
	GetApplicationAccounts() []*ListApplicationAccountsResponseBodyApplicationAccounts
	SetRequestId(v string) *ListApplicationAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationAccountsResponseBody
	GetTotalCount() *int64
}

type ListApplicationAccountsResponseBody struct {
	// The list of account information.
	ApplicationAccounts []*ListApplicationAccountsResponseBodyApplicationAccounts `json:"ApplicationAccounts,omitempty" xml:"ApplicationAccounts,omitempty" type:"Repeated"`
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

func (s ListApplicationAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationAccountsResponseBody) GetApplicationAccounts() []*ListApplicationAccountsResponseBodyApplicationAccounts {
	return s.ApplicationAccounts
}

func (s *ListApplicationAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationAccountsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationAccountsResponseBody) SetApplicationAccounts(v []*ListApplicationAccountsResponseBodyApplicationAccounts) *ListApplicationAccountsResponseBody {
	s.ApplicationAccounts = v
	return s
}

func (s *ListApplicationAccountsResponseBody) SetRequestId(v string) *ListApplicationAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationAccountsResponseBody) SetTotalCount(v int64) *ListApplicationAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationAccountsResponseBody) Validate() error {
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

type ListApplicationAccountsResponseBodyApplicationAccounts struct {
	// The application account ID.
	//
	// example:
	//
	// aac_m66smhbq5krept7nza54hxxxx
	ApplicationAccountId *string `json:"ApplicationAccountId,omitempty" xml:"ApplicationAccountId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// app_m43o4h5adeo5klvbbxgxxxxx
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
	// 1737510353000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// idaas_eznwtkkaucljizh6qqu7ptxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// user_tkmboufpnvpbitdpzrlngxzxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListApplicationAccountsResponseBodyApplicationAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccountsResponseBodyApplicationAccounts) GoString() string {
	return s.String()
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) GetApplicationAccountId() *string {
	return s.ApplicationAccountId
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) GetApplicationUsername() *string {
	return s.ApplicationUsername
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) GetUserId() *string {
	return s.UserId
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) SetApplicationAccountId(v string) *ListApplicationAccountsResponseBodyApplicationAccounts {
	s.ApplicationAccountId = &v
	return s
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) SetApplicationId(v string) *ListApplicationAccountsResponseBodyApplicationAccounts {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) SetApplicationUsername(v string) *ListApplicationAccountsResponseBodyApplicationAccounts {
	s.ApplicationUsername = &v
	return s
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) SetCreateTime(v int64) *ListApplicationAccountsResponseBodyApplicationAccounts {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) SetInstanceId(v string) *ListApplicationAccountsResponseBodyApplicationAccounts {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) SetUserId(v string) *ListApplicationAccountsResponseBodyApplicationAccounts {
	s.UserId = &v
	return s
}

func (s *ListApplicationAccountsResponseBodyApplicationAccounts) Validate() error {
	return dara.Validate(s)
}
