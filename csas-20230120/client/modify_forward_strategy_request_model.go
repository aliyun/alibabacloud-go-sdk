// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyForwardStrategyRequest
	GetDescription() *string
	SetDestinationId(v string) *ModifyForwardStrategyRequest
	GetDestinationId() *string
	SetDestinationType(v string) *ModifyForwardStrategyRequest
	GetDestinationType() *string
	SetForwardId(v string) *ModifyForwardStrategyRequest
	GetForwardId() *string
	SetName(v string) *ModifyForwardStrategyRequest
	GetName() *string
	SetPriority(v int32) *ModifyForwardStrategyRequest
	GetPriority() *int32
	SetStatus(v string) *ModifyForwardStrategyRequest
	GetStatus() *string
}

type ModifyForwardStrategyRequest struct {
	// The policy description.
	//
	// example:
	//
	// This is an internal access policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The target instance ID.
	//
	// example:
	//
	// connector-e3152978fb32443b
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
	// This parameter is required.
	//
	// example:
	//
	// fs-051199361a1fbefc
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// CollegeStudentsOnline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy priority. A value of 1 indicates the highest priority, and a value of 100 indicates the lowest priority.
	//
	// example:
	//
	// 100
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
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

func (s ModifyForwardStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyForwardStrategyRequest) GetDestinationId() *string {
	return s.DestinationId
}

func (s *ModifyForwardStrategyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ModifyForwardStrategyRequest) GetForwardId() *string {
	return s.ForwardId
}

func (s *ModifyForwardStrategyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyForwardStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyForwardStrategyRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyForwardStrategyRequest) SetDescription(v string) *ModifyForwardStrategyRequest {
	s.Description = &v
	return s
}

func (s *ModifyForwardStrategyRequest) SetDestinationId(v string) *ModifyForwardStrategyRequest {
	s.DestinationId = &v
	return s
}

func (s *ModifyForwardStrategyRequest) SetDestinationType(v string) *ModifyForwardStrategyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyForwardStrategyRequest) SetForwardId(v string) *ModifyForwardStrategyRequest {
	s.ForwardId = &v
	return s
}

func (s *ModifyForwardStrategyRequest) SetName(v string) *ModifyForwardStrategyRequest {
	s.Name = &v
	return s
}

func (s *ModifyForwardStrategyRequest) SetPriority(v int32) *ModifyForwardStrategyRequest {
	s.Priority = &v
	return s
}

func (s *ModifyForwardStrategyRequest) SetStatus(v string) *ModifyForwardStrategyRequest {
	s.Status = &v
	return s
}

func (s *ModifyForwardStrategyRequest) Validate() error {
	return dara.Validate(s)
}
