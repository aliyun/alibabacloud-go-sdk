// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOutboundCallRestrictionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateOutboundCallRestrictionShrinkRequest
	GetInstanceId() *string
	SetOutboundCallRestrictionShrink(v string) *CreateOutboundCallRestrictionShrinkRequest
	GetOutboundCallRestrictionShrink() *string
	SetPolicy(v int32) *CreateOutboundCallRestrictionShrinkRequest
	GetPolicy() *int32
}

type CreateOutboundCallRestrictionShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The outbound call restriction.
	OutboundCallRestrictionShrink *string `json:"OutboundCallRestriction,omitempty" xml:"OutboundCallRestriction,omitempty"`
	// The policy. Valid values:
	//
	// 0: blacklist.
	//
	// 1: whitelist.
	//
	// example:
	//
	// 0
	Policy *int32 `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CreateOutboundCallRestrictionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOutboundCallRestrictionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOutboundCallRestrictionShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateOutboundCallRestrictionShrinkRequest) GetOutboundCallRestrictionShrink() *string {
	return s.OutboundCallRestrictionShrink
}

func (s *CreateOutboundCallRestrictionShrinkRequest) GetPolicy() *int32 {
	return s.Policy
}

func (s *CreateOutboundCallRestrictionShrinkRequest) SetInstanceId(v string) *CreateOutboundCallRestrictionShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOutboundCallRestrictionShrinkRequest) SetOutboundCallRestrictionShrink(v string) *CreateOutboundCallRestrictionShrinkRequest {
	s.OutboundCallRestrictionShrink = &v
	return s
}

func (s *CreateOutboundCallRestrictionShrinkRequest) SetPolicy(v int32) *CreateOutboundCallRestrictionShrinkRequest {
	s.Policy = &v
	return s
}

func (s *CreateOutboundCallRestrictionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
