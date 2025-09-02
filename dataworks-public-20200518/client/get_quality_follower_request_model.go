// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityFollowerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *GetQualityFollowerRequest
	GetEntityId() *int64
	SetProjectId(v int64) *GetQualityFollowerRequest
	GetProjectId() *int64
	SetProjectName(v string) *GetQualityFollowerRequest
	GetProjectName() *string
}

type GetQualityFollowerRequest struct {
	// The ID of the partition filter expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetQualityFollowerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityFollowerRequest) GoString() string {
	return s.String()
}

func (s *GetQualityFollowerRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *GetQualityFollowerRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityFollowerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityFollowerRequest) SetEntityId(v int64) *GetQualityFollowerRequest {
	s.EntityId = &v
	return s
}

func (s *GetQualityFollowerRequest) SetProjectId(v int64) *GetQualityFollowerRequest {
	s.ProjectId = &v
	return s
}

func (s *GetQualityFollowerRequest) SetProjectName(v string) *GetQualityFollowerRequest {
	s.ProjectName = &v
	return s
}

func (s *GetQualityFollowerRequest) Validate() error {
	return dara.Validate(s)
}
