// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListSceneRequest
	GetName() *string
	SetPageNum(v int64) *ListSceneRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListSceneRequest
	GetPageSize() *int64
	SetProjectId(v string) *ListSceneRequest
	GetProjectId() *string
}

type ListSceneRequest struct {
	// example:
	//
	// 厨房
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSceneRequest) GoString() string {
	return s.String()
}

func (s *ListSceneRequest) GetName() *string {
	return s.Name
}

func (s *ListSceneRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListSceneRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSceneRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListSceneRequest) SetName(v string) *ListSceneRequest {
	s.Name = &v
	return s
}

func (s *ListSceneRequest) SetPageNum(v int64) *ListSceneRequest {
	s.PageNum = &v
	return s
}

func (s *ListSceneRequest) SetPageSize(v int64) *ListSceneRequest {
	s.PageSize = &v
	return s
}

func (s *ListSceneRequest) SetProjectId(v string) *ListSceneRequest {
	s.ProjectId = &v
	return s
}

func (s *ListSceneRequest) Validate() error {
	return dara.Validate(s)
}
