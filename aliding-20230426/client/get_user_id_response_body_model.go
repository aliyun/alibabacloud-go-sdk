// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserIdResponseBody
	GetRequestId() *string
	SetUserId(v string) *GetUserIdResponseBody
	GetUserId() *string
	SetVendorRequestId(v string) *GetUserIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetUserIdResponseBody
	GetVendorType() *string
}

type GetUserIdResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 012345
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

func (s GetUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserIdResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetUserIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetUserIdResponseBody) SetRequestId(v string) *GetUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserIdResponseBody) SetUserId(v string) *GetUserIdResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserIdResponseBody) SetVendorRequestId(v string) *GetUserIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetUserIdResponseBody) SetVendorType(v string) *GetUserIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}
