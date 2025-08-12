// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMetricRuleResourcesResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMetricRuleResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMetricRuleResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMetricRuleResourcesResponseBody
	GetSuccess() *bool
}

type DeleteMetricRuleResourcesResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The alert does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D8A35882-90C6-4F03-BBEB-153C180398EA
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

func (s DeleteMetricRuleResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMetricRuleResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMetricRuleResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetricRuleResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetricRuleResourcesResponseBody) SetCode(v string) *DeleteMetricRuleResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRuleResourcesResponseBody) SetMessage(v string) *DeleteMetricRuleResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRuleResourcesResponseBody) SetRequestId(v string) *DeleteMetricRuleResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRuleResourcesResponseBody) SetSuccess(v bool) *DeleteMetricRuleResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetricRuleResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
