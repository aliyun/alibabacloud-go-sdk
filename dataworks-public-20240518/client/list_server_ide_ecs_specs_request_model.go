// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeEcsSpecsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *ListServerIdeEcsSpecsRequest
	GetAcceleratorType() *string
	SetMaxResults(v int32) *ListServerIdeEcsSpecsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerIdeEcsSpecsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListServerIdeEcsSpecsRequest
	GetResourceGroupId() *string
}

type ListServerIdeEcsSpecsRequest struct {
	// The accelerator type. Valid values:
	//
	// - CPU: uses only CPU.
	//
	// - GPU: uses GPU acceleration.
	//
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The maximum number of records to return in a single request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token used to retrieve the next page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// CAESG****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The DataWorks resource group identifier. You can specify the numeric ID of the resource group or the full identifier in the Serverless_res_group_{tenantId}_{resgId} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_123456789012345_9876543210****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListServerIdeEcsSpecsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeEcsSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListServerIdeEcsSpecsRequest) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListServerIdeEcsSpecsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerIdeEcsSpecsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerIdeEcsSpecsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServerIdeEcsSpecsRequest) SetAcceleratorType(v string) *ListServerIdeEcsSpecsRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListServerIdeEcsSpecsRequest) SetMaxResults(v int32) *ListServerIdeEcsSpecsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerIdeEcsSpecsRequest) SetNextToken(v string) *ListServerIdeEcsSpecsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerIdeEcsSpecsRequest) SetResourceGroupId(v string) *ListServerIdeEcsSpecsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerIdeEcsSpecsRequest) Validate() error {
	return dara.Validate(s)
}
