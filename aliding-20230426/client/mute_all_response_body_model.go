// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MuteAllResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MuteAllResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *MuteAllResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *MuteAllResponseBody
	GetVendorType() *string
}

type MuteAllResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s MuteAllResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MuteAllResponseBody) GoString() string {
	return s.String()
}

func (s *MuteAllResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MuteAllResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MuteAllResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *MuteAllResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *MuteAllResponseBody) SetRequestId(v string) *MuteAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteAllResponseBody) SetSuccess(v bool) *MuteAllResponseBody {
	s.Success = &v
	return s
}

func (s *MuteAllResponseBody) SetVendorRequestId(v string) *MuteAllResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *MuteAllResponseBody) SetVendorType(v string) *MuteAllResponseBody {
	s.VendorType = &v
	return s
}

func (s *MuteAllResponseBody) Validate() error {
	return dara.Validate(s)
}
