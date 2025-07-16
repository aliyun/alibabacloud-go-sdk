// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddPermissionResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *AddPermissionResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *AddPermissionResponseBody
	GetVendorType() *string
}

type AddPermissionResponseBody struct {
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

func (s AddPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AddPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddPermissionResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *AddPermissionResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *AddPermissionResponseBody) SetRequestId(v string) *AddPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPermissionResponseBody) SetSuccess(v bool) *AddPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *AddPermissionResponseBody) SetVendorRequestId(v string) *AddPermissionResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *AddPermissionResponseBody) SetVendorType(v string) *AddPermissionResponseBody {
	s.VendorType = &v
	return s
}

func (s *AddPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
