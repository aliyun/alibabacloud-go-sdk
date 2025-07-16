// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEmployeeFieldValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchEmployeeFieldValuesResponseBody
	GetRequestId() *string
	SetResult(v string) *SearchEmployeeFieldValuesResponseBody
	GetResult() *string
	SetVendorRequestId(v string) *SearchEmployeeFieldValuesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SearchEmployeeFieldValuesResponseBody
	GetVendorType() *string
}

type SearchEmployeeFieldValuesResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// manager123
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SearchEmployeeFieldValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchEmployeeFieldValuesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchEmployeeFieldValuesResponseBody) GetResult() *string {
	return s.Result
}

func (s *SearchEmployeeFieldValuesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SearchEmployeeFieldValuesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SearchEmployeeFieldValuesResponseBody) SetRequestId(v string) *SearchEmployeeFieldValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEmployeeFieldValuesResponseBody) SetResult(v string) *SearchEmployeeFieldValuesResponseBody {
	s.Result = &v
	return s
}

func (s *SearchEmployeeFieldValuesResponseBody) SetVendorRequestId(v string) *SearchEmployeeFieldValuesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SearchEmployeeFieldValuesResponseBody) SetVendorType(v string) *SearchEmployeeFieldValuesResponseBody {
	s.VendorType = &v
	return s
}

func (s *SearchEmployeeFieldValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
