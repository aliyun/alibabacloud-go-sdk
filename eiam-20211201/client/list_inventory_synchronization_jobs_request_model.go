// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventorySynchronizationJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListInventorySynchronizationJobsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListInventorySynchronizationJobsRequest
	GetInstanceId() *string
	SetInventorySourceIds(v []*string) *ListInventorySynchronizationJobsRequest
	GetInventorySourceIds() []*string
	SetMaxResults(v int32) *ListInventorySynchronizationJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInventorySynchronizationJobsRequest
	GetNextToken() *string
	SetStartTime(v int64) *ListInventorySynchronizationJobsRequest
	GetStartTime() *int64
	SetStatus(v string) *ListInventorySynchronizationJobsRequest
	GetStatus() *string
}

type ListInventorySynchronizationJobsRequest struct {
	// The end time.
	//
	// example:
	//
	// 1234567890
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the synchronization targets.
	InventorySourceIds []*string `json:"InventorySourceIds,omitempty" xml:"InventorySourceIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return per page in a paging query. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. The value is the NextToken parameter value returned from the previous API call.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1234567890
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The synchronization status.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInventorySynchronizationJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsRequest) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListInventorySynchronizationJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInventorySynchronizationJobsRequest) GetInventorySourceIds() []*string {
	return s.InventorySourceIds
}

func (s *ListInventorySynchronizationJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInventorySynchronizationJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInventorySynchronizationJobsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListInventorySynchronizationJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInventorySynchronizationJobsRequest) SetEndTime(v int64) *ListInventorySynchronizationJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListInventorySynchronizationJobsRequest) SetInstanceId(v string) *ListInventorySynchronizationJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInventorySynchronizationJobsRequest) SetInventorySourceIds(v []*string) *ListInventorySynchronizationJobsRequest {
	s.InventorySourceIds = v
	return s
}

func (s *ListInventorySynchronizationJobsRequest) SetMaxResults(v int32) *ListInventorySynchronizationJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInventorySynchronizationJobsRequest) SetNextToken(v string) *ListInventorySynchronizationJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListInventorySynchronizationJobsRequest) SetStartTime(v int64) *ListInventorySynchronizationJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListInventorySynchronizationJobsRequest) SetStatus(v string) *ListInventorySynchronizationJobsRequest {
	s.Status = &v
	return s
}

func (s *ListInventorySynchronizationJobsRequest) Validate() error {
	return dara.Validate(s)
}
