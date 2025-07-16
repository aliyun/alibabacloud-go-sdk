// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOrgIdAndStaffIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserIdByOrgIdAndStaffIdResponseBody
	GetRequestId() *string
	SetUserId(v string) *GetUserIdByOrgIdAndStaffIdResponseBody
	GetUserId() *string
	SetVendorRequestId(v string) *GetUserIdByOrgIdAndStaffIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetUserIdByOrgIdAndStaffIdResponseBody
	GetVendorType() *string
}

type GetUserIdByOrgIdAndStaffIdResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 01223245436
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetUserIdByOrgIdAndStaffIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOrgIdAndStaffIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) SetRequestId(v string) *GetUserIdByOrgIdAndStaffIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) SetUserId(v string) *GetUserIdByOrgIdAndStaffIdResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) SetVendorRequestId(v string) *GetUserIdByOrgIdAndStaffIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) SetVendorType(v string) *GetUserIdByOrgIdAndStaffIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdResponseBody) Validate() error {
	return dara.Validate(s)
}
