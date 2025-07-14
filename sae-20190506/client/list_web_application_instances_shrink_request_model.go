// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListWebApplicationInstancesShrinkRequest
	GetEndTime() *int64
	SetInstanceIdsShrink(v string) *ListWebApplicationInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetLimit(v string) *ListWebApplicationInstancesShrinkRequest
	GetLimit() *string
	SetNamespaceId(v string) *ListWebApplicationInstancesShrinkRequest
	GetNamespaceId() *string
	SetStartTime(v int64) *ListWebApplicationInstancesShrinkRequest
	GetStartTime() *int64
	SetStatusesShrink(v string) *ListWebApplicationInstancesShrinkRequest
	GetStatusesShrink() *string
	SetVersionIdsShrink(v string) *ListWebApplicationInstancesShrinkRequest
	GetVersionIdsShrink() *string
}

type ListWebApplicationInstancesShrinkRequest struct {
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
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
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
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
	// The ID of the application version.
	//
	// example:
	//
	// 001
	VersionIdsShrink *string `json:"VersionIds,omitempty" xml:"VersionIds,omitempty"`
}

func (s ListWebApplicationInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWebApplicationInstancesShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListWebApplicationInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ListWebApplicationInstancesShrinkRequest) GetLimit() *string {
	return s.Limit
}

func (s *ListWebApplicationInstancesShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListWebApplicationInstancesShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListWebApplicationInstancesShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListWebApplicationInstancesShrinkRequest) GetVersionIdsShrink() *string {
	return s.VersionIdsShrink
}

func (s *ListWebApplicationInstancesShrinkRequest) SetEndTime(v int64) *ListWebApplicationInstancesShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListWebApplicationInstancesShrinkRequest) SetInstanceIdsShrink(v string) *ListWebApplicationInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ListWebApplicationInstancesShrinkRequest) SetLimit(v string) *ListWebApplicationInstancesShrinkRequest {
	s.Limit = &v
	return s
}

func (s *ListWebApplicationInstancesShrinkRequest) SetNamespaceId(v string) *ListWebApplicationInstancesShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListWebApplicationInstancesShrinkRequest) SetStartTime(v int64) *ListWebApplicationInstancesShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListWebApplicationInstancesShrinkRequest) SetStatusesShrink(v string) *ListWebApplicationInstancesShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListWebApplicationInstancesShrinkRequest) SetVersionIdsShrink(v string) *ListWebApplicationInstancesShrinkRequest {
	s.VersionIdsShrink = &v
	return s
}

func (s *ListWebApplicationInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
