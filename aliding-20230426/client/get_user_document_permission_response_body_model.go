// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDocumentPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthLevel(v int64) *GetUserDocumentPermissionResponseBody
	GetAuthLevel() *int64
	SetHasPermission(v bool) *GetUserDocumentPermissionResponseBody
	GetHasPermission() *bool
	SetRequestId(v string) *GetUserDocumentPermissionResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetUserDocumentPermissionResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetUserDocumentPermissionResponseBody
	GetVendorType() *string
}

type GetUserDocumentPermissionResponseBody struct {
	// example:
	//
	// -3
	AuthLevel *int64 `json:"authLevel,omitempty" xml:"authLevel,omitempty"`
	// example:
	//
	// true
	HasPermission *bool `json:"hasPermission,omitempty" xml:"hasPermission,omitempty"`
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

func (s GetUserDocumentPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserDocumentPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDocumentPermissionResponseBody) GetAuthLevel() *int64 {
	return s.AuthLevel
}

func (s *GetUserDocumentPermissionResponseBody) GetHasPermission() *bool {
	return s.HasPermission
}

func (s *GetUserDocumentPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserDocumentPermissionResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetUserDocumentPermissionResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetUserDocumentPermissionResponseBody) SetAuthLevel(v int64) *GetUserDocumentPermissionResponseBody {
	s.AuthLevel = &v
	return s
}

func (s *GetUserDocumentPermissionResponseBody) SetHasPermission(v bool) *GetUserDocumentPermissionResponseBody {
	s.HasPermission = &v
	return s
}

func (s *GetUserDocumentPermissionResponseBody) SetRequestId(v string) *GetUserDocumentPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserDocumentPermissionResponseBody) SetVendorRequestId(v string) *GetUserDocumentPermissionResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetUserDocumentPermissionResponseBody) SetVendorType(v string) *GetUserDocumentPermissionResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetUserDocumentPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
