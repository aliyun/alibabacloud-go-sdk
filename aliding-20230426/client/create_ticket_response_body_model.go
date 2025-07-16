// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpenTicketId(v string) *CreateTicketResponseBody
	GetOpenTicketId() *string
	SetRequestId(v string) *CreateTicketResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *CreateTicketResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateTicketResponseBody
	GetVendorType() *string
}

type CreateTicketResponseBody struct {
	// example:
	//
	// a8iSxxxxtgiE
	OpenTicketId *string `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
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

func (s CreateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *CreateTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTicketResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateTicketResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateTicketResponseBody) SetOpenTicketId(v string) *CreateTicketResponseBody {
	s.OpenTicketId = &v
	return s
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetVendorRequestId(v string) *CreateTicketResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetVendorType(v string) *CreateTicketResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
