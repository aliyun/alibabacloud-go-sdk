// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocUpdateContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DocUpdateContentResponseBody
	GetRequestId() *string
	SetValue(v bool) *DocUpdateContentResponseBody
	GetValue() *bool
	SetVendorRequestId(v string) *DocUpdateContentResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DocUpdateContentResponseBody
	GetVendorType() *string
}

type DocUpdateContentResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Value *bool `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s DocUpdateContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocUpdateContentResponseBody) GoString() string {
	return s.String()
}

func (s *DocUpdateContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocUpdateContentResponseBody) GetValue() *bool {
	return s.Value
}

func (s *DocUpdateContentResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DocUpdateContentResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DocUpdateContentResponseBody) SetRequestId(v string) *DocUpdateContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocUpdateContentResponseBody) SetValue(v bool) *DocUpdateContentResponseBody {
	s.Value = &v
	return s
}

func (s *DocUpdateContentResponseBody) SetVendorRequestId(v string) *DocUpdateContentResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DocUpdateContentResponseBody) SetVendorType(v string) *DocUpdateContentResponseBody {
	s.VendorType = &v
	return s
}

func (s *DocUpdateContentResponseBody) Validate() error {
	return dara.Validate(s)
}
