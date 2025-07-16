// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePermissionResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *DeletePermissionResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DeletePermissionResponseBody
	GetVendorType() *string
}

type DeletePermissionResponseBody struct {
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

func (s DeletePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePermissionResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DeletePermissionResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DeletePermissionResponseBody) SetRequestId(v string) *DeletePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePermissionResponseBody) SetSuccess(v bool) *DeletePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePermissionResponseBody) SetVendorRequestId(v string) *DeletePermissionResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DeletePermissionResponseBody) SetVendorType(v string) *DeletePermissionResponseBody {
	s.VendorType = &v
	return s
}

func (s *DeletePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
