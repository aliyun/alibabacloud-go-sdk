// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignLeniPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnassignLeniPrivateIpAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UnassignLeniPrivateIpAddressResponseBody
	GetCode() *int32
	SetContent(v *UnassignLeniPrivateIpAddressResponseBodyContent) *UnassignLeniPrivateIpAddressResponseBody
	GetContent() *UnassignLeniPrivateIpAddressResponseBodyContent
	SetMessage(v string) *UnassignLeniPrivateIpAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnassignLeniPrivateIpAddressResponseBody
	GetRequestId() *string
}

type UnassignLeniPrivateIpAddressResponseBody struct {
	// The details about the access denial.
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
	//
	// example:
	//
	// {}
	Content *UnassignLeniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassignLeniPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassignLeniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnassignLeniPrivateIpAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnassignLeniPrivateIpAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnassignLeniPrivateIpAddressResponseBody) GetContent() *UnassignLeniPrivateIpAddressResponseBodyContent {
	return s.Content
}

func (s *UnassignLeniPrivateIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnassignLeniPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *UnassignLeniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetCode(v int32) *UnassignLeniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetContent(v *UnassignLeniPrivateIpAddressResponseBodyContent) *UnassignLeniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetMessage(v string) *UnassignLeniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) SetRequestId(v string) *UnassignLeniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnassignLeniPrivateIpAddressResponseBodyContent struct {
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
	// sip-dqvs****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
}

func (s UnassignLeniPrivateIpAddressResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UnassignLeniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UnassignLeniPrivateIpAddressResponseBodyContent) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *UnassignLeniPrivateIpAddressResponseBodyContent) GetIpName() *string {
	return s.IpName
}

func (s *UnassignLeniPrivateIpAddressResponseBodyContent) SetElasticNetworkInterfaceId(v string) *UnassignLeniPrivateIpAddressResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBodyContent) SetIpName(v string) *UnassignLeniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
