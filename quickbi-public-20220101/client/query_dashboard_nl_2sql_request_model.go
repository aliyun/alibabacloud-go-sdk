// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDashboardNl2sqlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QueryDashboardNl2sqlRequest
	GetUserId() *string
	SetWorkspaceId(v string) *QueryDashboardNl2sqlRequest
	GetWorkspaceId() *string
}

type QueryDashboardNl2sqlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3d7ebb8***********500078f4
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3d7ebb8***********500078f4
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryDashboardNl2sqlRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDashboardNl2sqlRequest) GoString() string {
	return s.String()
}

func (s *QueryDashboardNl2sqlRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryDashboardNl2sqlRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryDashboardNl2sqlRequest) SetUserId(v string) *QueryDashboardNl2sqlRequest {
	s.UserId = &v
	return s
}

func (s *QueryDashboardNl2sqlRequest) SetWorkspaceId(v string) *QueryDashboardNl2sqlRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryDashboardNl2sqlRequest) Validate() error {
	return dara.Validate(s)
}
