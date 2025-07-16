// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveFormRemarkResponseBody
	GetRequestId() *string
	SetResult(v int64) *SaveFormRemarkResponseBody
	GetResult() *int64
	SetVendorRequestId(v string) *SaveFormRemarkResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SaveFormRemarkResponseBody
	GetVendorType() *string
}

type SaveFormRemarkResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *int64  `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SaveFormRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveFormRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveFormRemarkResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *SaveFormRemarkResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SaveFormRemarkResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SaveFormRemarkResponseBody) SetRequestId(v string) *SaveFormRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveFormRemarkResponseBody) SetResult(v int64) *SaveFormRemarkResponseBody {
	s.Result = &v
	return s
}

func (s *SaveFormRemarkResponseBody) SetVendorRequestId(v string) *SaveFormRemarkResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SaveFormRemarkResponseBody) SetVendorType(v string) *SaveFormRemarkResponseBody {
	s.VendorType = &v
	return s
}

func (s *SaveFormRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
