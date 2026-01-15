// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *DescribeModelFeatureRequest
	GetCustomerModuleId() *int32
	SetFeatureTemplate(v string) *DescribeModelFeatureRequest
	GetFeatureTemplate() *string
}

type DescribeModelFeatureRequest struct {
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
	// example:
	//
	// FINANCE_51
	FeatureTemplate *string `json:"FeatureTemplate,omitempty" xml:"FeatureTemplate,omitempty"`
}

func (s DescribeModelFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelFeatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelFeatureRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DescribeModelFeatureRequest) GetFeatureTemplate() *string {
	return s.FeatureTemplate
}

func (s *DescribeModelFeatureRequest) SetCustomerModuleId(v int32) *DescribeModelFeatureRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DescribeModelFeatureRequest) SetFeatureTemplate(v string) *DescribeModelFeatureRequest {
	s.FeatureTemplate = &v
	return s
}

func (s *DescribeModelFeatureRequest) Validate() error {
	return dara.Validate(s)
}
