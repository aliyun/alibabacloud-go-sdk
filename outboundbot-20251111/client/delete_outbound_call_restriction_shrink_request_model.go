// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundCallRestrictionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteOutboundCallRestrictionShrinkRequest
	GetInstanceId() *string
	SetRestrictionIdListShrink(v string) *DeleteOutboundCallRestrictionShrinkRequest
	GetRestrictionIdListShrink() *string
}

type DeleteOutboundCallRestrictionShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of outbound restriction IDs.
	RestrictionIdListShrink *string `json:"RestrictionIdList,omitempty" xml:"RestrictionIdList,omitempty"`
}

func (s DeleteOutboundCallRestrictionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundCallRestrictionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteOutboundCallRestrictionShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteOutboundCallRestrictionShrinkRequest) GetRestrictionIdListShrink() *string {
	return s.RestrictionIdListShrink
}

func (s *DeleteOutboundCallRestrictionShrinkRequest) SetInstanceId(v string) *DeleteOutboundCallRestrictionShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOutboundCallRestrictionShrinkRequest) SetRestrictionIdListShrink(v string) *DeleteOutboundCallRestrictionShrinkRequest {
	s.RestrictionIdListShrink = &v
	return s
}

func (s *DeleteOutboundCallRestrictionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
