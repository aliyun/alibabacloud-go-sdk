// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRuleResponseBody
	GetCode() *string
	SetData(v bool) *UpdateRuleResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRuleResponseBody
	GetSuccess() *bool
}

type UpdateRuleResponseBody struct {
	// The response code. The value Success indicates that the request was successful. Other values indicate that the request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// EventRuleFilterPatternSchemaError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7439944D-D9FC-5614-B65E-07A468988924
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRuleResponseBody) SetCode(v string) *UpdateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleResponseBody) SetData(v bool) *UpdateRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleResponseBody) SetMessage(v string) *UpdateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleResponseBody) SetRequestId(v string) *UpdateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleResponseBody) SetSuccess(v bool) *UpdateRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
