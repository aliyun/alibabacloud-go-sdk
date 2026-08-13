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
	SetPageNumber(v int32) *ListCampaignsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCampaignsRequest
	GetPageSize() *int32
	SetPlannedStartTimeFrom(v string) *ListCampaignsRequest
	GetPlannedStartTimeFrom() *string
	SetPlannedStartTimeTo(v string) *ListCampaignsRequest
	GetPlannedStartTimeTo() *string
	SetState(v string) *ListCampaignsRequest
	GetState() *string
}

type ListCampaignsRequest struct {
	// The earliest actual start time.
	//
	// example:
	//
	// 1634054400000
	ActualStartTimeFrom *string `json:"ActualStartTimeFrom,omitempty" xml:"ActualStartTimeFrom,omitempty"`
	// The latest actual start time.
	//
	// example:
	//
	// 1634054400000
	ActualStartTimeTo *string `json:"ActualStartTimeTo,omitempty" xml:"ActualStartTimeTo,omitempty"`
	// The instance ID of the outbound robot.
	//
	// This parameter is required.
	//
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the campaign.
	//
	// example:
	//
	// SatisfactionSurvey
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The earliest planned start time.
	//
	// example:
	//
	// 1634054400000
	PlannedStartTimeFrom *string `json:"PlannedStartTimeFrom,omitempty" xml:"PlannedStartTimeFrom,omitempty"`
	// The latest planned start time.
	//
	// example:
	//
	// 1634054400000
	PlannedStartTimeTo *string `json:"PlannedStartTimeTo,omitempty" xml:"PlannedStartTimeTo,omitempty"`
	// The status of the campaign.
	//
	// example:
	//
	// Completed
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

func (s *ListCampaignsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCampaignsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCampaignsRequest) GetPlannedStartTimeFrom() *string {
	return s.PlannedStartTimeFrom
}

func (s *ListCampaignsRequest) GetPlannedStartTimeTo() *string {
	return s.PlannedStartTimeTo
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

func (s *ListCampaignsRequest) SetPageNumber(v int32) *ListCampaignsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsRequest) SetPageSize(v int32) *ListCampaignsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsRequest) SetPlannedStartTimeFrom(v string) *ListCampaignsRequest {
	s.PlannedStartTimeFrom = &v
	return s
}

func (s *ListCampaignsRequest) SetPlannedStartTimeTo(v string) *ListCampaignsRequest {
	s.PlannedStartTimeTo = &v
	return s
}

func (s *ListCampaignsRequest) SetState(v string) *ListCampaignsRequest {
	s.State = &v
	return s
}

func (s *ListCampaignsRequest) Validate() error {
	return dara.Validate(s)
}
