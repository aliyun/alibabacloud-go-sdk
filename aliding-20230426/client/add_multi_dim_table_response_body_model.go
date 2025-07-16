// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMultiDimTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *AddMultiDimTableResponseBody
	GetId() *string
	SetName(v string) *AddMultiDimTableResponseBody
	GetName() *string
	SetRequestId(v string) *AddMultiDimTableResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *AddMultiDimTableResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *AddMultiDimTableResponseBody
	GetVendorType() *string
}

type AddMultiDimTableResponseBody struct {
	// example:
	//
	// r1R7q3QmWew5lo02fxB7xxxxxxx
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s AddMultiDimTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableResponseBody) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableResponseBody) GetId() *string {
	return s.Id
}

func (s *AddMultiDimTableResponseBody) GetName() *string {
	return s.Name
}

func (s *AddMultiDimTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMultiDimTableResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *AddMultiDimTableResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *AddMultiDimTableResponseBody) SetId(v string) *AddMultiDimTableResponseBody {
	s.Id = &v
	return s
}

func (s *AddMultiDimTableResponseBody) SetName(v string) *AddMultiDimTableResponseBody {
	s.Name = &v
	return s
}

func (s *AddMultiDimTableResponseBody) SetRequestId(v string) *AddMultiDimTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMultiDimTableResponseBody) SetVendorRequestId(v string) *AddMultiDimTableResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *AddMultiDimTableResponseBody) SetVendorType(v string) *AddMultiDimTableResponseBody {
	s.VendorType = &v
	return s
}

func (s *AddMultiDimTableResponseBody) Validate() error {
	return dara.Validate(s)
}
