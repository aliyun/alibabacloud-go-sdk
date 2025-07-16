// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveFormDataResponseBody
	GetRequestId() *string
	SetResult(v string) *SaveFormDataResponseBody
	GetResult() *string
	SetVendorRequestId(v string) *SaveFormDataResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SaveFormDataResponseBody
	GetVendorType() *string
}

type SaveFormDataResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// FORM-EF6xxx
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SaveFormDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveFormDataResponseBody) GetResult() *string {
	return s.Result
}

func (s *SaveFormDataResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SaveFormDataResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SaveFormDataResponseBody) SetRequestId(v string) *SaveFormDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveFormDataResponseBody) SetResult(v string) *SaveFormDataResponseBody {
	s.Result = &v
	return s
}

func (s *SaveFormDataResponseBody) SetVendorRequestId(v string) *SaveFormDataResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SaveFormDataResponseBody) SetVendorType(v string) *SaveFormDataResponseBody {
	s.VendorType = &v
	return s
}

func (s *SaveFormDataResponseBody) Validate() error {
	return dara.Validate(s)
}
