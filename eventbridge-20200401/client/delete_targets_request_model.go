// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *DeleteTargetsRequest
	GetEventBusName() *string
	SetRuleName(v string) *DeleteTargetsRequest
	GetRuleName() *string
	SetTargetIds(v []*string) *DeleteTargetsRequest
	GetTargetIds() []*string
}

type DeleteTargetsRequest struct {
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
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
}

func (s DeleteTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTargetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTargetsRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *DeleteTargetsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteTargetsRequest) GetTargetIds() []*string {
	return s.TargetIds
}

func (s *DeleteTargetsRequest) SetEventBusName(v string) *DeleteTargetsRequest {
	s.EventBusName = &v
	return s
}

func (s *DeleteTargetsRequest) SetRuleName(v string) *DeleteTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteTargetsRequest) SetTargetIds(v []*string) *DeleteTargetsRequest {
	s.TargetIds = v
	return s
}

func (s *DeleteTargetsRequest) Validate() error {
	return dara.Validate(s)
}
