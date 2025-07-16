// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDriveSpaceResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *DeleteDriveSpaceResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DeleteDriveSpaceResponseBody
	GetVendorType() *string
}

type DeleteDriveSpaceResponseBody struct {
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

func (s DeleteDriveSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDriveSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDriveSpaceResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DeleteDriveSpaceResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DeleteDriveSpaceResponseBody) SetRequestId(v string) *DeleteDriveSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDriveSpaceResponseBody) SetVendorRequestId(v string) *DeleteDriveSpaceResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DeleteDriveSpaceResponseBody) SetVendorType(v string) *DeleteDriveSpaceResponseBody {
	s.VendorType = &v
	return s
}

func (s *DeleteDriveSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
