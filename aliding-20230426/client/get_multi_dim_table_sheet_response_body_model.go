// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableSheetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetMultiDimTableSheetResponseBody
	GetId() *string
	SetName(v string) *GetMultiDimTableSheetResponseBody
	GetName() *string
	SetRequestId(v string) *GetMultiDimTableSheetResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetMultiDimTableSheetResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetMultiDimTableSheetResponseBody
	GetVendorType() *string
}

type GetMultiDimTableSheetResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// Sheet1
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

func (s GetMultiDimTableSheetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableSheetResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableSheetResponseBody) GetId() *string {
	return s.Id
}

func (s *GetMultiDimTableSheetResponseBody) GetName() *string {
	return s.Name
}

func (s *GetMultiDimTableSheetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiDimTableSheetResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetMultiDimTableSheetResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetMultiDimTableSheetResponseBody) SetId(v string) *GetMultiDimTableSheetResponseBody {
	s.Id = &v
	return s
}

func (s *GetMultiDimTableSheetResponseBody) SetName(v string) *GetMultiDimTableSheetResponseBody {
	s.Name = &v
	return s
}

func (s *GetMultiDimTableSheetResponseBody) SetRequestId(v string) *GetMultiDimTableSheetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiDimTableSheetResponseBody) SetVendorRequestId(v string) *GetMultiDimTableSheetResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetMultiDimTableSheetResponseBody) SetVendorType(v string) *GetMultiDimTableSheetResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetMultiDimTableSheetResponseBody) Validate() error {
	return dara.Validate(s)
}
