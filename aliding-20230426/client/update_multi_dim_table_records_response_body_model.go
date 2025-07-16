// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v []*UpdateMultiDimTableRecordsResponseBodyValue) *UpdateMultiDimTableRecordsResponseBody
	GetValue() []*UpdateMultiDimTableRecordsResponseBodyValue
	SetRequestId(v string) *UpdateMultiDimTableRecordsResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateMultiDimTableRecordsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateMultiDimTableRecordsResponseBody
	GetVendorType() *string
}

type UpdateMultiDimTableRecordsResponseBody struct {
	Value []*UpdateMultiDimTableRecordsResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
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

func (s UpdateMultiDimTableRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsResponseBody) GetValue() []*UpdateMultiDimTableRecordsResponseBodyValue {
	return s.Value
}

func (s *UpdateMultiDimTableRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMultiDimTableRecordsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateMultiDimTableRecordsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateMultiDimTableRecordsResponseBody) SetValue(v []*UpdateMultiDimTableRecordsResponseBodyValue) *UpdateMultiDimTableRecordsResponseBody {
	s.Value = v
	return s
}

func (s *UpdateMultiDimTableRecordsResponseBody) SetRequestId(v string) *UpdateMultiDimTableRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMultiDimTableRecordsResponseBody) SetVendorRequestId(v string) *UpdateMultiDimTableRecordsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateMultiDimTableRecordsResponseBody) SetVendorType(v string) *UpdateMultiDimTableRecordsResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateMultiDimTableRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateMultiDimTableRecordsResponseBodyValue struct {
	// example:
	//
	// []
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateMultiDimTableRecordsResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsResponseBodyValue) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsResponseBodyValue) GetId() *string {
	return s.Id
}

func (s *UpdateMultiDimTableRecordsResponseBodyValue) SetId(v string) *UpdateMultiDimTableRecordsResponseBodyValue {
	s.Id = &v
	return s
}

func (s *UpdateMultiDimTableRecordsResponseBodyValue) Validate() error {
	return dara.Validate(s)
}
