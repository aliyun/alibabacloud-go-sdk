// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddRecordPermissionResponseBody
	GetCode() *string
	SetRequestId(v string) *AddRecordPermissionResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *AddRecordPermissionResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *AddRecordPermissionResponseBody
	GetVendorType() *string
}

type AddRecordPermissionResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
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

func (s AddRecordPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddRecordPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRecordPermissionResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *AddRecordPermissionResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *AddRecordPermissionResponseBody) SetCode(v string) *AddRecordPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *AddRecordPermissionResponseBody) SetRequestId(v string) *AddRecordPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecordPermissionResponseBody) SetVendorRequestId(v string) *AddRecordPermissionResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *AddRecordPermissionResponseBody) SetVendorType(v string) *AddRecordPermissionResponseBody {
	s.VendorType = &v
	return s
}

func (s *AddRecordPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
