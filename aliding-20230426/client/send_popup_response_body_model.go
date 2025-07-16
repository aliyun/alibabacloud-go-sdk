// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPopupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v []interface{}) *SendPopupResponseBody
	GetArguments() []interface{}
	SetRequestId(v string) *SendPopupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendPopupResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *SendPopupResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SendPopupResponseBody
	GetVendorType() *string
}

type SendPopupResponseBody struct {
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

func (s SendPopupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendPopupResponseBody) GoString() string {
	return s.String()
}

func (s *SendPopupResponseBody) GetArguments() []interface{} {
	return s.Arguments
}

func (s *SendPopupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendPopupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendPopupResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SendPopupResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SendPopupResponseBody) SetArguments(v []interface{}) *SendPopupResponseBody {
	s.Arguments = v
	return s
}

func (s *SendPopupResponseBody) SetRequestId(v string) *SendPopupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendPopupResponseBody) SetSuccess(v bool) *SendPopupResponseBody {
	s.Success = &v
	return s
}

func (s *SendPopupResponseBody) SetVendorRequestId(v string) *SendPopupResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SendPopupResponseBody) SetVendorType(v string) *SendPopupResponseBody {
	s.VendorType = &v
	return s
}

func (s *SendPopupResponseBody) Validate() error {
	return dara.Validate(s)
}
