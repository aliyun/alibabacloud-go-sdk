// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAirflowVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *ListAirflowVersionsRequest
	GetWorkspaceId() *string
}

type ListAirflowVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAirflowVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAirflowVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAirflowVersionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAirflowVersionsRequest) SetWorkspaceId(v string) *ListAirflowVersionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAirflowVersionsRequest) Validate() error {
	return dara.Validate(s)
}
