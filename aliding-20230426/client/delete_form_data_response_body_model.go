// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFormDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFormDataResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *DeleteFormDataResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DeleteFormDataResponseBody
	GetVendorType() *string
}

type DeleteFormDataResponseBody struct {
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

func (s DeleteFormDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFormDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFormDataResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DeleteFormDataResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DeleteFormDataResponseBody) SetRequestId(v string) *DeleteFormDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFormDataResponseBody) SetVendorRequestId(v string) *DeleteFormDataResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DeleteFormDataResponseBody) SetVendorType(v string) *DeleteFormDataResponseBody {
	s.VendorType = &v
	return s
}

func (s *DeleteFormDataResponseBody) Validate() error {
	return dara.Validate(s)
}
