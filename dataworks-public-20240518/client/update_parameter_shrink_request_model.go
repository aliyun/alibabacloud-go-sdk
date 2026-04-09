// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateParameterShrinkRequest
	GetDescription() *string
	SetId(v int64) *UpdateParameterShrinkRequest
	GetId() *int64
	SetOwner(v string) *UpdateParameterShrinkRequest
	GetOwner() *string
	SetPropertiesShrink(v string) *UpdateParameterShrinkRequest
	GetPropertiesShrink() *string
}

type UpdateParameterShrinkRequest struct {
	// example:
	//
	// 这是一个测试参数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123456789
	Owner            *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s UpdateParameterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateParameterShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateParameterShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateParameterShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateParameterShrinkRequest) GetPropertiesShrink() *string {
	return s.PropertiesShrink
}

func (s *UpdateParameterShrinkRequest) SetDescription(v string) *UpdateParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateParameterShrinkRequest) SetId(v int64) *UpdateParameterShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateParameterShrinkRequest) SetOwner(v string) *UpdateParameterShrinkRequest {
	s.Owner = &v
	return s
}

func (s *UpdateParameterShrinkRequest) SetPropertiesShrink(v string) *UpdateParameterShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *UpdateParameterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
