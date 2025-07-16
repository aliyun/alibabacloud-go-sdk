// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBannerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v []interface{}) *SendBannerResponseBody
	GetArguments() []interface{}
	SetRequestId(v string) *SendBannerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendBannerResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *SendBannerResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SendBannerResponseBody
	GetVendorType() *string
}

type SendBannerResponseBody struct {
	Arguments []interface{} `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SendBannerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendBannerResponseBody) GoString() string {
	return s.String()
}

func (s *SendBannerResponseBody) GetArguments() []interface{} {
	return s.Arguments
}

func (s *SendBannerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendBannerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendBannerResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SendBannerResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SendBannerResponseBody) SetArguments(v []interface{}) *SendBannerResponseBody {
	s.Arguments = v
	return s
}

func (s *SendBannerResponseBody) SetRequestId(v string) *SendBannerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendBannerResponseBody) SetSuccess(v bool) *SendBannerResponseBody {
	s.Success = &v
	return s
}

func (s *SendBannerResponseBody) SetVendorRequestId(v string) *SendBannerResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SendBannerResponseBody) SetVendorType(v string) *SendBannerResponseBody {
	s.VendorType = &v
	return s
}

func (s *SendBannerResponseBody) Validate() error {
	return dara.Validate(s)
}
