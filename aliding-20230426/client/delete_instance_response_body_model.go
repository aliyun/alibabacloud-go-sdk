// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInstanceResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *DeleteInstanceResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DeleteInstanceResponseBody
	GetVendorType() *string
}

type DeleteInstanceResponseBody struct {
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

func (s DeleteInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DeleteInstanceResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetVendorRequestId(v string) *DeleteInstanceResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetVendorType(v string) *DeleteInstanceResponseBody {
	s.VendorType = &v
	return s
}

func (s *DeleteInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
