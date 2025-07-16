// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubscribeEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubscribeEventResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *SubscribeEventResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SubscribeEventResponseBody
	GetVendorType() *string
}

type SubscribeEventResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SubscribeEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubscribeEventResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubscribeEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubscribeEventResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SubscribeEventResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SubscribeEventResponseBody) SetRequestId(v string) *SubscribeEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeEventResponseBody) SetSuccess(v bool) *SubscribeEventResponseBody {
	s.Success = &v
	return s
}

func (s *SubscribeEventResponseBody) SetVendorRequestId(v string) *SubscribeEventResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SubscribeEventResponseBody) SetVendorType(v string) *SubscribeEventResponseBody {
	s.VendorType = &v
	return s
}

func (s *SubscribeEventResponseBody) Validate() error {
	return dara.Validate(s)
}
