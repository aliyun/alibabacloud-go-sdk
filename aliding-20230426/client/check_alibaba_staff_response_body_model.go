// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAlibabaStaffResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsAlibabaStaff(v bool) *CheckAlibabaStaffResponseBody
	GetIsAlibabaStaff() *bool
	SetRequestId(v string) *CheckAlibabaStaffResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *CheckAlibabaStaffResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CheckAlibabaStaffResponseBody
	GetVendorType() *string
}

type CheckAlibabaStaffResponseBody struct {
	// example:
	//
	// true
	IsAlibabaStaff *bool `json:"isAlibabaStaff,omitempty" xml:"isAlibabaStaff,omitempty"`
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

func (s CheckAlibabaStaffResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAlibabaStaffResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAlibabaStaffResponseBody) GetIsAlibabaStaff() *bool {
	return s.IsAlibabaStaff
}

func (s *CheckAlibabaStaffResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAlibabaStaffResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CheckAlibabaStaffResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CheckAlibabaStaffResponseBody) SetIsAlibabaStaff(v bool) *CheckAlibabaStaffResponseBody {
	s.IsAlibabaStaff = &v
	return s
}

func (s *CheckAlibabaStaffResponseBody) SetRequestId(v string) *CheckAlibabaStaffResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAlibabaStaffResponseBody) SetVendorRequestId(v string) *CheckAlibabaStaffResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CheckAlibabaStaffResponseBody) SetVendorType(v string) *CheckAlibabaStaffResponseBody {
	s.VendorType = &v
	return s
}

func (s *CheckAlibabaStaffResponseBody) Validate() error {
	return dara.Validate(s)
}
