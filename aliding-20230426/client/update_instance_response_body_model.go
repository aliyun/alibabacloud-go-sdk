// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateInstanceResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateInstanceResponseBody
	GetVendorType() *string
}

type UpdateInstanceResponseBody struct {
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

func (s UpdateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateInstanceResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetVendorRequestId(v string) *UpdateInstanceResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetVendorType(v string) *UpdateInstanceResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
