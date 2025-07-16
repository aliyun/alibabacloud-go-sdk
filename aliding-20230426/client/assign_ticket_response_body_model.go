// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssignTicketResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *AssignTicketResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *AssignTicketResponseBody
	GetVendorType() *string
}

type AssignTicketResponseBody struct {
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

func (s AssignTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketResponseBody) GoString() string {
	return s.String()
}

func (s *AssignTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignTicketResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *AssignTicketResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *AssignTicketResponseBody) SetRequestId(v string) *AssignTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignTicketResponseBody) SetVendorRequestId(v string) *AssignTicketResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *AssignTicketResponseBody) SetVendorType(v string) *AssignTicketResponseBody {
	s.VendorType = &v
	return s
}

func (s *AssignTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
