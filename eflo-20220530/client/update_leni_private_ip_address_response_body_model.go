// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLeniPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateLeniPrivateIpAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UpdateLeniPrivateIpAddressResponseBody
	GetCode() *int32
	SetContent(v *UpdateLeniPrivateIpAddressResponseBodyContent) *UpdateLeniPrivateIpAddressResponseBody
	GetContent() *UpdateLeniPrivateIpAddressResponseBodyContent
	SetMessage(v string) *UpdateLeniPrivateIpAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLeniPrivateIpAddressResponseBody
	GetRequestId() *string
}

type UpdateLeniPrivateIpAddressResponseBody struct {
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
	Content *UpdateLeniPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLeniPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLeniPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLeniPrivateIpAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateLeniPrivateIpAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateLeniPrivateIpAddressResponseBody) GetContent() *UpdateLeniPrivateIpAddressResponseBodyContent {
	return s.Content
}

func (s *UpdateLeniPrivateIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLeniPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *UpdateLeniPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetCode(v int32) *UpdateLeniPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetContent(v *UpdateLeniPrivateIpAddressResponseBodyContent) *UpdateLeniPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetMessage(v string) *UpdateLeniPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) SetRequestId(v string) *UpdateLeniPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLeniPrivateIpAddressResponseBodyContent struct {
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
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
}

func (s UpdateLeniPrivateIpAddressResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateLeniPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateLeniPrivateIpAddressResponseBodyContent) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *UpdateLeniPrivateIpAddressResponseBodyContent) GetIpName() *string {
	return s.IpName
}

func (s *UpdateLeniPrivateIpAddressResponseBodyContent) SetElasticNetworkInterfaceId(v string) *UpdateLeniPrivateIpAddressResponseBodyContent {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBodyContent) SetIpName(v string) *UpdateLeniPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
