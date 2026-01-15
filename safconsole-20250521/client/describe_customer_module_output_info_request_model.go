// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleOutputInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *DescribeCustomerModuleOutputInfoRequest
	GetAuthType() *string
	SetCustomerModuleId(v int32) *DescribeCustomerModuleOutputInfoRequest
	GetCustomerModuleId() *int32
}

type DescribeCustomerModuleOutputInfoRequest struct {
	// example:
	//
	// READ
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s DescribeCustomerModuleOutputInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleOutputInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleOutputInfoRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeCustomerModuleOutputInfoRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DescribeCustomerModuleOutputInfoRequest) SetAuthType(v string) *DescribeCustomerModuleOutputInfoRequest {
	s.AuthType = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoRequest) SetCustomerModuleId(v int32) *DescribeCustomerModuleOutputInfoRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoRequest) Validate() error {
	return dara.Validate(s)
}
