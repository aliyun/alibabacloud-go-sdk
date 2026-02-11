// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployModelFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DeployModelFileRequest
	GetGroupId() *int64
	SetProjectId(v int64) *DeployModelFileRequest
	GetProjectId() *int64
}

type DeployModelFileRequest struct {
	// File Group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeployModelFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployModelFileRequest) GoString() string {
	return s.String()
}

func (s *DeployModelFileRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeployModelFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeployModelFileRequest) SetGroupId(v int64) *DeployModelFileRequest {
	s.GroupId = &v
	return s
}

func (s *DeployModelFileRequest) SetProjectId(v int64) *DeployModelFileRequest {
	s.ProjectId = &v
	return s
}

func (s *DeployModelFileRequest) Validate() error {
	return dara.Validate(s)
}
