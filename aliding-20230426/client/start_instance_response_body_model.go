// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartInstanceResponseBody
	GetRequestId() *string
	SetResult(v string) *StartInstanceResponseBody
	GetResult() *string
	SetVendorRequestId(v string) *StartInstanceResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *StartInstanceResponseBody
	GetVendorType() *string
}

type StartInstanceResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// f30233fb-72e1-xxx-xxx
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

func (s StartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstanceResponseBody) GetResult() *string {
	return s.Result
}

func (s *StartInstanceResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *StartInstanceResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) SetResult(v string) *StartInstanceResponseBody {
	s.Result = &v
	return s
}

func (s *StartInstanceResponseBody) SetVendorRequestId(v string) *StartInstanceResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *StartInstanceResponseBody) SetVendorType(v string) *StartInstanceResponseBody {
	s.VendorType = &v
	return s
}

func (s *StartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
