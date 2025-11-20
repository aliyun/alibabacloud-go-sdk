// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *DocBlocksModifyResponseBody
	GetSuccess() *bool
	SetRequestId(v string) *DocBlocksModifyResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *DocBlocksModifyResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DocBlocksModifyResponseBody
	GetVendorType() *string
}

type DocBlocksModifyResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s DocBlocksModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksModifyResponseBody) GoString() string {
	return s.String()
}

func (s *DocBlocksModifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DocBlocksModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocBlocksModifyResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DocBlocksModifyResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DocBlocksModifyResponseBody) SetSuccess(v bool) *DocBlocksModifyResponseBody {
	s.Success = &v
	return s
}

func (s *DocBlocksModifyResponseBody) SetRequestId(v string) *DocBlocksModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocBlocksModifyResponseBody) SetVendorRequestId(v string) *DocBlocksModifyResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DocBlocksModifyResponseBody) SetVendorType(v string) *DocBlocksModifyResponseBody {
	s.VendorType = &v
	return s
}

func (s *DocBlocksModifyResponseBody) Validate() error {
	return dara.Validate(s)
}
