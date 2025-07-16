// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMultiDimTableRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v []*InsertMultiDimTableRecordResponseBodyValue) *InsertMultiDimTableRecordResponseBody
	GetValue() []*InsertMultiDimTableRecordResponseBodyValue
	SetRequestId(v string) *InsertMultiDimTableRecordResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *InsertMultiDimTableRecordResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *InsertMultiDimTableRecordResponseBody
	GetVendorType() *string
}

type InsertMultiDimTableRecordResponseBody struct {
	Value []*InsertMultiDimTableRecordResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
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

func (s InsertMultiDimTableRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordResponseBody) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordResponseBody) GetValue() []*InsertMultiDimTableRecordResponseBodyValue {
	return s.Value
}

func (s *InsertMultiDimTableRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertMultiDimTableRecordResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *InsertMultiDimTableRecordResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *InsertMultiDimTableRecordResponseBody) SetValue(v []*InsertMultiDimTableRecordResponseBodyValue) *InsertMultiDimTableRecordResponseBody {
	s.Value = v
	return s
}

func (s *InsertMultiDimTableRecordResponseBody) SetRequestId(v string) *InsertMultiDimTableRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertMultiDimTableRecordResponseBody) SetVendorRequestId(v string) *InsertMultiDimTableRecordResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *InsertMultiDimTableRecordResponseBody) SetVendorType(v string) *InsertMultiDimTableRecordResponseBody {
	s.VendorType = &v
	return s
}

func (s *InsertMultiDimTableRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type InsertMultiDimTableRecordResponseBodyValue struct {
	// example:
	//
	// hfauVBFJIo
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s InsertMultiDimTableRecordResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordResponseBodyValue) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordResponseBodyValue) GetId() *string {
	return s.Id
}

func (s *InsertMultiDimTableRecordResponseBodyValue) SetId(v string) *InsertMultiDimTableRecordResponseBodyValue {
	s.Id = &v
	return s
}

func (s *InsertMultiDimTableRecordResponseBodyValue) Validate() error {
	return dara.Validate(s)
}
