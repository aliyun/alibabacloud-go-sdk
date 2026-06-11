// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionOp interface {
	dara.Model
	String() string
	GoString() string
	SetOp(v string) *SubscriptionOp
	GetOp() *string
	SetPayload(v *SubscriptionForModify) *SubscriptionOp
	GetPayload() *SubscriptionForModify
	SetUuid(v string) *SubscriptionOp
	GetUuid() *string
}

type SubscriptionOp struct {
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// create/update 必填
	Payload *SubscriptionForModify `json:"payload,omitempty" xml:"payload,omitempty"`
	// update/remove 必填
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s SubscriptionOp) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionOp) GoString() string {
	return s.String()
}

func (s *SubscriptionOp) GetOp() *string {
	return s.Op
}

func (s *SubscriptionOp) GetPayload() *SubscriptionForModify {
	return s.Payload
}

func (s *SubscriptionOp) GetUuid() *string {
	return s.Uuid
}

func (s *SubscriptionOp) SetOp(v string) *SubscriptionOp {
	s.Op = &v
	return s
}

func (s *SubscriptionOp) SetPayload(v *SubscriptionForModify) *SubscriptionOp {
	s.Payload = v
	return s
}

func (s *SubscriptionOp) SetUuid(v string) *SubscriptionOp {
	s.Uuid = &v
	return s
}

func (s *SubscriptionOp) Validate() error {
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}
