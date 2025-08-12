// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMetricRuleResourcesResponseBody
	GetCode() *string
	SetMessage(v string) *CreateMetricRuleResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMetricRuleResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMetricRuleResourcesResponseBody
	GetSuccess() *bool
}

type CreateMetricRuleResourcesResponseBody struct {
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
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
	// The request ID.
	//
	// example:
	//
	// 0671A721-0D7A-4F11-BB77-2416325D65AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMetricRuleResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMetricRuleResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMetricRuleResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMetricRuleResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMetricRuleResourcesResponseBody) SetCode(v string) *CreateMetricRuleResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMetricRuleResourcesResponseBody) SetMessage(v string) *CreateMetricRuleResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMetricRuleResourcesResponseBody) SetRequestId(v string) *CreateMetricRuleResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetricRuleResourcesResponseBody) SetSuccess(v bool) *CreateMetricRuleResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMetricRuleResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
