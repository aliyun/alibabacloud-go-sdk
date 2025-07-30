// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAssignRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTaskAssignRuleResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateTaskAssignRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTaskAssignRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskAssignRuleResponseBody
	GetSuccess() *bool
}

type UpdateTaskAssignRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskAssignRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTaskAssignRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTaskAssignRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskAssignRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskAssignRuleResponseBody) SetCode(v string) *UpdateTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetMessage(v string) *UpdateTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetRequestId(v string) *UpdateTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetSuccess(v bool) *UpdateTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
