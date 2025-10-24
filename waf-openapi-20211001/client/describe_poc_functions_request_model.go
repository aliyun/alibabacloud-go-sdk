// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocFunctionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePocFunctionsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribePocFunctionsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribePocFunctionsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribePocFunctionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribePocFunctionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePocFunctionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePocFunctionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePocFunctionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePocFunctionsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribePocFunctionsRequest) SetInstanceId(v string) *DescribePocFunctionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePocFunctionsRequest) SetRegionId(v string) *DescribePocFunctionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePocFunctionsRequest) SetResourceManagerResourceGroupId(v string) *DescribePocFunctionsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribePocFunctionsRequest) Validate() error {
	return dara.Validate(s)
}
