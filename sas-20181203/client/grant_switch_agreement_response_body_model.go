// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSwitchAgreementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GrantSwitchAgreementResponseBody
	GetCode() *string
	SetMessage(v string) *GrantSwitchAgreementResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrantSwitchAgreementResponseBody
	GetRequestId() *string
}

type GrantSwitchAgreementResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 11C96623-E106-59C9-866D-A6C82911459F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantSwitchAgreementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantSwitchAgreementResponseBody) GoString() string {
	return s.String()
}

func (s *GrantSwitchAgreementResponseBody) GetCode() *string {
	return s.Code
}

func (s *GrantSwitchAgreementResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrantSwitchAgreementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantSwitchAgreementResponseBody) SetCode(v string) *GrantSwitchAgreementResponseBody {
	s.Code = &v
	return s
}

func (s *GrantSwitchAgreementResponseBody) SetMessage(v string) *GrantSwitchAgreementResponseBody {
	s.Message = &v
	return s
}

func (s *GrantSwitchAgreementResponseBody) SetRequestId(v string) *GrantSwitchAgreementResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantSwitchAgreementResponseBody) Validate() error {
	return dara.Validate(s)
}
