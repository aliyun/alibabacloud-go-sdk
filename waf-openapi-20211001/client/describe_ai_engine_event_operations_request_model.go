// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeAiEngineEventOperationsRequest
	GetEventId() *string
	SetInstanceId(v string) *DescribeAiEngineEventOperationsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeAiEngineEventOperationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEventOperationsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeAiEngineEventOperationsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventOperationsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeAiEngineEventOperationsRequest struct {
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EVT-20241025***109
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_dddsdf3*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page during paging in a paged query. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set the value to the **NextToken*	- value that is returned from the previous API call. You do not need to specify this parameter for the first query.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the WAF instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-de****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeAiEngineEventOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventOperationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventOperationsRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEventOperationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEventOperationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEventOperationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEventOperationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEventOperationsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEventOperationsRequest) SetEventId(v string) *DescribeAiEngineEventOperationsRequest {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEventOperationsRequest) SetInstanceId(v string) *DescribeAiEngineEventOperationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEventOperationsRequest) SetMaxResults(v int32) *DescribeAiEngineEventOperationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEventOperationsRequest) SetNextToken(v string) *DescribeAiEngineEventOperationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEventOperationsRequest) SetRegionId(v string) *DescribeAiEngineEventOperationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEventOperationsRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventOperationsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEventOperationsRequest) Validate() error {
	return dara.Validate(s)
}
