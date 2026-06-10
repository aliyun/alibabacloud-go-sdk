// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalTimerBatchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeGlobalTimerBatchesRequest
	GetGroupId() *string
	SetMaxResults(v string) *DescribeGlobalTimerBatchesRequest
	GetMaxResults() *string
	SetNextToken(v string) *DescribeGlobalTimerBatchesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeGlobalTimerBatchesRequest
	GetRegionId() *string
	SetResourceTypes(v []*string) *DescribeGlobalTimerBatchesRequest
	GetResourceTypes() []*string
	SetSearchRegionId(v string) *DescribeGlobalTimerBatchesRequest
	GetSearchRegionId() *string
	SetTimerType(v string) *DescribeGlobalTimerBatchesRequest
	GetTimerType() *string
}

type DescribeGlobalTimerBatchesRequest struct {
	// The ID of the scheduled task group.
	//
	// example:
	//
	// ccg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. Set this parameter to the `NextToken` value from a previous response.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// - China (Shanghai)
	//
	// - Singapore (Singapore)
	//
	// example:
	//
	// cn-shanghai
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The ID of the region where the cloud computers are located. This parameter filters the results to include only cloud computers in the specified region.
	//
	// example:
	//
	// cn-hangzhou
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// The type of the scheduled task. This operation returns batch information for timer-based scheduled tasks only.
	//
	// - `TimerBoot`: scheduled startup
	//
	// - `TimerShutdown`: scheduled shutdown
	//
	// - `TimerReboot`: scheduled reboot
	//
	// - `TimerReset`: scheduled reset
	//
	// - `TimerMaintenance`: scheduled maintenance
	//
	// - `TimerHibernate`: scheduled hibernation
	//
	// example:
	//
	// 1
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeGlobalTimerBatchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerBatchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerBatchesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGlobalTimerBatchesRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeGlobalTimerBatchesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGlobalTimerBatchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalTimerBatchesRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribeGlobalTimerBatchesRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeGlobalTimerBatchesRequest) GetTimerType() *string {
	return s.TimerType
}

func (s *DescribeGlobalTimerBatchesRequest) SetGroupId(v string) *DescribeGlobalTimerBatchesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetMaxResults(v string) *DescribeGlobalTimerBatchesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetNextToken(v string) *DescribeGlobalTimerBatchesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetRegionId(v string) *DescribeGlobalTimerBatchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetResourceTypes(v []*string) *DescribeGlobalTimerBatchesRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetSearchRegionId(v string) *DescribeGlobalTimerBatchesRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) SetTimerType(v string) *DescribeGlobalTimerBatchesRequest {
	s.TimerType = &v
	return s
}

func (s *DescribeGlobalTimerBatchesRequest) Validate() error {
	return dara.Validate(s)
}
