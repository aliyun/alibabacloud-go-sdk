// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateMultiDimTableFieldResponseBody
	GetId() *string
	SetRequestId(v string) *UpdateMultiDimTableFieldResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateMultiDimTableFieldResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateMultiDimTableFieldResponseBody
	GetVendorType() *string
}

type UpdateMultiDimTableFieldResponseBody struct {
	// example:
	//
	// r1R7q3QmWew5lo02fxB7xxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s UpdateMultiDimTableFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableFieldResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableFieldResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateMultiDimTableFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMultiDimTableFieldResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateMultiDimTableFieldResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateMultiDimTableFieldResponseBody) SetId(v string) *UpdateMultiDimTableFieldResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateMultiDimTableFieldResponseBody) SetRequestId(v string) *UpdateMultiDimTableFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMultiDimTableFieldResponseBody) SetVendorRequestId(v string) *UpdateMultiDimTableFieldResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateMultiDimTableFieldResponseBody) SetVendorType(v string) *UpdateMultiDimTableFieldResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateMultiDimTableFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
