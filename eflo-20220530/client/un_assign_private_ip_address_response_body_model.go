// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssignPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnAssignPrivateIpAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UnAssignPrivateIpAddressResponseBody
	GetCode() *int32
	SetContent(v *UnAssignPrivateIpAddressResponseBodyContent) *UnAssignPrivateIpAddressResponseBody
	GetContent() *UnAssignPrivateIpAddressResponseBodyContent
	SetMessage(v string) *UnAssignPrivateIpAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnAssignPrivateIpAddressResponseBody
	GetRequestId() *string
}

type UnAssignPrivateIpAddressResponseBody struct {
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
	// The response parameters.
	Content *UnAssignPrivateIpAddressResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission to do this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnAssignPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnAssignPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnAssignPrivateIpAddressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnAssignPrivateIpAddressResponseBody) GetContent() *UnAssignPrivateIpAddressResponseBodyContent {
	return s.Content
}

func (s *UnAssignPrivateIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnAssignPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnAssignPrivateIpAddressResponseBody) SetAccessDeniedDetail(v string) *UnAssignPrivateIpAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBody) SetCode(v int32) *UnAssignPrivateIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBody) SetContent(v *UnAssignPrivateIpAddressResponseBodyContent) *UnAssignPrivateIpAddressResponseBody {
	s.Content = v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBody) SetMessage(v string) *UnAssignPrivateIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBody) SetRequestId(v string) *UnAssignPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnAssignPrivateIpAddressResponseBodyContent struct {
	// IP unique identifier
	//
	// example:
	//
	// sip-xxxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-bp164jwjpdq4lnsy83s5
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s UnAssignPrivateIpAddressResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UnAssignPrivateIpAddressResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressResponseBodyContent) GetIpName() *string {
	return s.IpName
}

func (s *UnAssignPrivateIpAddressResponseBodyContent) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *UnAssignPrivateIpAddressResponseBodyContent) SetIpName(v string) *UnAssignPrivateIpAddressResponseBodyContent {
	s.IpName = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBodyContent) SetNetworkInterfaceId(v string) *UnAssignPrivateIpAddressResponseBodyContent {
	s.NetworkInterfaceId = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
