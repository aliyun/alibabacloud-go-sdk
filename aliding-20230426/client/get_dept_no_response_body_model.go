// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeptNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeptNo(v string) *GetDeptNoResponseBody
	GetDeptNo() *string
	SetRequestId(v string) *GetDeptNoResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetDeptNoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDeptNoResponseBody
	GetVendorType() *string
}

type GetDeptNoResponseBody struct {
	// example:
	//
	// 12345
	DeptNo *string `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
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

func (s GetDeptNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeptNoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeptNoResponseBody) GetDeptNo() *string {
	return s.DeptNo
}

func (s *GetDeptNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeptNoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDeptNoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDeptNoResponseBody) SetDeptNo(v string) *GetDeptNoResponseBody {
	s.DeptNo = &v
	return s
}

func (s *GetDeptNoResponseBody) SetRequestId(v string) *GetDeptNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeptNoResponseBody) SetVendorRequestId(v string) *GetDeptNoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDeptNoResponseBody) SetVendorType(v string) *GetDeptNoResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDeptNoResponseBody) Validate() error {
	return dara.Validate(s)
}
