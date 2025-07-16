// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TerminateInstanceResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *TerminateInstanceResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *TerminateInstanceResponseBody
	GetVendorType() *string
}

type TerminateInstanceResponseBody struct {
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

func (s TerminateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateInstanceResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *TerminateInstanceResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *TerminateInstanceResponseBody) SetRequestId(v string) *TerminateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateInstanceResponseBody) SetVendorRequestId(v string) *TerminateInstanceResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *TerminateInstanceResponseBody) SetVendorType(v string) *TerminateInstanceResponseBody {
	s.VendorType = &v
	return s
}

func (s *TerminateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
