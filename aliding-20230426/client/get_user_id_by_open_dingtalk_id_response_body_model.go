// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOpenDingtalkIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserIdByOpenDingtalkIdResponseBody
	GetRequestId() *string
	SetUserId(v string) *GetUserIdByOpenDingtalkIdResponseBody
	GetUserId() *string
	SetVendorRequestId(v string) *GetUserIdByOpenDingtalkIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetUserIdByOpenDingtalkIdResponseBody
	GetVendorType() *string
}

type GetUserIdByOpenDingtalkIdResponseBody struct {
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

func (s GetUserIdByOpenDingtalkIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOpenDingtalkIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) SetRequestId(v string) *GetUserIdByOpenDingtalkIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) SetUserId(v string) *GetUserIdByOpenDingtalkIdResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) SetVendorRequestId(v string) *GetUserIdByOpenDingtalkIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) SetVendorType(v string) *GetUserIdByOpenDingtalkIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdResponseBody) Validate() error {
	return dara.Validate(s)
}
