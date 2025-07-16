// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateMultiDimTableResponseBody
	GetId() *string
	SetName(v string) *UpdateMultiDimTableResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateMultiDimTableResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateMultiDimTableResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateMultiDimTableResponseBody
	GetVendorType() *string
}

type UpdateMultiDimTableResponseBody struct {
	// example:
	//
	// []
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// []
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s UpdateMultiDimTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateMultiDimTableResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateMultiDimTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMultiDimTableResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateMultiDimTableResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateMultiDimTableResponseBody) SetId(v string) *UpdateMultiDimTableResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateMultiDimTableResponseBody) SetName(v string) *UpdateMultiDimTableResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateMultiDimTableResponseBody) SetRequestId(v string) *UpdateMultiDimTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMultiDimTableResponseBody) SetVendorRequestId(v string) *UpdateMultiDimTableResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateMultiDimTableResponseBody) SetVendorType(v string) *UpdateMultiDimTableResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateMultiDimTableResponseBody) Validate() error {
	return dara.Validate(s)
}
