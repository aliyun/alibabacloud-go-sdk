// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteMetricRuleTemplateResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteMetricRuleTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMetricRuleTemplateResponseBody
	GetRequestId() *string
	SetResource(v *DeleteMetricRuleTemplateResponseBodyResource) *DeleteMetricRuleTemplateResponseBody
	GetResource() *DeleteMetricRuleTemplateResponseBodyResource
	SetSuccess(v bool) *DeleteMetricRuleTemplateResponseBody
	GetSuccess() *bool
}

type DeleteMetricRuleTemplateResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B63F8CF-D48D-4608-A402-04FB5B2B4B6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template information.
	Resource *DeleteMetricRuleTemplateResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetricRuleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteMetricRuleTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMetricRuleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetricRuleTemplateResponseBody) GetResource() *DeleteMetricRuleTemplateResponseBodyResource {
	return s.Resource
}

func (s *DeleteMetricRuleTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetricRuleTemplateResponseBody) SetCode(v int32) *DeleteMetricRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) SetMessage(v string) *DeleteMetricRuleTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) SetRequestId(v string) *DeleteMetricRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) SetResource(v *DeleteMetricRuleTemplateResponseBodyResource) *DeleteMetricRuleTemplateResponseBody {
	s.Resource = v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) SetSuccess(v bool) *DeleteMetricRuleTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteMetricRuleTemplateResponseBodyResource struct {
	// The ID of the template.
	//
	// example:
	//
	// 123****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMetricRuleTemplateResponseBodyResource) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTemplateResponseBodyResource) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTemplateResponseBodyResource) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteMetricRuleTemplateResponseBodyResource) SetTemplateId(v string) *DeleteMetricRuleTemplateResponseBodyResource {
	s.TemplateId = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBodyResource) Validate() error {
	return dara.Validate(s)
}
