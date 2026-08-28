// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAuthorizationRuleResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAuthorizationRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAuthorizationRuleResponseBody
	GetRequestId() *string
}

type UpdateAuthorizationRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 019FCFBD-4B8D-57A4-91DA-A3B517002442
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAuthorizationRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthorizationRuleResponseBody) SetCode(v string) *UpdateAuthorizationRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAuthorizationRuleResponseBody) SetMessage(v string) *UpdateAuthorizationRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAuthorizationRuleResponseBody) SetRequestId(v string) *UpdateAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
