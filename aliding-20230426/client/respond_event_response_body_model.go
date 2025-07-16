// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RespondEventResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *RespondEventResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *RespondEventResponseBody
	GetVendorType() *string
}

type RespondEventResponseBody struct {
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

func (s RespondEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RespondEventResponseBody) GoString() string {
	return s.String()
}

func (s *RespondEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RespondEventResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *RespondEventResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *RespondEventResponseBody) SetRequestId(v string) *RespondEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *RespondEventResponseBody) SetVendorRequestId(v string) *RespondEventResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *RespondEventResponseBody) SetVendorType(v string) *RespondEventResponseBody {
	s.VendorType = &v
	return s
}

func (s *RespondEventResponseBody) Validate() error {
	return dara.Validate(s)
}
