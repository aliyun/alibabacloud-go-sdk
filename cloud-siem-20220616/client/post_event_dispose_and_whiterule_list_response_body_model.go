// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEventDisposeAndWhiteruleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PostEventDisposeAndWhiteruleListResponseBody
	GetCode() *int32
	SetData(v string) *PostEventDisposeAndWhiteruleListResponseBody
	GetData() *string
	SetMessage(v string) *PostEventDisposeAndWhiteruleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *PostEventDisposeAndWhiteruleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostEventDisposeAndWhiteruleListResponseBody
	GetSuccess() *bool
}

type PostEventDisposeAndWhiteruleListResponseBody struct {
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

func (s PostEventDisposeAndWhiteruleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListResponseBody) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) GetData() *string {
	return s.Data
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetCode(v int32) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Code = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetData(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Data = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetMessage(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Message = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetRequestId(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetSuccess(v bool) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Success = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) Validate() error {
	return dara.Validate(s)
}
