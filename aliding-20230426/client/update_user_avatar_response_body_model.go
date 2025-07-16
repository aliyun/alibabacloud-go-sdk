// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAvatarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrcode(v string) *UpdateUserAvatarResponseBody
	GetErrcode() *string
	SetErrmsg(v string) *UpdateUserAvatarResponseBody
	GetErrmsg() *string
	SetRequestId(v string) *UpdateUserAvatarResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateUserAvatarResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateUserAvatarResponseBody
	GetVendorType() *string
}

type UpdateUserAvatarResponseBody struct {
	// example:
	//
	// 0
	Errcode *string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	Errmsg  *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
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

func (s UpdateUserAvatarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserAvatarResponseBody) GetErrcode() *string {
	return s.Errcode
}

func (s *UpdateUserAvatarResponseBody) GetErrmsg() *string {
	return s.Errmsg
}

func (s *UpdateUserAvatarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserAvatarResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateUserAvatarResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateUserAvatarResponseBody) SetErrcode(v string) *UpdateUserAvatarResponseBody {
	s.Errcode = &v
	return s
}

func (s *UpdateUserAvatarResponseBody) SetErrmsg(v string) *UpdateUserAvatarResponseBody {
	s.Errmsg = &v
	return s
}

func (s *UpdateUserAvatarResponseBody) SetRequestId(v string) *UpdateUserAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserAvatarResponseBody) SetVendorRequestId(v string) *UpdateUserAvatarResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateUserAvatarResponseBody) SetVendorType(v string) *UpdateUserAvatarResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateUserAvatarResponseBody) Validate() error {
	return dara.Validate(s)
}
