// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDsType(v string) *ListDataSourceRequest
	GetDsType() *string
	SetWorkspaceId(v string) *ListDataSourceRequest
	GetWorkspaceId() *string
}

type ListDataSourceRequest struct {
	// Data source type.
	//
	// example:
	//
	// mysql
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// Workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-******c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceRequest) GetDsType() *string {
	return s.DsType
}

func (s *ListDataSourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataSourceRequest) SetDsType(v string) *ListDataSourceRequest {
	s.DsType = &v
	return s
}

func (s *ListDataSourceRequest) SetWorkspaceId(v string) *ListDataSourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
