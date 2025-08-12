// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMetricRuleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyMetricRuleTemplateResponseBody
	GetCode() *int32
	SetMessage(v string) *ModifyMetricRuleTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMetricRuleTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMetricRuleTemplateResponseBody
	GetSuccess() *bool
}

type ModifyMetricRuleTemplateResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E07117F-F6AE-4F1C-81E8-36FBB4892235
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMetricRuleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyMetricRuleTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMetricRuleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMetricRuleTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMetricRuleTemplateResponseBody) SetCode(v int32) *ModifyMetricRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMetricRuleTemplateResponseBody) SetMessage(v string) *ModifyMetricRuleTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMetricRuleTemplateResponseBody) SetRequestId(v string) *ModifyMetricRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMetricRuleTemplateResponseBody) SetSuccess(v bool) *ModifyMetricRuleTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMetricRuleTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
