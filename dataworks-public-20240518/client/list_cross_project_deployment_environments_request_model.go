// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectDeploymentEnvironmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCrossProjectDeploymentEnvironmentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCrossProjectDeploymentEnvironmentsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCrossProjectDeploymentEnvironmentsRequest
	GetProjectId() *int64
}

type ListCrossProjectDeploymentEnvironmentsRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListCrossProjectDeploymentEnvironmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentEnvironmentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrossProjectDeploymentEnvironmentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrossProjectDeploymentEnvironmentsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCrossProjectDeploymentEnvironmentsRequest) SetPageNumber(v int32) *ListCrossProjectDeploymentEnvironmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsRequest) SetPageSize(v int32) *ListCrossProjectDeploymentEnvironmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsRequest) SetProjectId(v int64) *ListCrossProjectDeploymentEnvironmentsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsRequest) Validate() error {
	return dara.Validate(s)
}
