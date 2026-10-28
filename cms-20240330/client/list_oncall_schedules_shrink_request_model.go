// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOncallSchedulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOncallScheduleIdsShrink(v string) *ListOncallSchedulesShrinkRequest
	GetOncallScheduleIdsShrink() *string
	SetOncallScheduleName(v string) *ListOncallSchedulesShrinkRequest
	GetOncallScheduleName() *string
	SetPageNumber(v int32) *ListOncallSchedulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOncallSchedulesShrinkRequest
	GetPageSize() *int32
	SetWorkspace(v string) *ListOncallSchedulesShrinkRequest
	GetWorkspace() *string
}

type ListOncallSchedulesShrinkRequest struct {
	OncallScheduleIdsShrink *string `json:"oncallScheduleIds,omitempty" xml:"oncallScheduleIds,omitempty"`
	// example:
	//
	// testOncallScheduleName
	OncallScheduleName *string `json:"oncallScheduleName,omitempty" xml:"oncallScheduleName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// default-cms-5884500827389836-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListOncallSchedulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOncallSchedulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOncallSchedulesShrinkRequest) GetOncallScheduleIdsShrink() *string {
	return s.OncallScheduleIdsShrink
}

func (s *ListOncallSchedulesShrinkRequest) GetOncallScheduleName() *string {
	return s.OncallScheduleName
}

func (s *ListOncallSchedulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOncallSchedulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOncallSchedulesShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListOncallSchedulesShrinkRequest) SetOncallScheduleIdsShrink(v string) *ListOncallSchedulesShrinkRequest {
	s.OncallScheduleIdsShrink = &v
	return s
}

func (s *ListOncallSchedulesShrinkRequest) SetOncallScheduleName(v string) *ListOncallSchedulesShrinkRequest {
	s.OncallScheduleName = &v
	return s
}

func (s *ListOncallSchedulesShrinkRequest) SetPageNumber(v int32) *ListOncallSchedulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOncallSchedulesShrinkRequest) SetPageSize(v int32) *ListOncallSchedulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListOncallSchedulesShrinkRequest) SetWorkspace(v string) *ListOncallSchedulesShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ListOncallSchedulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
