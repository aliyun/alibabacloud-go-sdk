// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MuteMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MuteMembersResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *MuteMembersResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *MuteMembersResponseBody
	GetVendorType() *string
}

type MuteMembersResponseBody struct {
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

func (s MuteMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MuteMembersResponseBody) GoString() string {
	return s.String()
}

func (s *MuteMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MuteMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MuteMembersResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *MuteMembersResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *MuteMembersResponseBody) SetRequestId(v string) *MuteMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteMembersResponseBody) SetSuccess(v bool) *MuteMembersResponseBody {
	s.Success = &v
	return s
}

func (s *MuteMembersResponseBody) SetVendorRequestId(v string) *MuteMembersResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *MuteMembersResponseBody) SetVendorType(v string) *MuteMembersResponseBody {
	s.VendorType = &v
	return s
}

func (s *MuteMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
