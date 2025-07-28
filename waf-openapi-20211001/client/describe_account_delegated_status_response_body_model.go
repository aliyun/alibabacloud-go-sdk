// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountDelegatedStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeAccountDelegatedStatusResponseBody
	GetAccountId() *string
	SetAccountName(v string) *DescribeAccountDelegatedStatusResponseBody
	GetAccountName() *string
	SetDelegatedStatus(v bool) *DescribeAccountDelegatedStatusResponseBody
	GetDelegatedStatus() *bool
	SetRequestId(v string) *DescribeAccountDelegatedStatusResponseBody
	GetRequestId() *string
}

type DescribeAccountDelegatedStatusResponseBody struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 10***********34
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account. This parameter is returned only if the account is the delegated administrator account.
	//
	// example:
	//
	// account_test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the Alibaba Cloud account is the delegated administrator account of the WAF instance.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DelegatedStatus *bool `json:"DelegatedStatus,omitempty" xml:"DelegatedStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8161375D-5958-5627-BFDE-DF14****3E87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountDelegatedStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountDelegatedStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountDelegatedStatusResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeAccountDelegatedStatusResponseBody) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountDelegatedStatusResponseBody) GetDelegatedStatus() *bool {
	return s.DelegatedStatus
}

func (s *DescribeAccountDelegatedStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountDelegatedStatusResponseBody) SetAccountId(v string) *DescribeAccountDelegatedStatusResponseBody {
	s.AccountId = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponseBody) SetAccountName(v string) *DescribeAccountDelegatedStatusResponseBody {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponseBody) SetDelegatedStatus(v bool) *DescribeAccountDelegatedStatusResponseBody {
	s.DelegatedStatus = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponseBody) SetRequestId(v string) *DescribeAccountDelegatedStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
