// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubSceneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateSubSceneShrinkRequest
	GetId() *string
	SetName(v string) *UpdateSubSceneShrinkRequest
	GetName() *string
	SetViewPointShrink(v string) *UpdateSubSceneShrinkRequest
	GetViewPointShrink() *string
}

type UpdateSubSceneShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 测试
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ViewPointShrink *string `json:"ViewPoint,omitempty" xml:"ViewPoint,omitempty"`
}

func (s UpdateSubSceneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateSubSceneShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSubSceneShrinkRequest) GetViewPointShrink() *string {
	return s.ViewPointShrink
}

func (s *UpdateSubSceneShrinkRequest) SetId(v string) *UpdateSubSceneShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateSubSceneShrinkRequest) SetName(v string) *UpdateSubSceneShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSubSceneShrinkRequest) SetViewPointShrink(v string) *UpdateSubSceneShrinkRequest {
	s.ViewPointShrink = &v
	return s
}

func (s *UpdateSubSceneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
