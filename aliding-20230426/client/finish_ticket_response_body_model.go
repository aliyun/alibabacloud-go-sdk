// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FinishTicketResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *FinishTicketResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *FinishTicketResponseBody
	GetVendorType() *string
}

type FinishTicketResponseBody struct {
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

func (s FinishTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketResponseBody) GoString() string {
	return s.String()
}

func (s *FinishTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FinishTicketResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *FinishTicketResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *FinishTicketResponseBody) SetRequestId(v string) *FinishTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishTicketResponseBody) SetVendorRequestId(v string) *FinishTicketResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *FinishTicketResponseBody) SetVendorType(v string) *FinishTicketResponseBody {
	s.VendorType = &v
	return s
}

func (s *FinishTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
