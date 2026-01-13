// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *DescribeCustomerModuleBasicInfoRequest
	GetAuthType() *string
	SetCustomerModuleId(v int32) *DescribeCustomerModuleBasicInfoRequest
	GetCustomerModuleId() *int32
}

type DescribeCustomerModuleBasicInfoRequest struct {
	// Authorization type.
	//
	// example:
	//
	// READ
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Customer model ID
	//
	// example:
	//
	// 1
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s DescribeCustomerModuleBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleBasicInfoRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeCustomerModuleBasicInfoRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DescribeCustomerModuleBasicInfoRequest) SetAuthType(v string) *DescribeCustomerModuleBasicInfoRequest {
	s.AuthType = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoRequest) SetCustomerModuleId(v int32) *DescribeCustomerModuleBasicInfoRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
