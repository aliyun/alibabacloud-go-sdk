// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachElasticNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DetachElasticNetworkInterfaceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DetachElasticNetworkInterfaceResponseBody
	GetCode() *int32
	SetMessage(v string) *DetachElasticNetworkInterfaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachElasticNetworkInterfaceResponseBody
	GetRequestId() *string
}

type DetachElasticNetworkInterfaceResponseBody struct {
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response
	//
	// example:
	//
	// You don\\"t have the permission to do this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachElasticNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachElasticNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachElasticNetworkInterfaceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DetachElasticNetworkInterfaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DetachElasticNetworkInterfaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachElasticNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachElasticNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *DetachElasticNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponseBody) SetCode(v int32) *DetachElasticNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponseBody) SetMessage(v string) *DetachElasticNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponseBody) SetRequestId(v string) *DetachElasticNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
