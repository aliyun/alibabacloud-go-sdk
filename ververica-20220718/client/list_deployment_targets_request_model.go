// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *ListDeploymentTargetsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDeploymentTargetsRequest
	GetPageSize() *int32
}

type ListDeploymentTargetsRequest struct {
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDeploymentTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentTargetsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDeploymentTargetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentTargetsRequest) SetPageIndex(v int32) *ListDeploymentTargetsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentTargetsRequest) SetPageSize(v int32) *ListDeploymentTargetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentTargetsRequest) Validate() error {
	return dara.Validate(s)
}
