// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *GetDingtalkProjectionInfoResponseBody
	GetData() interface{}
	SetRequestId(v string) *GetDingtalkProjectionInfoResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetDingtalkProjectionInfoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDingtalkProjectionInfoResponseBody
	GetVendorType() *string
}

type GetDingtalkProjectionInfoResponseBody struct {
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
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

func (s GetDingtalkProjectionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionInfoResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetDingtalkProjectionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDingtalkProjectionInfoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDingtalkProjectionInfoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDingtalkProjectionInfoResponseBody) SetData(v interface{}) *GetDingtalkProjectionInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDingtalkProjectionInfoResponseBody) SetRequestId(v string) *GetDingtalkProjectionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDingtalkProjectionInfoResponseBody) SetVendorRequestId(v string) *GetDingtalkProjectionInfoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDingtalkProjectionInfoResponseBody) SetVendorType(v string) *GetDingtalkProjectionInfoResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDingtalkProjectionInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
