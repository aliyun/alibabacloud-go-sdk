// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostRuleStatusChangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PostRuleStatusChangeResponseBody
	GetCode() *int32
	SetData(v interface{}) *PostRuleStatusChangeResponseBody
	GetData() interface{}
	SetMessage(v string) *PostRuleStatusChangeResponseBody
	GetMessage() *string
	SetRequestId(v string) *PostRuleStatusChangeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostRuleStatusChangeResponseBody
	GetSuccess() *bool
}

type PostRuleStatusChangeResponseBody struct {
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

func (s PostRuleStatusChangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostRuleStatusChangeResponseBody) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PostRuleStatusChangeResponseBody) GetData() interface{} {
	return s.Data
}

func (s *PostRuleStatusChangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PostRuleStatusChangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostRuleStatusChangeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostRuleStatusChangeResponseBody) SetCode(v int32) *PostRuleStatusChangeResponseBody {
	s.Code = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetData(v interface{}) *PostRuleStatusChangeResponseBody {
	s.Data = v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetMessage(v string) *PostRuleStatusChangeResponseBody {
	s.Message = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetRequestId(v string) *PostRuleStatusChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetSuccess(v bool) *PostRuleStatusChangeResponseBody {
	s.Success = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) Validate() error {
	return dara.Validate(s)
}
