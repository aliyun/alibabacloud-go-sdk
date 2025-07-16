// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMinutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopMinutesResponseBody
	GetCode() *string
	SetRequestId(v string) *StopMinutesResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *StopMinutesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *StopMinutesResponseBody
	GetVendorType() *string
}

type StopMinutesResponseBody struct {
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

func (s StopMinutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *StopMinutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopMinutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopMinutesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *StopMinutesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *StopMinutesResponseBody) SetCode(v string) *StopMinutesResponseBody {
	s.Code = &v
	return s
}

func (s *StopMinutesResponseBody) SetRequestId(v string) *StopMinutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMinutesResponseBody) SetVendorRequestId(v string) *StopMinutesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *StopMinutesResponseBody) SetVendorType(v string) *StopMinutesResponseBody {
	s.VendorType = &v
	return s
}

func (s *StopMinutesResponseBody) Validate() error {
	return dara.Validate(s)
}
