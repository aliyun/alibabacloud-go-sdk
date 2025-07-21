// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountKmsStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountStatus(v string) *DescribeAccountKmsStatusResponseBody
	GetAccountStatus() *string
	SetRequestId(v string) *DescribeAccountKmsStatusResponseBody
	GetRequestId() *string
}

type DescribeAccountKmsStatusResponseBody struct {
	// The status of KMS within your Alibaba cloud account. Valid values:
	//
	// 	- Enabled: KMS is enabled.
	//
	// 	- NotEnabled: KMS is disabled.
	//
	// 	- InDebt: Your account is overdue, and KMS stops providing services.
	//
	// > If your Alibaba Cloud account is overdue, top up your account at the earliest opportunity to avoid impacts on your services.
	//
	// 	- Suspended: KMS is suspended.
	//
	// example:
	//
	// Enabled
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3ac84333-d64d-4784-a8bc-997834a7ac6c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountKmsStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountKmsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountKmsStatusResponseBody) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeAccountKmsStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountKmsStatusResponseBody) SetAccountStatus(v string) *DescribeAccountKmsStatusResponseBody {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountKmsStatusResponseBody) SetRequestId(v string) *DescribeAccountKmsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountKmsStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
