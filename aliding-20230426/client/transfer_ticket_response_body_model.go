// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferTicketResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *TransferTicketResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *TransferTicketResponseBody
	GetVendorType() *string
}

type TransferTicketResponseBody struct {
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

func (s TransferTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketResponseBody) GoString() string {
	return s.String()
}

func (s *TransferTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferTicketResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *TransferTicketResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *TransferTicketResponseBody) SetRequestId(v string) *TransferTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferTicketResponseBody) SetVendorRequestId(v string) *TransferTicketResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *TransferTicketResponseBody) SetVendorType(v string) *TransferTicketResponseBody {
	s.VendorType = &v
	return s
}

func (s *TransferTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
