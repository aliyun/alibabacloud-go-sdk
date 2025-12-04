// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubnetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteSubnetResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteSubnetResponseBody
	GetCode() *int32
	SetContent(v interface{}) *DeleteSubnetResponseBody
	GetContent() interface{}
	SetMessage(v string) *DeleteSubnetResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSubnetResponseBody
	GetRequestId() *string
}

type DeleteSubnetResponseBody struct {
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response content (if the resource has dependent resources, the existing dependent resources will be returned)
	//
	// example:
	//
	// {
	//
	//       "nc": [
	//
	//             {}
	//
	//       ]
	//
	// }
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A56F7D3C-8850-5AF4-A342-2D71C9A9D1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubnetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubnetResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteSubnetResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteSubnetResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *DeleteSubnetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSubnetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubnetResponseBody) SetAccessDeniedDetail(v string) *DeleteSubnetResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteSubnetResponseBody) SetCode(v int32) *DeleteSubnetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSubnetResponseBody) SetContent(v interface{}) *DeleteSubnetResponseBody {
	s.Content = v
	return s
}

func (s *DeleteSubnetResponseBody) SetMessage(v string) *DeleteSubnetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSubnetResponseBody) SetRequestId(v string) *DeleteSubnetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubnetResponseBody) Validate() error {
	return dara.Validate(s)
}
