// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachElasticNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AttachElasticNetworkInterfaceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *AttachElasticNetworkInterfaceResponseBody
	GetCode() *int32
	SetContent(v interface{}) *AttachElasticNetworkInterfaceResponseBody
	GetContent() interface{}
	SetMessage(v string) *AttachElasticNetworkInterfaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AttachElasticNetworkInterfaceResponseBody
	GetRequestId() *string
}

type AttachElasticNetworkInterfaceResponseBody struct {
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
	// Response body
	//
	// example:
	//
	// []
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachElasticNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachElasticNetworkInterfaceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AttachElasticNetworkInterfaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AttachElasticNetworkInterfaceResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *AttachElasticNetworkInterfaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AttachElasticNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *AttachElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetCode(v int32) *AttachElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetContent(v interface{}) *AttachElasticNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetMessage(v string) *AttachElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) SetRequestId(v string) *AttachElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
