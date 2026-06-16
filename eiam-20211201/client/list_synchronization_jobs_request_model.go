// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSynchronizationJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListSynchronizationJobsRequest
	GetDirection() *string
	SetEndTime(v int64) *ListSynchronizationJobsRequest
	GetEndTime() *int64
	SetFilters(v []*ListSynchronizationJobsRequestFilters) *ListSynchronizationJobsRequest
	GetFilters() []*ListSynchronizationJobsRequestFilters
	SetInstanceId(v string) *ListSynchronizationJobsRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListSynchronizationJobsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListSynchronizationJobsRequest
	GetNextToken() *string
	SetPageNumber(v int64) *ListSynchronizationJobsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSynchronizationJobsRequest
	GetPageSize() *int64
	SetStartTime(v int64) *ListSynchronizationJobsRequest
	GetStartTime() *int64
	SetStatus(v string) *ListSynchronizationJobsRequest
	GetStatus() *string
	SetTargetIds(v []*string) *ListSynchronizationJobsRequest
	GetTargetIds() []*string
	SetTargetType(v string) *ListSynchronizationJobsRequest
	GetTargetType() *string
}

type ListSynchronizationJobsRequest struct {
	// The direction of the sync task. Valid values:
	//
	// - ingress: Inbound.
	//
	// - egress: Outbound.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The synchronization end time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter parameters.
	Filters []*ListSynchronizationJobsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. If no more pages exist, this parameter is not returned.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. The value starts from 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The maximum value is 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The synchronization start time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the sync task. Valid values:
	//
	// - pending: The task is pending.
	//
	// - running: The task is running.
	//
	// - failed: The task failed.
	//
	// - partial_success: The task is partially successful.
	//
	// - success: The task is successful.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of synchronization target IDs. For example, \\`[idp_111XXXX,idp_222XXXX]\\`.
	//
	// example:
	//
	// target_001
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
	// The type of the synchronization target. Valid values:
	//
	// - identity_provider: Identity provider.
	//
	// - application: Application.
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListSynchronizationJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListSynchronizationJobsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListSynchronizationJobsRequest) GetFilters() []*ListSynchronizationJobsRequestFilters {
	return s.Filters
}

func (s *ListSynchronizationJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSynchronizationJobsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListSynchronizationJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSynchronizationJobsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSynchronizationJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSynchronizationJobsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListSynchronizationJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSynchronizationJobsRequest) GetTargetIds() []*string {
	return s.TargetIds
}

func (s *ListSynchronizationJobsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListSynchronizationJobsRequest) SetDirection(v string) *ListSynchronizationJobsRequest {
	s.Direction = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetEndTime(v int64) *ListSynchronizationJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetFilters(v []*ListSynchronizationJobsRequestFilters) *ListSynchronizationJobsRequest {
	s.Filters = v
	return s
}

func (s *ListSynchronizationJobsRequest) SetInstanceId(v string) *ListSynchronizationJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetMaxResults(v int64) *ListSynchronizationJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetNextToken(v string) *ListSynchronizationJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetPageNumber(v int64) *ListSynchronizationJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetPageSize(v int64) *ListSynchronizationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetStartTime(v int64) *ListSynchronizationJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetStatus(v string) *ListSynchronizationJobsRequest {
	s.Status = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetTargetIds(v []*string) *ListSynchronizationJobsRequest {
	s.TargetIds = v
	return s
}

func (s *ListSynchronizationJobsRequest) SetTargetType(v string) *ListSynchronizationJobsRequest {
	s.TargetType = &v
	return s
}

func (s *ListSynchronizationJobsRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSynchronizationJobsRequestFilters struct {
	// The name of the dynamic parameter.
	//
	// example:
	//
	// qps
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the dynamic parameter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListSynchronizationJobsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListSynchronizationJobsRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *ListSynchronizationJobsRequestFilters) SetKey(v string) *ListSynchronizationJobsRequestFilters {
	s.Key = &v
	return s
}

func (s *ListSynchronizationJobsRequestFilters) SetValues(v []*string) *ListSynchronizationJobsRequestFilters {
	s.Values = v
	return s
}

func (s *ListSynchronizationJobsRequestFilters) Validate() error {
	return dara.Validate(s)
}
