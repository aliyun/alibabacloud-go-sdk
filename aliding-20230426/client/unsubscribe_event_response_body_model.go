// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnsubscribeEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnsubscribeEventResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *UnsubscribeEventResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UnsubscribeEventResponseBody
	GetVendorType() *string
}

type UnsubscribeEventResponseBody struct {
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

func (s UnsubscribeEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeEventResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnsubscribeEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnsubscribeEventResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UnsubscribeEventResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UnsubscribeEventResponseBody) SetRequestId(v string) *UnsubscribeEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeEventResponseBody) SetSuccess(v bool) *UnsubscribeEventResponseBody {
	s.Success = &v
	return s
}

func (s *UnsubscribeEventResponseBody) SetVendorRequestId(v string) *UnsubscribeEventResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UnsubscribeEventResponseBody) SetVendorType(v string) *UnsubscribeEventResponseBody {
	s.VendorType = &v
	return s
}

func (s *UnsubscribeEventResponseBody) Validate() error {
	return dara.Validate(s)
}
