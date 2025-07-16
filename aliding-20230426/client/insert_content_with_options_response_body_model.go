// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertContentWithOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InsertContentWithOptionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsertContentWithOptionsResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *InsertContentWithOptionsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *InsertContentWithOptionsResponseBody
	GetVendorType() *string
}

type InsertContentWithOptionsResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s InsertContentWithOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertContentWithOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *InsertContentWithOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertContentWithOptionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsertContentWithOptionsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *InsertContentWithOptionsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *InsertContentWithOptionsResponseBody) SetRequestId(v string) *InsertContentWithOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertContentWithOptionsResponseBody) SetSuccess(v bool) *InsertContentWithOptionsResponseBody {
	s.Success = &v
	return s
}

func (s *InsertContentWithOptionsResponseBody) SetVendorRequestId(v string) *InsertContentWithOptionsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *InsertContentWithOptionsResponseBody) SetVendorType(v string) *InsertContentWithOptionsResponseBody {
	s.VendorType = &v
	return s
}

func (s *InsertContentWithOptionsResponseBody) Validate() error {
	return dara.Validate(s)
}
