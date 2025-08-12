// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomEventRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutCustomEventRuleResponseBody
	GetCode() *string
	SetMessage(v string) *PutCustomEventRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutCustomEventRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutCustomEventRuleResponseBody
	GetSuccess() *bool
}

type PutCustomEventRuleResponseBody struct {
	// The HTTP status code.
	//
	// >  The value 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The request has failed due to a temporary failure of the server.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD5DCD82-BD1C-405F-BAED-32302DE9F498
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutCustomEventRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutCustomEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomEventRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutCustomEventRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutCustomEventRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutCustomEventRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutCustomEventRuleResponseBody) SetCode(v string) *PutCustomEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutCustomEventRuleResponseBody) SetMessage(v string) *PutCustomEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomEventRuleResponseBody) SetRequestId(v string) *PutCustomEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCustomEventRuleResponseBody) SetSuccess(v bool) *PutCustomEventRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PutCustomEventRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
