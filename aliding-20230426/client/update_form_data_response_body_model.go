// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFormDataResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateFormDataResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateFormDataResponseBody
	GetVendorType() *string
}

type UpdateFormDataResponseBody struct {
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

func (s UpdateFormDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFormDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFormDataResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateFormDataResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateFormDataResponseBody) SetRequestId(v string) *UpdateFormDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFormDataResponseBody) SetVendorRequestId(v string) *UpdateFormDataResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateFormDataResponseBody) SetVendorType(v string) *UpdateFormDataResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateFormDataResponseBody) Validate() error {
	return dara.Validate(s)
}
