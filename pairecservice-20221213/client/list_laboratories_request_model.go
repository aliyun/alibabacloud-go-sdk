// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLaboratoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *ListLaboratoriesRequest
	GetEnvironment() *string
	SetInstanceId(v string) *ListLaboratoriesRequest
	GetInstanceId() *string
	SetSceneId(v string) *ListLaboratoriesRequest
	GetSceneId() *string
	SetStatus(v string) *ListLaboratoriesRequest
	GetStatus() *string
}

type ListLaboratoriesRequest struct {
	// The laboratory environment.
	//
	// - `Daily`: the daily environment
	//
	// - `Pre`: the pre-production environment
	//
	// - `Prod`: the production environment
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The ID of the instance. You can obtain this ID by calling the `ListInstances` API.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the scene. You can obtain this ID by calling the `ListScenes` API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The laboratory status.
	//
	// - `Offline`: The laboratory is offline.
	//
	// - `Online`: The laboratory is online.
	//
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLaboratoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLaboratoriesRequest) GoString() string {
	return s.String()
}

func (s *ListLaboratoriesRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ListLaboratoriesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLaboratoriesRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListLaboratoriesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListLaboratoriesRequest) SetEnvironment(v string) *ListLaboratoriesRequest {
	s.Environment = &v
	return s
}

func (s *ListLaboratoriesRequest) SetInstanceId(v string) *ListLaboratoriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLaboratoriesRequest) SetSceneId(v string) *ListLaboratoriesRequest {
	s.SceneId = &v
	return s
}

func (s *ListLaboratoriesRequest) SetStatus(v string) *ListLaboratoriesRequest {
	s.Status = &v
	return s
}

func (s *ListLaboratoriesRequest) Validate() error {
	return dara.Validate(s)
}
