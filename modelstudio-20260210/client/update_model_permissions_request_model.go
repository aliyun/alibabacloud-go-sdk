// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessAllEntities(v string) *UpdateModelPermissionsRequest
	GetAccessAllEntities() *string
	SetModels(v []*UpdateModelPermissionsRequestModels) *UpdateModelPermissionsRequest
	GetModels() []*UpdateModelPermissionsRequestModels
	SetWorkspaceId(v string) *UpdateModelPermissionsRequest
	GetWorkspaceId() *string
}

type UpdateModelPermissionsRequest struct {
	// example:
	//
	// OPEN
	AccessAllEntities *string                                `json:"accessAllEntities,omitempty" xml:"accessAllEntities,omitempty"`
	Models            []*UpdateModelPermissionsRequestModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ws-32klhjk2312334jkh
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateModelPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelPermissionsRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelPermissionsRequest) GetAccessAllEntities() *string {
	return s.AccessAllEntities
}

func (s *UpdateModelPermissionsRequest) GetModels() []*UpdateModelPermissionsRequestModels {
	return s.Models
}

func (s *UpdateModelPermissionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateModelPermissionsRequest) SetAccessAllEntities(v string) *UpdateModelPermissionsRequest {
	s.AccessAllEntities = &v
	return s
}

func (s *UpdateModelPermissionsRequest) SetModels(v []*UpdateModelPermissionsRequestModels) *UpdateModelPermissionsRequest {
	s.Models = v
	return s
}

func (s *UpdateModelPermissionsRequest) SetWorkspaceId(v string) *UpdateModelPermissionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateModelPermissionsRequest) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateModelPermissionsRequestModels struct {
	Deploy    *bool `json:"deploy,omitempty" xml:"deploy,omitempty"`
	FineTune  *bool `json:"fineTune,omitempty" xml:"fineTune,omitempty"`
	Inference *bool `json:"inference,omitempty" xml:"inference,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
}

func (s UpdateModelPermissionsRequestModels) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelPermissionsRequestModels) GoString() string {
	return s.String()
}

func (s *UpdateModelPermissionsRequestModels) GetDeploy() *bool {
	return s.Deploy
}

func (s *UpdateModelPermissionsRequestModels) GetFineTune() *bool {
	return s.FineTune
}

func (s *UpdateModelPermissionsRequestModels) GetInference() *bool {
	return s.Inference
}

func (s *UpdateModelPermissionsRequestModels) GetModel() *string {
	return s.Model
}

func (s *UpdateModelPermissionsRequestModels) SetDeploy(v bool) *UpdateModelPermissionsRequestModels {
	s.Deploy = &v
	return s
}

func (s *UpdateModelPermissionsRequestModels) SetFineTune(v bool) *UpdateModelPermissionsRequestModels {
	s.FineTune = &v
	return s
}

func (s *UpdateModelPermissionsRequestModels) SetInference(v bool) *UpdateModelPermissionsRequestModels {
	s.Inference = &v
	return s
}

func (s *UpdateModelPermissionsRequestModels) SetModel(v string) *UpdateModelPermissionsRequestModels {
	s.Model = &v
	return s
}

func (s *UpdateModelPermissionsRequestModels) Validate() error {
	return dara.Validate(s)
}
