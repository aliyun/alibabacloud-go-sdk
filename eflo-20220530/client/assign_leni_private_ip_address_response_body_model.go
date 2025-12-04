// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignLeniPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AssignLeniPrivateIpAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *AssignLeniPrivateIpAddressResponseBody
	GetCode() *int32
	SetContent(v *AssignLeniPrivateIpAddressResponseBodyContent) *AssignLeniPrivateIpAddressResponseBody
	GetContent() *AssignLeniPrivateIpAddressResponseBodyContent
	SetMessage(v string) *AssignLeniPrivateIpAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssignLeniPrivateIpAddressResponseBody
	GetRequestId() *string
}

type AssignLeniPrivateIpAddressResponseBody struct {
	// The details about the access denial.
	//
	// >  This parameter is returned only if Resource Access Management (RAM) permission verification failed.
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
	// The response data.
	Content *AssignLeniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignLeniPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignLeniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssignLeniPrivateIpAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AssignLeniPrivateIpAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AssignLeniPrivateIpAddressResponseBody) GetContent() *AssignLeniPrivateIpAddressResponseBodyContent {
	return s.Content
}

func (s *AssignLeniPrivateIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssignLeniPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *AssignLeniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetCode(v int32) *AssignLeniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetContent(v *AssignLeniPrivateIpAddressResponseBodyContent) *AssignLeniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetMessage(v string) *AssignLeniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) SetRequestId(v string) *AssignLeniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssignLeniPrivateIpAddressResponseBodyContent struct {
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-lzwx****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
}

func (s AssignLeniPrivateIpAddressResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s AssignLeniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AssignLeniPrivateIpAddressResponseBodyContent) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *AssignLeniPrivateIpAddressResponseBodyContent) GetIpName() *string {
	return s.IpName
}

func (s *AssignLeniPrivateIpAddressResponseBodyContent) SetElasticNetworkInterfaceId(v string) *AssignLeniPrivateIpAddressResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBodyContent) SetIpName(v string) *AssignLeniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
