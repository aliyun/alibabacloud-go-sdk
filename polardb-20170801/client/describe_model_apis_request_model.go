// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DescribeModelApisRequest
	GetGwClusterId() *string
	SetModelApiIds(v string) *DescribeModelApisRequest
	GetModelApiIds() *string
	SetModelCategory(v string) *DescribeModelApisRequest
	GetModelCategory() *string
	SetName(v string) *DescribeModelApisRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeModelApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeModelApisRequest
	GetPageSize() *int32
	SetPathPrefix(v string) *DescribeModelApisRequest
	GetPathPrefix() *string
	SetProtocol(v string) *DescribeModelApisRequest
	GetProtocol() *string
	SetRegionId(v string) *DescribeModelApisRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeModelApisRequest
	GetStatus() *string
}

type DescribeModelApisRequest struct {
	// The ID of the gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The IDs of the model APIs. Separate multiple IDs with a comma.
	//
	// example:
	//
	// mi-xxx,mi-xxxx
	ModelApiIds *string `json:"ModelApiIds,omitempty" xml:"ModelApiIds,omitempty"`
	// The model category. Valid values:
	//
	// - **text**
	//
	// - **embedding**
	//
	// - **rerank**
	//
	// example:
	//
	// text
	ModelCategory *string `json:"ModelCategory,omitempty" xml:"ModelCategory,omitempty"`
	// The name of the model API.
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
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	//   The default value is **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The API path prefix.
	//
	// example:
	//
	// /test
	PathPrefix *string `json:"PathPrefix,omitempty" xml:"PathPrefix,omitempty"`
	// The protocol. Valid values:
	//
	// - **openai**
	//
	// - **anthropic**
	//
	// - **bailian**
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
	// The model API status.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeModelApisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelApisRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeModelApisRequest) GetModelApiIds() *string {
	return s.ModelApiIds
}

func (s *DescribeModelApisRequest) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *DescribeModelApisRequest) GetName() *string {
	return s.Name
}

func (s *DescribeModelApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeModelApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeModelApisRequest) GetPathPrefix() *string {
	return s.PathPrefix
}

func (s *DescribeModelApisRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeModelApisRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeModelApisRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeModelApisRequest) SetGwClusterId(v string) *DescribeModelApisRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeModelApisRequest) SetModelApiIds(v string) *DescribeModelApisRequest {
	s.ModelApiIds = &v
	return s
}

func (s *DescribeModelApisRequest) SetModelCategory(v string) *DescribeModelApisRequest {
	s.ModelCategory = &v
	return s
}

func (s *DescribeModelApisRequest) SetName(v string) *DescribeModelApisRequest {
	s.Name = &v
	return s
}

func (s *DescribeModelApisRequest) SetPageNumber(v int32) *DescribeModelApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelApisRequest) SetPageSize(v int32) *DescribeModelApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeModelApisRequest) SetPathPrefix(v string) *DescribeModelApisRequest {
	s.PathPrefix = &v
	return s
}

func (s *DescribeModelApisRequest) SetProtocol(v string) *DescribeModelApisRequest {
	s.Protocol = &v
	return s
}

func (s *DescribeModelApisRequest) SetRegionId(v string) *DescribeModelApisRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeModelApisRequest) SetStatus(v string) *DescribeModelApisRequest {
	s.Status = &v
	return s
}

func (s *DescribeModelApisRequest) Validate() error {
	return dara.Validate(s)
}
