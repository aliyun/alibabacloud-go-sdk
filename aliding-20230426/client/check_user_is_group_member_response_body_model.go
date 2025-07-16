// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserIsGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckUserIsGroupMemberResponseBody
	GetRequestId() *string
	SetResult(v bool) *CheckUserIsGroupMemberResponseBody
	GetResult() *bool
	SetVendorRequestId(v string) *CheckUserIsGroupMemberResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CheckUserIsGroupMemberResponseBody
	GetVendorType() *string
}

type CheckUserIsGroupMemberResponseBody struct {
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

func (s CheckUserIsGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUserIsGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUserIsGroupMemberResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CheckUserIsGroupMemberResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CheckUserIsGroupMemberResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CheckUserIsGroupMemberResponseBody) SetRequestId(v string) *CheckUserIsGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserIsGroupMemberResponseBody) SetResult(v bool) *CheckUserIsGroupMemberResponseBody {
	s.Result = &v
	return s
}

func (s *CheckUserIsGroupMemberResponseBody) SetVendorRequestId(v string) *CheckUserIsGroupMemberResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CheckUserIsGroupMemberResponseBody) SetVendorType(v string) *CheckUserIsGroupMemberResponseBody {
	s.VendorType = &v
	return s
}

func (s *CheckUserIsGroupMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
