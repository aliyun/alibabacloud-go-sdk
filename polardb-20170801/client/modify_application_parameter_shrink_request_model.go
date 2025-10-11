// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationParameterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationParameterShrinkRequest
	GetApplicationId() *string
	SetParameterName(v string) *ModifyApplicationParameterShrinkRequest
	GetParameterName() *string
	SetParameterValue(v string) *ModifyApplicationParameterShrinkRequest
	GetParameterValue() *string
	SetParametersShrink(v string) *ModifyApplicationParameterShrinkRequest
	GetParametersShrink() *string
}

type ModifyApplicationParameterShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// name
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// example:
	//
	// value
	ParameterValue   *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ModifyApplicationParameterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationParameterShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationParameterShrinkRequest) GetParameterName() *string {
	return s.ParameterName
}

func (s *ModifyApplicationParameterShrinkRequest) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *ModifyApplicationParameterShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *ModifyApplicationParameterShrinkRequest) SetApplicationId(v string) *ModifyApplicationParameterShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationParameterShrinkRequest) SetParameterName(v string) *ModifyApplicationParameterShrinkRequest {
	s.ParameterName = &v
	return s
}

func (s *ModifyApplicationParameterShrinkRequest) SetParameterValue(v string) *ModifyApplicationParameterShrinkRequest {
	s.ParameterValue = &v
	return s
}

func (s *ModifyApplicationParameterShrinkRequest) SetParametersShrink(v string) *ModifyApplicationParameterShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *ModifyApplicationParameterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
