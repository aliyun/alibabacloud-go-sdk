// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateMetricRuleTemplateResponseBody
	GetCode() *int32
	SetId(v int64) *CreateMetricRuleTemplateResponseBody
	GetId() *int64
	SetMessage(v string) *CreateMetricRuleTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMetricRuleTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMetricRuleTemplateResponseBody
	GetSuccess() *bool
}

type CreateMetricRuleTemplateResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the alert template.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// 9763ED1A-4D09-41BF-851E-310421750204
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

func (s CreateMetricRuleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateMetricRuleTemplateResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateMetricRuleTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMetricRuleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMetricRuleTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMetricRuleTemplateResponseBody) SetCode(v int32) *CreateMetricRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) SetId(v int64) *CreateMetricRuleTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) SetMessage(v string) *CreateMetricRuleTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) SetRequestId(v string) *CreateMetricRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) SetSuccess(v bool) *CreateMetricRuleTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
