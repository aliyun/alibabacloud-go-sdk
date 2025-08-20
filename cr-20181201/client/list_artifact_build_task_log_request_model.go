// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactBuildTaskLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildTaskId(v string) *ListArtifactBuildTaskLogRequest
	GetBuildTaskId() *string
	SetInstanceId(v string) *ListArtifactBuildTaskLogRequest
	GetInstanceId() *string
	SetPage(v int32) *ListArtifactBuildTaskLogRequest
	GetPage() *int32
	SetPageSize(v int32) *ListArtifactBuildTaskLogRequest
	GetPageSize() *int32
}

type ListArtifactBuildTaskLogRequest struct {
	// The ID of the artifact build task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i2a-1yu****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page. Maximum value: 100. If you specify a value greater than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListArtifactBuildTaskLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildTaskLogRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildTaskLogRequest) GetBuildTaskId() *string {
	return s.BuildTaskId
}

func (s *ListArtifactBuildTaskLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListArtifactBuildTaskLogRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListArtifactBuildTaskLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListArtifactBuildTaskLogRequest) SetBuildTaskId(v string) *ListArtifactBuildTaskLogRequest {
	s.BuildTaskId = &v
	return s
}

func (s *ListArtifactBuildTaskLogRequest) SetInstanceId(v string) *ListArtifactBuildTaskLogRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactBuildTaskLogRequest) SetPage(v int32) *ListArtifactBuildTaskLogRequest {
	s.Page = &v
	return s
}

func (s *ListArtifactBuildTaskLogRequest) SetPageSize(v int32) *ListArtifactBuildTaskLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListArtifactBuildTaskLogRequest) Validate() error {
	return dara.Validate(s)
}
