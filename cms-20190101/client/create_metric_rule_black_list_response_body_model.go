// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleBlackListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMetricRuleBlackListResponseBody
	GetCode() *string
	SetId(v string) *CreateMetricRuleBlackListResponseBody
	GetId() *string
	SetMessage(v string) *CreateMetricRuleBlackListResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMetricRuleBlackListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMetricRuleBlackListResponseBody
	GetSuccess() *bool
}

type CreateMetricRuleBlackListResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the blacklist policy.
	//
	// example:
	//
	// 93514c96-ceb8-47d8-8ee3-93b6d98b****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message.
	//
	// example:
	//
	// The request has failed due to a temporary failure of the server.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5C637FA3-3959-3352-8BE6-78CE2ED2D33D
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

func (s CreateMetricRuleBlackListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleBlackListResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMetricRuleBlackListResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateMetricRuleBlackListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMetricRuleBlackListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMetricRuleBlackListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMetricRuleBlackListResponseBody) SetCode(v string) *CreateMetricRuleBlackListResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMetricRuleBlackListResponseBody) SetId(v string) *CreateMetricRuleBlackListResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMetricRuleBlackListResponseBody) SetMessage(v string) *CreateMetricRuleBlackListResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMetricRuleBlackListResponseBody) SetRequestId(v string) *CreateMetricRuleBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetricRuleBlackListResponseBody) SetSuccess(v bool) *CreateMetricRuleBlackListResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMetricRuleBlackListResponseBody) Validate() error {
	return dara.Validate(s)
}
