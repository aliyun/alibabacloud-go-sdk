// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceIds(v string) *ListInstanceStatisticsRequest
	GetWorkspaceIds() *string
}

type ListInstanceStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 27218,34956
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s ListInstanceStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListInstanceStatisticsRequest) SetWorkspaceIds(v string) *ListInstanceStatisticsRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListInstanceStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
