// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *GetEventsRequest
	GetDeploymentId() *string
	SetPageIndex(v int32) *GetEventsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetEventsRequest
	GetPageSize() *int32
}

type GetEventsRequest struct {
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEventsRequest) GoString() string {
	return s.String()
}

func (s *GetEventsRequest) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *GetEventsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetEventsRequest) SetDeploymentId(v string) *GetEventsRequest {
	s.DeploymentId = &v
	return s
}

func (s *GetEventsRequest) SetPageIndex(v int32) *GetEventsRequest {
	s.PageIndex = &v
	return s
}

func (s *GetEventsRequest) SetPageSize(v int32) *GetEventsRequest {
	s.PageSize = &v
	return s
}

func (s *GetEventsRequest) Validate() error {
	return dara.Validate(s)
}
