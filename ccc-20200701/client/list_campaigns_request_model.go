// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCampaignsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActualStartTimeFrom(v string) *ListCampaignsRequest
	GetActualStartTimeFrom() *string
	SetActualStartTimeTo(v string) *ListCampaignsRequest
	GetActualStartTimeTo() *string
	SetInstanceId(v string) *ListCampaignsRequest
	GetInstanceId() *string
	SetName(v string) *ListCampaignsRequest
	GetName() *string
	SetPageNumber(v int64) *ListCampaignsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCampaignsRequest
	GetPageSize() *int64
	SetPlanedStartTimeFrom(v string) *ListCampaignsRequest
	GetPlanedStartTimeFrom() *string
	SetPlanedStartTimeTo(v string) *ListCampaignsRequest
	GetPlanedStartTimeTo() *string
	SetQueueId(v string) *ListCampaignsRequest
	GetQueueId() *string
	SetState(v string) *ListCampaignsRequest
	GetState() *string
}

type ListCampaignsRequest struct {
	// example:
	//
	// 2021-10-14 20:59:59
	ActualStartTimeFrom *string `json:"ActualStartTimeFrom,omitempty" xml:"ActualStartTimeFrom,omitempty"`
	// example:
	//
	// 2021-10-14 20:59:59
	ActualStartTimeTo *string `json:"ActualStartTimeTo,omitempty" xml:"ActualStartTimeTo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test-campaign
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2021-10-14 00:00:00
	PlanedStartTimeFrom *string `json:"PlanedStartTimeFrom,omitempty" xml:"PlanedStartTimeFrom,omitempty"`
	// example:
	//
	// 2021-10-14 20:59:59
	PlanedStartTimeTo *string `json:"PlanedStartTimeTo,omitempty" xml:"PlanedStartTimeTo,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// example:
	//
	// Draft
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListCampaignsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignsRequest) GoString() string {
	return s.String()
}

func (s *ListCampaignsRequest) GetActualStartTimeFrom() *string {
	return s.ActualStartTimeFrom
}

func (s *ListCampaignsRequest) GetActualStartTimeTo() *string {
	return s.ActualStartTimeTo
}

func (s *ListCampaignsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCampaignsRequest) GetName() *string {
	return s.Name
}

func (s *ListCampaignsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCampaignsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCampaignsRequest) GetPlanedStartTimeFrom() *string {
	return s.PlanedStartTimeFrom
}

func (s *ListCampaignsRequest) GetPlanedStartTimeTo() *string {
	return s.PlanedStartTimeTo
}

func (s *ListCampaignsRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *ListCampaignsRequest) GetState() *string {
	return s.State
}

func (s *ListCampaignsRequest) SetActualStartTimeFrom(v string) *ListCampaignsRequest {
	s.ActualStartTimeFrom = &v
	return s
}

func (s *ListCampaignsRequest) SetActualStartTimeTo(v string) *ListCampaignsRequest {
	s.ActualStartTimeTo = &v
	return s
}

func (s *ListCampaignsRequest) SetInstanceId(v string) *ListCampaignsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCampaignsRequest) SetName(v string) *ListCampaignsRequest {
	s.Name = &v
	return s
}

func (s *ListCampaignsRequest) SetPageNumber(v int64) *ListCampaignsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsRequest) SetPageSize(v int64) *ListCampaignsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsRequest) SetPlanedStartTimeFrom(v string) *ListCampaignsRequest {
	s.PlanedStartTimeFrom = &v
	return s
}

func (s *ListCampaignsRequest) SetPlanedStartTimeTo(v string) *ListCampaignsRequest {
	s.PlanedStartTimeTo = &v
	return s
}

func (s *ListCampaignsRequest) SetQueueId(v string) *ListCampaignsRequest {
	s.QueueId = &v
	return s
}

func (s *ListCampaignsRequest) SetState(v string) *ListCampaignsRequest {
	s.State = &v
	return s
}

func (s *ListCampaignsRequest) Validate() error {
	return dara.Validate(s)
}
