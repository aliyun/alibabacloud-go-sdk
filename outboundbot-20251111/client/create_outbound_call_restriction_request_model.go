// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOutboundCallRestrictionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateOutboundCallRestrictionRequest
	GetInstanceId() *string
	SetOutboundCallRestriction(v []*CreateOutboundCallRestrictionRequestOutboundCallRestriction) *CreateOutboundCallRestrictionRequest
	GetOutboundCallRestriction() []*CreateOutboundCallRestrictionRequestOutboundCallRestriction
	SetPolicy(v int32) *CreateOutboundCallRestrictionRequest
	GetPolicy() *int32
}

type CreateOutboundCallRestrictionRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The outbound call restriction.
	OutboundCallRestriction []*CreateOutboundCallRestrictionRequestOutboundCallRestriction `json:"OutboundCallRestriction,omitempty" xml:"OutboundCallRestriction,omitempty" type:"Repeated"`
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

func (s CreateOutboundCallRestrictionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOutboundCallRestrictionRequest) GoString() string {
	return s.String()
}

func (s *CreateOutboundCallRestrictionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateOutboundCallRestrictionRequest) GetOutboundCallRestriction() []*CreateOutboundCallRestrictionRequestOutboundCallRestriction {
	return s.OutboundCallRestriction
}

func (s *CreateOutboundCallRestrictionRequest) GetPolicy() *int32 {
	return s.Policy
}

func (s *CreateOutboundCallRestrictionRequest) SetInstanceId(v string) *CreateOutboundCallRestrictionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOutboundCallRestrictionRequest) SetOutboundCallRestriction(v []*CreateOutboundCallRestrictionRequestOutboundCallRestriction) *CreateOutboundCallRestrictionRequest {
	s.OutboundCallRestriction = v
	return s
}

func (s *CreateOutboundCallRestrictionRequest) SetPolicy(v int32) *CreateOutboundCallRestrictionRequest {
	s.Policy = &v
	return s
}

func (s *CreateOutboundCallRestrictionRequest) Validate() error {
	if s.OutboundCallRestriction != nil {
		for _, item := range s.OutboundCallRestriction {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOutboundCallRestrictionRequestOutboundCallRestriction struct {
	// The phone number.
	//
	// example:
	//
	// 02032734241
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Do-not-disturb user
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateOutboundCallRestrictionRequestOutboundCallRestriction) String() string {
	return dara.Prettify(s)
}

func (s CreateOutboundCallRestrictionRequestOutboundCallRestriction) GoString() string {
	return s.String()
}

func (s *CreateOutboundCallRestrictionRequestOutboundCallRestriction) GetNumber() *string {
	return s.Number
}

func (s *CreateOutboundCallRestrictionRequestOutboundCallRestriction) GetRemark() *string {
	return s.Remark
}

func (s *CreateOutboundCallRestrictionRequestOutboundCallRestriction) SetNumber(v string) *CreateOutboundCallRestrictionRequestOutboundCallRestriction {
	s.Number = &v
	return s
}

func (s *CreateOutboundCallRestrictionRequestOutboundCallRestriction) SetRemark(v string) *CreateOutboundCallRestrictionRequestOutboundCallRestriction {
	s.Remark = &v
	return s
}

func (s *CreateOutboundCallRestrictionRequestOutboundCallRestriction) Validate() error {
	return dara.Validate(s)
}
