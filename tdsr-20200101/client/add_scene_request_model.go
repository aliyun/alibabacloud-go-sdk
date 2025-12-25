// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerUid(v string) *AddSceneRequest
	GetCustomerUid() *string
	SetName(v string) *AddSceneRequest
	GetName() *string
	SetProjectId(v string) *AddSceneRequest
	GetProjectId() *string
	SetType(v string) *AddSceneRequest
	GetType() *string
}

type AddSceneRequest struct {
	// example:
	//
	// 2345****
	CustomerUid *string `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MODEL_3D
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSceneRequest) GoString() string {
	return s.String()
}

func (s *AddSceneRequest) GetCustomerUid() *string {
	return s.CustomerUid
}

func (s *AddSceneRequest) GetName() *string {
	return s.Name
}

func (s *AddSceneRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *AddSceneRequest) GetType() *string {
	return s.Type
}

func (s *AddSceneRequest) SetCustomerUid(v string) *AddSceneRequest {
	s.CustomerUid = &v
	return s
}

func (s *AddSceneRequest) SetName(v string) *AddSceneRequest {
	s.Name = &v
	return s
}

func (s *AddSceneRequest) SetProjectId(v string) *AddSceneRequest {
	s.ProjectId = &v
	return s
}

func (s *AddSceneRequest) SetType(v string) *AddSceneRequest {
	s.Type = &v
	return s
}

func (s *AddSceneRequest) Validate() error {
	return dara.Validate(s)
}
