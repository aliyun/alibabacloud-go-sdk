// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketMemoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTicketMemoResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *AddTicketMemoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *AddTicketMemoResponseBody
	GetVendorType() *string
}

type AddTicketMemoResponseBody struct {
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

func (s AddTicketMemoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoResponseBody) GoString() string {
	return s.String()
}

func (s *AddTicketMemoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTicketMemoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *AddTicketMemoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *AddTicketMemoResponseBody) SetRequestId(v string) *AddTicketMemoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTicketMemoResponseBody) SetVendorRequestId(v string) *AddTicketMemoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *AddTicketMemoResponseBody) SetVendorType(v string) *AddTicketMemoResponseBody {
	s.VendorType = &v
	return s
}

func (s *AddTicketMemoResponseBody) Validate() error {
	return dara.Validate(s)
}
