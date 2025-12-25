// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateSubSceneRequest
	GetId() *string
	SetName(v string) *UpdateSubSceneRequest
	GetName() *string
	SetViewPoint(v []*float64) *UpdateSubSceneRequest
	GetViewPoint() []*float64
}

type UpdateSubSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 测试
	Name      *string    `json:"Name,omitempty" xml:"Name,omitempty"`
	ViewPoint []*float64 `json:"ViewPoint,omitempty" xml:"ViewPoint,omitempty" type:"Repeated"`
}

func (s UpdateSubSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneRequest) GetId() *string {
	return s.Id
}

func (s *UpdateSubSceneRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSubSceneRequest) GetViewPoint() []*float64 {
	return s.ViewPoint
}

func (s *UpdateSubSceneRequest) SetId(v string) *UpdateSubSceneRequest {
	s.Id = &v
	return s
}

func (s *UpdateSubSceneRequest) SetName(v string) *UpdateSubSceneRequest {
	s.Name = &v
	return s
}

func (s *UpdateSubSceneRequest) SetViewPoint(v []*float64) *UpdateSubSceneRequest {
	s.ViewPoint = v
	return s
}

func (s *UpdateSubSceneRequest) Validate() error {
	return dara.Validate(s)
}
