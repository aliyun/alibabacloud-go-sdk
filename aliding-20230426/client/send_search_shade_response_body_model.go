// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSearchShadeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v []interface{}) *SendSearchShadeResponseBody
	GetArguments() []interface{}
	SetRequestId(v string) *SendSearchShadeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendSearchShadeResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *SendSearchShadeResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SendSearchShadeResponseBody
	GetVendorType() *string
}

type SendSearchShadeResponseBody struct {
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

func (s SendSearchShadeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendSearchShadeResponseBody) GoString() string {
	return s.String()
}

func (s *SendSearchShadeResponseBody) GetArguments() []interface{} {
	return s.Arguments
}

func (s *SendSearchShadeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendSearchShadeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendSearchShadeResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SendSearchShadeResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SendSearchShadeResponseBody) SetArguments(v []interface{}) *SendSearchShadeResponseBody {
	s.Arguments = v
	return s
}

func (s *SendSearchShadeResponseBody) SetRequestId(v string) *SendSearchShadeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSearchShadeResponseBody) SetSuccess(v bool) *SendSearchShadeResponseBody {
	s.Success = &v
	return s
}

func (s *SendSearchShadeResponseBody) SetVendorRequestId(v string) *SendSearchShadeResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SendSearchShadeResponseBody) SetVendorType(v string) *SendSearchShadeResponseBody {
	s.VendorType = &v
	return s
}

func (s *SendSearchShadeResponseBody) Validate() error {
	return dara.Validate(s)
}
