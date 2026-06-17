// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DescribeModelServicesRequest
	GetGwClusterId() *string
	SetModelCategory(v string) *DescribeModelServicesRequest
	GetModelCategory() *string
	SetModelServiceIds(v string) *DescribeModelServicesRequest
	GetModelServiceIds() *string
	SetName(v string) *DescribeModelServicesRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeModelServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeModelServicesRequest
	GetPageSize() *int32
	SetProtocol(v string) *DescribeModelServicesRequest
	GetProtocol() *string
	SetRegionId(v string) *DescribeModelServicesRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeModelServicesRequest
	GetStatus() *string
}

type DescribeModelServicesRequest struct {
	// The ID of the gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The model category.
	//
	// example:
	//
	// text
	ModelCategory *string `json:"ModelCategory,omitempty" xml:"ModelCategory,omitempty"`
	// The IDs of the model services. Separate multiple IDs with a comma (,).
	//
	// example:
	//
	// ms-xxx,ms-xxxx
	ModelServiceIds *string `json:"ModelServiceIds,omitempty" xml:"ModelServiceIds,omitempty"`
	// The name of the model service.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol type. Valid values:
	//
	// - **openai**
	//
	// - **anthropic**
	//
	// - **bailian**: Models from Model Studio.
	//
	// - **vllm**
	//
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the model service.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeModelServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelServicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelServicesRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeModelServicesRequest) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *DescribeModelServicesRequest) GetModelServiceIds() *string {
	return s.ModelServiceIds
}

func (s *DescribeModelServicesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeModelServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeModelServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeModelServicesRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeModelServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeModelServicesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeModelServicesRequest) SetGwClusterId(v string) *DescribeModelServicesRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeModelServicesRequest) SetModelCategory(v string) *DescribeModelServicesRequest {
	s.ModelCategory = &v
	return s
}

func (s *DescribeModelServicesRequest) SetModelServiceIds(v string) *DescribeModelServicesRequest {
	s.ModelServiceIds = &v
	return s
}

func (s *DescribeModelServicesRequest) SetName(v string) *DescribeModelServicesRequest {
	s.Name = &v
	return s
}

func (s *DescribeModelServicesRequest) SetPageNumber(v int32) *DescribeModelServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelServicesRequest) SetPageSize(v int32) *DescribeModelServicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeModelServicesRequest) SetProtocol(v string) *DescribeModelServicesRequest {
	s.Protocol = &v
	return s
}

func (s *DescribeModelServicesRequest) SetRegionId(v string) *DescribeModelServicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeModelServicesRequest) SetStatus(v string) *DescribeModelServicesRequest {
	s.Status = &v
	return s
}

func (s *DescribeModelServicesRequest) Validate() error {
	return dara.Validate(s)
}
