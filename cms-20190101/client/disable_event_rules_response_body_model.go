// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEventRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableEventRulesResponseBody
	GetCode() *string
	SetMessage(v string) *DisableEventRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableEventRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableEventRulesResponseBody
	GetSuccess() *bool
}

type DisableEventRulesResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3FD0E8B5-F132-4F4E-A081-2878AF378B12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. The value true indicates a success. The value false indicates a failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableEventRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableEventRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableEventRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableEventRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableEventRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableEventRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableEventRulesResponseBody) SetCode(v string) *DisableEventRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableEventRulesResponseBody) SetMessage(v string) *DisableEventRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableEventRulesResponseBody) SetRequestId(v string) *DisableEventRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableEventRulesResponseBody) SetSuccess(v bool) *DisableEventRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DisableEventRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
