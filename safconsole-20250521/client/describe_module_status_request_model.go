// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *DescribeModuleStatusRequest
	GetCustomerModuleId() *int32
}

type DescribeModuleStatusRequest struct {
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s DescribeModuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeModuleStatusRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DescribeModuleStatusRequest) SetCustomerModuleId(v int32) *DescribeModuleStatusRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DescribeModuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
