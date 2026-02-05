// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateInspectionTaskRequest
	GetEndTime() *string
	SetInspectionItems(v string) *CreateInspectionTaskRequest
	GetInspectionItems() *string
	SetInstanceIds(v string) *CreateInspectionTaskRequest
	GetInstanceIds() *string
	SetStartTime(v string) *CreateInspectionTaskRequest
	GetStartTime() *string
}

type CreateInspectionTaskRequest struct {
	// example:
	//
	// 2026-01-30T02:10:48Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// instance_info, resource_usage
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 2025-12-28T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInspectionTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateInspectionTaskRequest) GetInspectionItems() *string {
	return s.InspectionItems
}

func (s *CreateInspectionTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *CreateInspectionTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateInspectionTaskRequest) SetEndTime(v string) *CreateInspectionTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetInspectionItems(v string) *CreateInspectionTaskRequest {
	s.InspectionItems = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetInstanceIds(v string) *CreateInspectionTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetStartTime(v string) *CreateInspectionTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
