// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocalityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateLocalityRuleResponseBody
	GetCode() *int32
	SetData(v string) *UpdateLocalityRuleResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateLocalityRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateLocalityRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLocalityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateLocalityRuleResponseBody
	GetSuccess() *string
}

type UpdateLocalityRuleResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//   "Message":"updateLocalityPolicy success",
	//
	//   "RequestId":"3B519913-7348-16AB-AD71-******",
	//
	//   "HttpStatusCode":200,
	//
	//   "Code":200,
	//
	//   "Success":true
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 78F05E89-D387-50CE-8186-2E27A8AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLocalityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocalityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLocalityRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateLocalityRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateLocalityRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateLocalityRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLocalityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLocalityRuleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateLocalityRuleResponseBody) SetCode(v int32) *UpdateLocalityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLocalityRuleResponseBody) SetData(v string) *UpdateLocalityRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLocalityRuleResponseBody) SetHttpStatusCode(v int32) *UpdateLocalityRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLocalityRuleResponseBody) SetMessage(v string) *UpdateLocalityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLocalityRuleResponseBody) SetRequestId(v string) *UpdateLocalityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLocalityRuleResponseBody) SetSuccess(v string) *UpdateLocalityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLocalityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
