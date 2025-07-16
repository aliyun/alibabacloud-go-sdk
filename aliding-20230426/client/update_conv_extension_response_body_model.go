// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConvExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *UpdateConvExtensionResponseBody
	GetSuccess() *bool
	SetRequestId(v string) *UpdateConvExtensionResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateConvExtensionResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateConvExtensionResponseBody
	GetVendorType() *string
}

type UpdateConvExtensionResponseBody struct {
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s UpdateConvExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConvExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConvExtensionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConvExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConvExtensionResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateConvExtensionResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateConvExtensionResponseBody) SetSuccess(v bool) *UpdateConvExtensionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConvExtensionResponseBody) SetRequestId(v string) *UpdateConvExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConvExtensionResponseBody) SetVendorRequestId(v string) *UpdateConvExtensionResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateConvExtensionResponseBody) SetVendorType(v string) *UpdateConvExtensionResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateConvExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}
