// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedirectTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RedirectTaskResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *RedirectTaskResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *RedirectTaskResponseBody
	GetVendorType() *string
}

type RedirectTaskResponseBody struct {
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

func (s RedirectTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RedirectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RedirectTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RedirectTaskResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *RedirectTaskResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *RedirectTaskResponseBody) SetRequestId(v string) *RedirectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RedirectTaskResponseBody) SetVendorRequestId(v string) *RedirectTaskResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *RedirectTaskResponseBody) SetVendorType(v string) *RedirectTaskResponseBody {
	s.VendorType = &v
	return s
}

func (s *RedirectTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
