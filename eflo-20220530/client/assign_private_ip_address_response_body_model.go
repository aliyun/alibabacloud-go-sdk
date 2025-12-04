// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AssignPrivateIpAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *AssignPrivateIpAddressResponseBody
	GetCode() *int32
	SetContent(v *AssignPrivateIpAddressResponseBodyContent) *AssignPrivateIpAddressResponseBody
	GetContent() *AssignPrivateIpAddressResponseBodyContent
	SetMessage(v string) *AssignPrivateIpAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssignPrivateIpAddressResponseBody
	GetRequestId() *string
}

type AssignPrivateIpAddressResponseBody struct {
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
	// The response data.
	Content *AssignPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
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

func (s AssignPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AssignPrivateIpAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AssignPrivateIpAddressResponseBody) GetContent() *AssignPrivateIpAddressResponseBodyContent {
	return s.Content
}

func (s *AssignPrivateIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssignPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *AssignPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBody) SetCode(v int32) *AssignPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBody) SetContent(v *AssignPrivateIpAddressResponseBodyContent) *AssignPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *AssignPrivateIpAddressResponseBody) SetMessage(v string) *AssignPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBody) SetRequestId(v string) *AssignPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssignPrivateIpAddressResponseBodyContent struct {
	// The unique IP identifier.
	//
	// example:
	//
	// sip-8exxqa2r
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s AssignPrivateIpAddressResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressResponseBodyContent) GetIpName() *string {
	return s.IpName
}

func (s *AssignPrivateIpAddressResponseBodyContent) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssignPrivateIpAddressResponseBodyContent) SetIpName(v string) *AssignPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBodyContent) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressResponseBodyContent {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignPrivateIpAddressResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
