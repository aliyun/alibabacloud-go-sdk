// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMinutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartMinutesResponseBody
	GetCode() *string
	SetRequestId(v string) *StartMinutesResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *StartMinutesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *StartMinutesResponseBody
	GetVendorType() *string
}

type StartMinutesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
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

func (s StartMinutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *StartMinutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartMinutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartMinutesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *StartMinutesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *StartMinutesResponseBody) SetCode(v string) *StartMinutesResponseBody {
	s.Code = &v
	return s
}

func (s *StartMinutesResponseBody) SetRequestId(v string) *StartMinutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMinutesResponseBody) SetVendorRequestId(v string) *StartMinutesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *StartMinutesResponseBody) SetVendorType(v string) *StartMinutesResponseBody {
	s.VendorType = &v
	return s
}

func (s *StartMinutesResponseBody) Validate() error {
	return dara.Validate(s)
}
