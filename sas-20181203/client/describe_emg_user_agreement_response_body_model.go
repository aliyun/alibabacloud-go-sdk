// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmgUserAgreementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v bool) *DescribeEmgUserAgreementResponseBody
	GetAuth() *bool
	SetRequestId(v string) *DescribeEmgUserAgreementResponseBody
	GetRequestId() *string
}

type DescribeEmgUserAgreementResponseBody struct {
	// Indicates whether Security Center is authorized to scan for urgent vulnerabilities. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Auth *bool `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 768BB9F5-8AF3-557F-A489-9BDD64CB3E4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEmgUserAgreementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmgUserAgreementResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmgUserAgreementResponseBody) GetAuth() *bool {
	return s.Auth
}

func (s *DescribeEmgUserAgreementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEmgUserAgreementResponseBody) SetAuth(v bool) *DescribeEmgUserAgreementResponseBody {
	s.Auth = &v
	return s
}

func (s *DescribeEmgUserAgreementResponseBody) SetRequestId(v string) *DescribeEmgUserAgreementResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEmgUserAgreementResponseBody) Validate() error {
	return dara.Validate(s)
}
