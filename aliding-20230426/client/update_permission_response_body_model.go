// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePermissionResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *UpdatePermissionResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdatePermissionResponseBody
	GetVendorType() *string
}

type UpdatePermissionResponseBody struct {
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

func (s UpdatePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePermissionResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdatePermissionResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdatePermissionResponseBody) SetRequestId(v string) *UpdatePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePermissionResponseBody) SetSuccess(v bool) *UpdatePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePermissionResponseBody) SetVendorRequestId(v string) *UpdatePermissionResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdatePermissionResponseBody) SetVendorType(v string) *UpdatePermissionResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdatePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
