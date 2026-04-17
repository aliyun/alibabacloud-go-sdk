// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertRobotsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAlertRobotsShrinkRequest
	GetName() *string
	SetPageNumber(v int64) *ListAlertRobotsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertRobotsShrinkRequest
	GetPageSize() *int64
	SetRobotIdsShrink(v string) *ListAlertRobotsShrinkRequest
	GetRobotIdsShrink() *string
	SetTypesShrink(v string) *ListAlertRobotsShrinkRequest
	GetTypesShrink() *string
	SetWorkspace(v string) *ListAlertRobotsShrinkRequest
	GetWorkspace() *string
}

type ListAlertRobotsShrinkRequest struct {
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize       *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RobotIdsShrink *string `json:"robotIds,omitempty" xml:"robotIds,omitempty"`
	TypesShrink    *string `json:"types,omitempty" xml:"types,omitempty"`
	Workspace      *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListAlertRobotsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRobotsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertRobotsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertRobotsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertRobotsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertRobotsShrinkRequest) GetRobotIdsShrink() *string {
	return s.RobotIdsShrink
}

func (s *ListAlertRobotsShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListAlertRobotsShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAlertRobotsShrinkRequest) SetName(v string) *ListAlertRobotsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListAlertRobotsShrinkRequest) SetPageNumber(v int64) *ListAlertRobotsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertRobotsShrinkRequest) SetPageSize(v int64) *ListAlertRobotsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertRobotsShrinkRequest) SetRobotIdsShrink(v string) *ListAlertRobotsShrinkRequest {
	s.RobotIdsShrink = &v
	return s
}

func (s *ListAlertRobotsShrinkRequest) SetTypesShrink(v string) *ListAlertRobotsShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListAlertRobotsShrinkRequest) SetWorkspace(v string) *ListAlertRobotsShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ListAlertRobotsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
