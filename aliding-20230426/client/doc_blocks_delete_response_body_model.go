// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *DocBlocksDeleteResponseBody
	GetSuccess() *bool
	SetRequestId(v string) *DocBlocksDeleteResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *DocBlocksDeleteResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DocBlocksDeleteResponseBody
	GetVendorType() *string
}

type DocBlocksDeleteResponseBody struct {
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

func (s DocBlocksDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *DocBlocksDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DocBlocksDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocBlocksDeleteResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DocBlocksDeleteResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DocBlocksDeleteResponseBody) SetSuccess(v bool) *DocBlocksDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *DocBlocksDeleteResponseBody) SetRequestId(v string) *DocBlocksDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocBlocksDeleteResponseBody) SetVendorRequestId(v string) *DocBlocksDeleteResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DocBlocksDeleteResponseBody) SetVendorType(v string) *DocBlocksDeleteResponseBody {
	s.VendorType = &v
	return s
}

func (s *DocBlocksDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
