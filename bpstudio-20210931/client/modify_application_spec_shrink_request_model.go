// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationSpecShrinkRequest
	GetApplicationId() *string
	SetInstanceSpecShrink(v string) *ModifyApplicationSpecShrinkRequest
	GetInstanceSpecShrink() *string
}

type ModifyApplicationSpecShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId      *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	InstanceSpecShrink *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
}

func (s ModifyApplicationSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationSpecShrinkRequest) GetInstanceSpecShrink() *string {
	return s.InstanceSpecShrink
}

func (s *ModifyApplicationSpecShrinkRequest) SetApplicationId(v string) *ModifyApplicationSpecShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationSpecShrinkRequest) SetInstanceSpecShrink(v string) *ModifyApplicationSpecShrinkRequest {
	s.InstanceSpecShrink = &v
	return s
}

func (s *ModifyApplicationSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
