// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleMetaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *DescribeCustomerModuleMetaInfoRequest
	GetCustomerModuleId() *int32
}

type DescribeCustomerModuleMetaInfoRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s DescribeCustomerModuleMetaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleMetaInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleMetaInfoRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DescribeCustomerModuleMetaInfoRequest) SetCustomerModuleId(v int32) *DescribeCustomerModuleMetaInfoRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoRequest) Validate() error {
	return dara.Validate(s)
}
