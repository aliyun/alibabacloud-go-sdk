// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutEventRuleResponseBody
	GetCode() *string
	SetData(v string) *PutEventRuleResponseBody
	GetData() *string
	SetMessage(v string) *PutEventRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutEventRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutEventRuleResponseBody
	GetSuccess() *bool
}

type PutEventRuleResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of event-triggered alert rules that were created or modified.
	//
	// example:
	//
	// 1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0B47C47B-E68A-4429-BB23-370E91889C7D
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

func (s PutEventRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutEventRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutEventRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *PutEventRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutEventRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutEventRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutEventRuleResponseBody) SetCode(v string) *PutEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutEventRuleResponseBody) SetData(v string) *PutEventRuleResponseBody {
	s.Data = &v
	return s
}

func (s *PutEventRuleResponseBody) SetMessage(v string) *PutEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutEventRuleResponseBody) SetRequestId(v string) *PutEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEventRuleResponseBody) SetSuccess(v bool) *PutEventRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PutEventRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
