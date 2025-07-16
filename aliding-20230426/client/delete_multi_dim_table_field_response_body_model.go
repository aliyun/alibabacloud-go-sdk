// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *DeleteMultiDimTableFieldResponseBody
	GetSuccess() *bool
	SetRequestId(v string) *DeleteMultiDimTableFieldResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *DeleteMultiDimTableFieldResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DeleteMultiDimTableFieldResponseBody
	GetVendorType() *string
}

type DeleteMultiDimTableFieldResponseBody struct {
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

func (s DeleteMultiDimTableFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableFieldResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMultiDimTableFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultiDimTableFieldResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DeleteMultiDimTableFieldResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DeleteMultiDimTableFieldResponseBody) SetSuccess(v bool) *DeleteMultiDimTableFieldResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMultiDimTableFieldResponseBody) SetRequestId(v string) *DeleteMultiDimTableFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultiDimTableFieldResponseBody) SetVendorRequestId(v string) *DeleteMultiDimTableFieldResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DeleteMultiDimTableFieldResponseBody) SetVendorType(v string) *DeleteMultiDimTableFieldResponseBody {
	s.VendorType = &v
	return s
}

func (s *DeleteMultiDimTableFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
