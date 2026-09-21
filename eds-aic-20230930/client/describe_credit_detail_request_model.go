// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentTypes(v []*string) *DescribeCreditDetailRequest
	GetAgentTypes() []*string
	SetEndTime(v int64) *DescribeCreditDetailRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribeCreditDetailRequest
	GetInstanceIds() []*string
	SetMaxResults(v int32) *DescribeCreditDetailRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCreditDetailRequest
	GetNextToken() *string
	SetPackageIds(v []*string) *DescribeCreditDetailRequest
	GetPackageIds() []*string
	SetPageNum(v string) *DescribeCreditDetailRequest
	GetPageNum() *string
	SetPageSize(v string) *DescribeCreditDetailRequest
	GetPageSize() *string
	SetStartTime(v int64) *DescribeCreditDetailRequest
	GetStartTime() *int64
}

type DescribeCreditDetailRequest struct {
	// The list of agent types, used to filter credit change details by specified agent types.
	AgentTypes []*string `json:"AgentTypes,omitempty" xml:"AgentTypes,omitempty" type:"Repeated"`
	// The end time.
	//
	// example:
	//
	// 1782906240000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The maximum number of entries to read in this request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. If a query does not return all results, the returned NextToken is not empty. You can pass the returned NextToken in the next query to continue retrieving results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of package or credit booster pack IDs.
	PackageIds []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// The page number for pagination. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1782819840000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCreditDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreditDetailRequest) GetAgentTypes() []*string {
	return s.AgentTypes
}

func (s *DescribeCreditDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeCreditDetailRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeCreditDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCreditDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCreditDetailRequest) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *DescribeCreditDetailRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *DescribeCreditDetailRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCreditDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeCreditDetailRequest) SetAgentTypes(v []*string) *DescribeCreditDetailRequest {
	s.AgentTypes = v
	return s
}

func (s *DescribeCreditDetailRequest) SetEndTime(v int64) *DescribeCreditDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCreditDetailRequest) SetInstanceIds(v []*string) *DescribeCreditDetailRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCreditDetailRequest) SetMaxResults(v int32) *DescribeCreditDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCreditDetailRequest) SetNextToken(v string) *DescribeCreditDetailRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCreditDetailRequest) SetPackageIds(v []*string) *DescribeCreditDetailRequest {
	s.PackageIds = v
	return s
}

func (s *DescribeCreditDetailRequest) SetPageNum(v string) *DescribeCreditDetailRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCreditDetailRequest) SetPageSize(v string) *DescribeCreditDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCreditDetailRequest) SetStartTime(v int64) *DescribeCreditDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCreditDetailRequest) Validate() error {
	return dara.Validate(s)
}
