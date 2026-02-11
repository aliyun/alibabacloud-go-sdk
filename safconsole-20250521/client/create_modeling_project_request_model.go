// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceSpec(v string) *CreateModelingProjectRequest
	GetInstanceSpec() *string
	SetProjectName(v string) *CreateModelingProjectRequest
	GetProjectName() *string
	SetRemark(v string) *CreateModelingProjectRequest
	GetRemark() *string
}

type CreateModelingProjectRequest struct {
	// Instance specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// SECURE_ENV_LITE
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// Project name
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Remark.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateModelingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelingProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateModelingProjectRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *CreateModelingProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateModelingProjectRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateModelingProjectRequest) SetInstanceSpec(v string) *CreateModelingProjectRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateModelingProjectRequest) SetProjectName(v string) *CreateModelingProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateModelingProjectRequest) SetRemark(v string) *CreateModelingProjectRequest {
	s.Remark = &v
	return s
}

func (s *CreateModelingProjectRequest) Validate() error {
	return dara.Validate(s)
}
