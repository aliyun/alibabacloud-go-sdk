// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityRuleSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateQualityRuleSwitchResponseBody
	GetCode() *string
	SetData(v int32) *UpdateQualityRuleSwitchResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *UpdateQualityRuleSwitchResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateQualityRuleSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateQualityRuleSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateQualityRuleSwitchResponseBody
	GetSuccess() *bool
}

type UpdateQualityRuleSwitchResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityRuleSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityRuleSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityRuleSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateQualityRuleSwitchResponseBody) GetData() *int32 {
	return s.Data
}

func (s *UpdateQualityRuleSwitchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateQualityRuleSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateQualityRuleSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQualityRuleSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateQualityRuleSwitchResponseBody) SetCode(v string) *UpdateQualityRuleSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQualityRuleSwitchResponseBody) SetData(v int32) *UpdateQualityRuleSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQualityRuleSwitchResponseBody) SetHttpStatusCode(v int32) *UpdateQualityRuleSwitchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateQualityRuleSwitchResponseBody) SetMessage(v string) *UpdateQualityRuleSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQualityRuleSwitchResponseBody) SetRequestId(v string) *UpdateQualityRuleSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityRuleSwitchResponseBody) SetSuccess(v bool) *UpdateQualityRuleSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQualityRuleSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
