// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllSheetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMultiDimTableAllSheetsResponseBody
	GetRequestId() *string
	SetValue(v []*GetMultiDimTableAllSheetsResponseBodyValue) *GetMultiDimTableAllSheetsResponseBody
	GetValue() []*GetMultiDimTableAllSheetsResponseBodyValue
	SetVendorRequestId(v string) *GetMultiDimTableAllSheetsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetMultiDimTableAllSheetsResponseBody
	GetVendorType() *string
}

type GetMultiDimTableAllSheetsResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// []
	Value []*GetMultiDimTableAllSheetsResponseBodyValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetMultiDimTableAllSheetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiDimTableAllSheetsResponseBody) GetValue() []*GetMultiDimTableAllSheetsResponseBodyValue {
	return s.Value
}

func (s *GetMultiDimTableAllSheetsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetMultiDimTableAllSheetsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetMultiDimTableAllSheetsResponseBody) SetRequestId(v string) *GetMultiDimTableAllSheetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiDimTableAllSheetsResponseBody) SetValue(v []*GetMultiDimTableAllSheetsResponseBodyValue) *GetMultiDimTableAllSheetsResponseBody {
	s.Value = v
	return s
}

func (s *GetMultiDimTableAllSheetsResponseBody) SetVendorRequestId(v string) *GetMultiDimTableAllSheetsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetMultiDimTableAllSheetsResponseBody) SetVendorType(v string) *GetMultiDimTableAllSheetsResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetMultiDimTableAllSheetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMultiDimTableAllSheetsResponseBodyValue struct {
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMultiDimTableAllSheetsResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsResponseBodyValue) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsResponseBodyValue) GetId() *string {
	return s.Id
}

func (s *GetMultiDimTableAllSheetsResponseBodyValue) GetName() *string {
	return s.Name
}

func (s *GetMultiDimTableAllSheetsResponseBodyValue) SetId(v string) *GetMultiDimTableAllSheetsResponseBodyValue {
	s.Id = &v
	return s
}

func (s *GetMultiDimTableAllSheetsResponseBodyValue) SetName(v string) *GetMultiDimTableAllSheetsResponseBodyValue {
	s.Name = &v
	return s
}

func (s *GetMultiDimTableAllSheetsResponseBodyValue) Validate() error {
	return dara.Validate(s)
}
