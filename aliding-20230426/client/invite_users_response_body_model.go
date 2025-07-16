// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InviteUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InviteUsersResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *InviteUsersResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *InviteUsersResponseBody
	GetVendorType() *string
}

type InviteUsersResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	VendorType      *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s InviteUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *InviteUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InviteUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InviteUsersResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *InviteUsersResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *InviteUsersResponseBody) SetRequestId(v string) *InviteUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *InviteUsersResponseBody) SetSuccess(v bool) *InviteUsersResponseBody {
	s.Success = &v
	return s
}

func (s *InviteUsersResponseBody) SetVendorRequestId(v string) *InviteUsersResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *InviteUsersResponseBody) SetVendorType(v string) *InviteUsersResponseBody {
	s.VendorType = &v
	return s
}

func (s *InviteUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
