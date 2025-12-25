// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateSceneRequest
	GetId() *string
	SetName(v string) *UpdateSceneRequest
	GetName() *string
}

type UpdateSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSceneRequest) GetId() *string {
	return s.Id
}

func (s *UpdateSceneRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSceneRequest) SetId(v string) *UpdateSceneRequest {
	s.Id = &v
	return s
}

func (s *UpdateSceneRequest) SetName(v string) *UpdateSceneRequest {
	s.Name = &v
	return s
}

func (s *UpdateSceneRequest) Validate() error {
	return dara.Validate(s)
}
