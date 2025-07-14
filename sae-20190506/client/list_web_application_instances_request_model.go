// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListWebApplicationInstancesRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *ListWebApplicationInstancesRequest
	GetInstanceIds() []*string
	SetLimit(v string) *ListWebApplicationInstancesRequest
	GetLimit() *string
	SetNamespaceId(v string) *ListWebApplicationInstancesRequest
	GetNamespaceId() *string
	SetStartTime(v int64) *ListWebApplicationInstancesRequest
	GetStartTime() *int64
	SetStatuses(v []*string) *ListWebApplicationInstancesRequest
	GetStatuses() []*string
	SetVersionIds(v []*string) *ListWebApplicationInstancesRequest
	GetVersionIds() []*string
}

type ListWebApplicationInstancesRequest struct {
	// The time when the operation ended.
	//
	// example:
	//
	// 1715567192
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// c-667d143a-17b4e0fa-46d3a2******
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of application instances returned.
	//
	// example:
	//
	// 10
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 1562831689704
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the application instance.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// The ID of the application version.
	//
	// example:
	//
	// 001
	VersionIds []*string `json:"VersionIds,omitempty" xml:"VersionIds,omitempty" type:"Repeated"`
}

func (s ListWebApplicationInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListWebApplicationInstancesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListWebApplicationInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListWebApplicationInstancesRequest) GetLimit() *string {
	return s.Limit
}

func (s *ListWebApplicationInstancesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListWebApplicationInstancesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListWebApplicationInstancesRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListWebApplicationInstancesRequest) GetVersionIds() []*string {
	return s.VersionIds
}

func (s *ListWebApplicationInstancesRequest) SetEndTime(v int64) *ListWebApplicationInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListWebApplicationInstancesRequest) SetInstanceIds(v []*string) *ListWebApplicationInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *ListWebApplicationInstancesRequest) SetLimit(v string) *ListWebApplicationInstancesRequest {
	s.Limit = &v
	return s
}

func (s *ListWebApplicationInstancesRequest) SetNamespaceId(v string) *ListWebApplicationInstancesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListWebApplicationInstancesRequest) SetStartTime(v int64) *ListWebApplicationInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListWebApplicationInstancesRequest) SetStatuses(v []*string) *ListWebApplicationInstancesRequest {
	s.Statuses = v
	return s
}

func (s *ListWebApplicationInstancesRequest) SetVersionIds(v []*string) *ListWebApplicationInstancesRequest {
	s.VersionIds = v
	return s
}

func (s *ListWebApplicationInstancesRequest) Validate() error {
	return dara.Validate(s)
}
