// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertRobotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAlertRobotsRequest
	GetName() *string
	SetPageNumber(v int64) *ListAlertRobotsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertRobotsRequest
	GetPageSize() *int64
	SetRobotIds(v []*string) *ListAlertRobotsRequest
	GetRobotIds() []*string
	SetTypes(v []*string) *ListAlertRobotsRequest
	GetTypes() []*string
	SetWorkspace(v string) *ListAlertRobotsRequest
	GetWorkspace() *string
}

type ListAlertRobotsRequest struct {
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
	PageSize  *int64    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RobotIds  []*string `json:"robotIds,omitempty" xml:"robotIds,omitempty" type:"Repeated"`
	Types     []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
	Workspace *string   `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListAlertRobotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRobotsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertRobotsRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertRobotsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertRobotsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertRobotsRequest) GetRobotIds() []*string {
	return s.RobotIds
}

func (s *ListAlertRobotsRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListAlertRobotsRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAlertRobotsRequest) SetName(v string) *ListAlertRobotsRequest {
	s.Name = &v
	return s
}

func (s *ListAlertRobotsRequest) SetPageNumber(v int64) *ListAlertRobotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlertRobotsRequest) SetPageSize(v int64) *ListAlertRobotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertRobotsRequest) SetRobotIds(v []*string) *ListAlertRobotsRequest {
	s.RobotIds = v
	return s
}

func (s *ListAlertRobotsRequest) SetTypes(v []*string) *ListAlertRobotsRequest {
	s.Types = v
	return s
}

func (s *ListAlertRobotsRequest) SetWorkspace(v string) *ListAlertRobotsRequest {
	s.Workspace = &v
	return s
}

func (s *ListAlertRobotsRequest) Validate() error {
	return dara.Validate(s)
}
