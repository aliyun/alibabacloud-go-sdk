// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutOrgAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SignOutOrgAccountResponseBody
	GetRequestId() *string
	SetResult(v bool) *SignOutOrgAccountResponseBody
	GetResult() *bool
	SetVendorRequestId(v string) *SignOutOrgAccountResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SignOutOrgAccountResponseBody
	GetVendorType() *string
}

type SignOutOrgAccountResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SignOutOrgAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SignOutOrgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *SignOutOrgAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SignOutOrgAccountResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SignOutOrgAccountResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SignOutOrgAccountResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SignOutOrgAccountResponseBody) SetRequestId(v string) *SignOutOrgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignOutOrgAccountResponseBody) SetResult(v bool) *SignOutOrgAccountResponseBody {
	s.Result = &v
	return s
}

func (s *SignOutOrgAccountResponseBody) SetVendorRequestId(v string) *SignOutOrgAccountResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SignOutOrgAccountResponseBody) SetVendorType(v string) *SignOutOrgAccountResponseBody {
	s.VendorType = &v
	return s
}

func (s *SignOutOrgAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
