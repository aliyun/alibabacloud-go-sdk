// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhiteRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteWhiteRuleListResponseBody
	GetCode() *int32
	SetData(v interface{}) *DeleteWhiteRuleListResponseBody
	GetData() interface{}
	SetMessage(v string) *DeleteWhiteRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWhiteRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWhiteRuleListResponseBody
	GetSuccess() *bool
}

type DeleteWhiteRuleListResponseBody struct {
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

func (s DeleteWhiteRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhiteRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteWhiteRuleListResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteWhiteRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWhiteRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWhiteRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWhiteRuleListResponseBody) SetCode(v int32) *DeleteWhiteRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetData(v interface{}) *DeleteWhiteRuleListResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetMessage(v string) *DeleteWhiteRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetRequestId(v string) *DeleteWhiteRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetSuccess(v bool) *DeleteWhiteRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) Validate() error {
	return dara.Validate(s)
}
