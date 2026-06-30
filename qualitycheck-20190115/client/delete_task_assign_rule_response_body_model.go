// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskAssignRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTaskAssignRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteTaskAssignRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTaskAssignRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTaskAssignRuleResponseBody
	GetSuccess() *bool
}

type DeleteTaskAssignRuleResponseBody struct {
	// The result code. **200*	- means the operation succeeded. Any other value means the operation failed. Use this field to identify the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error details if the operation failed. If the operation succeeded, the value is successful.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. A value of true means success. A value of false or null means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTaskAssignRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTaskAssignRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTaskAssignRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTaskAssignRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTaskAssignRuleResponseBody) SetCode(v string) *DeleteTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetMessage(v string) *DeleteTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetRequestId(v string) *DeleteTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetSuccess(v bool) *DeleteTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
