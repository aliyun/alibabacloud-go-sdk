// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostFinishCustomizeRuleTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PostFinishCustomizeRuleTestResponseBody
	GetCode() *int32
	SetData(v interface{}) *PostFinishCustomizeRuleTestResponseBody
	GetData() interface{}
	SetMessage(v string) *PostFinishCustomizeRuleTestResponseBody
	GetMessage() *string
	SetRequestId(v string) *PostFinishCustomizeRuleTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostFinishCustomizeRuleTestResponseBody
	GetSuccess() *bool
}

type PostFinishCustomizeRuleTestResponseBody struct {
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

func (s PostFinishCustomizeRuleTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostFinishCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *PostFinishCustomizeRuleTestResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PostFinishCustomizeRuleTestResponseBody) GetData() interface{} {
	return s.Data
}

func (s *PostFinishCustomizeRuleTestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PostFinishCustomizeRuleTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostFinishCustomizeRuleTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetCode(v int32) *PostFinishCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetData(v interface{}) *PostFinishCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetMessage(v string) *PostFinishCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetRequestId(v string) *PostFinishCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetSuccess(v bool) *PostFinishCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) Validate() error {
	return dara.Validate(s)
}
