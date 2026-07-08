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
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 10***********34
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The Alibaba Cloud account name. This parameter is returned only when the account is a delegated administrator.
	//
	// example:
	//
	// account_test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the user is a delegated administrator of WAF. Valid values:
	//
	// - **true**: The user is a delegated administrator of WAF.
	//
	// - **false**: The user is not a delegated administrator of WAF.
	//
	// example:
	//
	// true
	DelegatedStatus *bool `json:"DelegatedStatus,omitempty" xml:"DelegatedStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8161375D-5958-5627-BFDE-DF1458A73E87
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
