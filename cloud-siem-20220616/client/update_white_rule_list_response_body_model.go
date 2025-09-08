// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateWhiteRuleListResponseBody
	GetCode() *int32
	SetData(v interface{}) *UpdateWhiteRuleListResponseBody
	GetData() interface{}
	SetMessage(v string) *UpdateWhiteRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWhiteRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWhiteRuleListResponseBody
	GetSuccess() *bool
}

type UpdateWhiteRuleListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
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

func (s UpdateWhiteRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateWhiteRuleListResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateWhiteRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWhiteRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWhiteRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWhiteRuleListResponseBody) SetCode(v int32) *UpdateWhiteRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetData(v interface{}) *UpdateWhiteRuleListResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetMessage(v string) *UpdateWhiteRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetRequestId(v string) *UpdateWhiteRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetSuccess(v bool) *UpdateWhiteRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) Validate() error {
	return dara.Validate(s)
}
