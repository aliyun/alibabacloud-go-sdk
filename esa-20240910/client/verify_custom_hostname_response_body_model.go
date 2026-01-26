// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCustomHostnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPassed(v bool) *VerifyCustomHostnameResponseBody
	GetPassed() *bool
	SetRequestId(v string) *VerifyCustomHostnameResponseBody
	GetRequestId() *string
}

type VerifyCustomHostnameResponseBody struct {
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyCustomHostnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyCustomHostnameResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCustomHostnameResponseBody) GetPassed() *bool {
	return s.Passed
}

func (s *VerifyCustomHostnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyCustomHostnameResponseBody) SetPassed(v bool) *VerifyCustomHostnameResponseBody {
	s.Passed = &v
	return s
}

func (s *VerifyCustomHostnameResponseBody) SetRequestId(v string) *VerifyCustomHostnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyCustomHostnameResponseBody) Validate() error {
	return dara.Validate(s)
}
