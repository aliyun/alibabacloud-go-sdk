// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTargetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *DeleteTargetsShrinkRequest
	GetEventBusName() *string
	SetRuleName(v string) *DeleteTargetsShrinkRequest
	GetRuleName() *string
	SetTargetIdsShrink(v string) *DeleteTargetsShrinkRequest
	GetTargetIdsShrink() *string
}

type DeleteTargetsShrinkRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ramrolechange-mns
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The IDs of the event targets that you want to delete.
	TargetIdsShrink *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
}

func (s DeleteTargetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTargetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteTargetsShrinkRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *DeleteTargetsShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteTargetsShrinkRequest) GetTargetIdsShrink() *string {
	return s.TargetIdsShrink
}

func (s *DeleteTargetsShrinkRequest) SetEventBusName(v string) *DeleteTargetsShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *DeleteTargetsShrinkRequest) SetRuleName(v string) *DeleteTargetsShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteTargetsShrinkRequest) SetTargetIdsShrink(v string) *DeleteTargetsShrinkRequest {
	s.TargetIdsShrink = &v
	return s
}

func (s *DeleteTargetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
