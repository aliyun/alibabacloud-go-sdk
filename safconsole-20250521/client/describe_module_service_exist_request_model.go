// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModuleServiceExistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *DescribeModuleServiceExistRequest
	GetCustomerModuleId() *int32
}

type DescribeModuleServiceExistRequest struct {
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s DescribeModuleServiceExistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleServiceExistRequest) GoString() string {
	return s.String()
}

func (s *DescribeModuleServiceExistRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DescribeModuleServiceExistRequest) SetCustomerModuleId(v int32) *DescribeModuleServiceExistRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DescribeModuleServiceExistRequest) Validate() error {
	return dara.Validate(s)
}
