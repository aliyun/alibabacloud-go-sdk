// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEventWhiteruleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PostEventWhiteruleListResponseBody
	GetCode() *int32
	SetData(v string) *PostEventWhiteruleListResponseBody
	GetData() *string
	SetMessage(v string) *PostEventWhiteruleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *PostEventWhiteruleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostEventWhiteruleListResponseBody
	GetSuccess() *bool
}

type PostEventWhiteruleListResponseBody struct {
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
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s PostEventWhiteruleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostEventWhiteruleListResponseBody) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PostEventWhiteruleListResponseBody) GetData() *string {
	return s.Data
}

func (s *PostEventWhiteruleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PostEventWhiteruleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostEventWhiteruleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostEventWhiteruleListResponseBody) SetCode(v int32) *PostEventWhiteruleListResponseBody {
	s.Code = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetData(v string) *PostEventWhiteruleListResponseBody {
	s.Data = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetMessage(v string) *PostEventWhiteruleListResponseBody {
	s.Message = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetRequestId(v string) *PostEventWhiteruleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetSuccess(v bool) *PostEventWhiteruleListResponseBody {
	s.Success = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) Validate() error {
	return dara.Validate(s)
}
