// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardStrategy(v *ModifyForwardStrategyResponseBodyForwardStrategy) *ModifyForwardStrategyResponseBody
	GetForwardStrategy() *ModifyForwardStrategyResponseBodyForwardStrategy
	SetRequestId(v string) *ModifyForwardStrategyResponseBody
	GetRequestId() *string
}

type ModifyForwardStrategyResponseBody struct {
	// The forwarding rule.
	ForwardStrategy *ModifyForwardStrategyResponseBodyForwardStrategy `json:"ForwardStrategy,omitempty" xml:"ForwardStrategy,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2EBEEB93-E7AF-5667-B492-FA95C70821A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyForwardStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyResponseBody) GetForwardStrategy() *ModifyForwardStrategyResponseBodyForwardStrategy {
	return s.ForwardStrategy
}

func (s *ModifyForwardStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyForwardStrategyResponseBody) SetForwardStrategy(v *ModifyForwardStrategyResponseBodyForwardStrategy) *ModifyForwardStrategyResponseBody {
	s.ForwardStrategy = v
	return s
}

func (s *ModifyForwardStrategyResponseBody) SetRequestId(v string) *ModifyForwardStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyForwardStrategyResponseBody) Validate() error {
	if s.ForwardStrategy != nil {
		if err := s.ForwardStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyForwardStrategyResponseBodyForwardStrategy struct {
	// The policy description.
	//
	// example:
	//
	// material_versions_rec
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The target instance ID.
	//
	// example:
	//
	// connector-f0b9195a6f2597fa
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The destination type. Valid values:
	//
	// - **Connector**: connector.
	//
	// example:
	//
	// Connector
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The forwarding rule ID.
	//
	// example:
	//
	// fs-037cee3b6ebaa919
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// SaseSSO
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority.
	//
	// example:
	//
	// 100
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The policy status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyForwardStrategyResponseBodyForwardStrategy) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyResponseBodyForwardStrategy) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) GetDescription() *string {
	return s.Description
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) GetDestinationId() *string {
	return s.DestinationId
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) GetForwardId() *string {
	return s.ForwardId
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) GetName() *string {
	return s.Name
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) GetPriority() *int64 {
	return s.Priority
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) GetStatus() *string {
	return s.Status
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) SetDescription(v string) *ModifyForwardStrategyResponseBodyForwardStrategy {
	s.Description = &v
	return s
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) SetDestinationId(v string) *ModifyForwardStrategyResponseBodyForwardStrategy {
	s.DestinationId = &v
	return s
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) SetDestinationType(v string) *ModifyForwardStrategyResponseBodyForwardStrategy {
	s.DestinationType = &v
	return s
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) SetForwardId(v string) *ModifyForwardStrategyResponseBodyForwardStrategy {
	s.ForwardId = &v
	return s
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) SetName(v string) *ModifyForwardStrategyResponseBodyForwardStrategy {
	s.Name = &v
	return s
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) SetPriority(v int64) *ModifyForwardStrategyResponseBodyForwardStrategy {
	s.Priority = &v
	return s
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) SetStatus(v string) *ModifyForwardStrategyResponseBodyForwardStrategy {
	s.Status = &v
	return s
}

func (s *ModifyForwardStrategyResponseBodyForwardStrategy) Validate() error {
	return dara.Validate(s)
}
