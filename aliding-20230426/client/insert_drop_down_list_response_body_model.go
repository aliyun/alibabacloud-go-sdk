// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDropDownListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetA1Notation(v string) *InsertDropDownListResponseBody
	GetA1Notation() *string
	SetRequestId(v string) *InsertDropDownListResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *InsertDropDownListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *InsertDropDownListResponseBody
	GetVendorType() *string
}

type InsertDropDownListResponseBody struct {
	// example:
	//
	// A3:C3
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
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

func (s InsertDropDownListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListResponseBody) GoString() string {
	return s.String()
}

func (s *InsertDropDownListResponseBody) GetA1Notation() *string {
	return s.A1Notation
}

func (s *InsertDropDownListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertDropDownListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *InsertDropDownListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *InsertDropDownListResponseBody) SetA1Notation(v string) *InsertDropDownListResponseBody {
	s.A1Notation = &v
	return s
}

func (s *InsertDropDownListResponseBody) SetRequestId(v string) *InsertDropDownListResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertDropDownListResponseBody) SetVendorRequestId(v string) *InsertDropDownListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *InsertDropDownListResponseBody) SetVendorType(v string) *InsertDropDownListResponseBody {
	s.VendorType = &v
	return s
}

func (s *InsertDropDownListResponseBody) Validate() error {
	return dara.Validate(s)
}
