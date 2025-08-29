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
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
