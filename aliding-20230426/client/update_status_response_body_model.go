// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateStatusResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateStatusResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateStatusResponseBody
	GetVendorType() *string
}

type UpdateStatusResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s UpdateStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStatusResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateStatusResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateStatusResponseBody) SetRequestId(v string) *UpdateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStatusResponseBody) SetVendorRequestId(v string) *UpdateStatusResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateStatusResponseBody) SetVendorType(v string) *UpdateStatusResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
