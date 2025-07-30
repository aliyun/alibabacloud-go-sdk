// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskAssignRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTaskAssignRuleResponseBody
	GetCode() *string
	SetData(v string) *CreateTaskAssignRuleResponseBody
	GetData() *string
	SetMessage(v string) *CreateTaskAssignRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTaskAssignRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTaskAssignRuleResponseBody
	GetSuccess() *bool
}

type CreateTaskAssignRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 54
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateTaskAssignRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTaskAssignRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateTaskAssignRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTaskAssignRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskAssignRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTaskAssignRuleResponseBody) SetCode(v string) *CreateTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetData(v string) *CreateTaskAssignRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetMessage(v string) *CreateTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetRequestId(v string) *CreateTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetSuccess(v bool) *CreateTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
