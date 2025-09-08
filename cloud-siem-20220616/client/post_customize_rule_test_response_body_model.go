// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostCustomizeRuleTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PostCustomizeRuleTestResponseBody
	GetCode() *int32
	SetData(v interface{}) *PostCustomizeRuleTestResponseBody
	GetData() interface{}
	SetMessage(v string) *PostCustomizeRuleTestResponseBody
	GetMessage() *string
	SetRequestId(v string) *PostCustomizeRuleTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostCustomizeRuleTestResponseBody
	GetSuccess() *bool
}

type PostCustomizeRuleTestResponseBody struct {
	// The HTTP status code that is returned.
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

func (s PostCustomizeRuleTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PostCustomizeRuleTestResponseBody) GetData() interface{} {
	return s.Data
}

func (s *PostCustomizeRuleTestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PostCustomizeRuleTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostCustomizeRuleTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostCustomizeRuleTestResponseBody) SetCode(v int32) *PostCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetData(v interface{}) *PostCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetMessage(v string) *PostCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetRequestId(v string) *PostCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetSuccess(v bool) *PostCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) Validate() error {
	return dara.Validate(s)
}
