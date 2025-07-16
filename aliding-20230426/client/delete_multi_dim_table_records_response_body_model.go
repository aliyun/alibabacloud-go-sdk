// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *DeleteMultiDimTableRecordsResponseBody
	GetSuccess() *bool
	SetRequestId(v string) *DeleteMultiDimTableRecordsResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *DeleteMultiDimTableRecordsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DeleteMultiDimTableRecordsResponseBody
	GetVendorType() *string
}

type DeleteMultiDimTableRecordsResponseBody struct {
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

func (s DeleteMultiDimTableRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMultiDimTableRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultiDimTableRecordsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DeleteMultiDimTableRecordsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DeleteMultiDimTableRecordsResponseBody) SetSuccess(v bool) *DeleteMultiDimTableRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMultiDimTableRecordsResponseBody) SetRequestId(v string) *DeleteMultiDimTableRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultiDimTableRecordsResponseBody) SetVendorRequestId(v string) *DeleteMultiDimTableRecordsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DeleteMultiDimTableRecordsResponseBody) SetVendorType(v string) *DeleteMultiDimTableRecordsResponseBody {
	s.VendorType = &v
	return s
}

func (s *DeleteMultiDimTableRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
