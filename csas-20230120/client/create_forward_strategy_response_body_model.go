// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardStrategy(v *CreateForwardStrategyResponseBodyForwardStrategy) *CreateForwardStrategyResponseBody
	GetForwardStrategy() *CreateForwardStrategyResponseBodyForwardStrategy
	SetRequestId(v string) *CreateForwardStrategyResponseBody
	GetRequestId() *string
}

type CreateForwardStrategyResponseBody struct {
	// The traffic forwarding rule.
	ForwardStrategy *CreateForwardStrategyResponseBodyForwardStrategy `json:"ForwardStrategy,omitempty" xml:"ForwardStrategy,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 60D4601C-B693-51A8-BB30-0944CE500B75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateForwardStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateForwardStrategyResponseBody) GetForwardStrategy() *CreateForwardStrategyResponseBodyForwardStrategy {
	return s.ForwardStrategy
}

func (s *CreateForwardStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateForwardStrategyResponseBody) SetForwardStrategy(v *CreateForwardStrategyResponseBodyForwardStrategy) *CreateForwardStrategyResponseBody {
	s.ForwardStrategy = v
	return s
}

func (s *CreateForwardStrategyResponseBody) SetRequestId(v string) *CreateForwardStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateForwardStrategyResponseBody) Validate() error {
	if s.ForwardStrategy != nil {
		if err := s.ForwardStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateForwardStrategyResponseBodyForwardStrategy struct {
	// The description of the traffic forwarding rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The target instance ID.
	//
	// example:
	//
	// connector-af9b4ee6fd15d82d
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The destination type. Valid values:
	//
	// - **Connector**: connector.
	//
	// example:
	//
	// Connector
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The ID of the traffic forwarding rule.
	//
	// example:
	//
	// fs-345d6ab82b5a43a3
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// xftp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the internal-facing access application. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateForwardStrategyResponseBodyForwardStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardStrategyResponseBodyForwardStrategy) GoString() string {
	return s.String()
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) GetDescription() *string {
	return s.Description
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) GetDestinationId() *string {
	return s.DestinationId
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) GetDestinationType() *string {
	return s.DestinationType
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) GetForwardId() *string {
	return s.ForwardId
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) GetName() *string {
	return s.Name
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) GetPriority() *int64 {
	return s.Priority
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) GetStatus() *string {
	return s.Status
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) SetDescription(v string) *CreateForwardStrategyResponseBodyForwardStrategy {
	s.Description = &v
	return s
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) SetDestinationId(v string) *CreateForwardStrategyResponseBodyForwardStrategy {
	s.DestinationId = &v
	return s
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) SetDestinationType(v string) *CreateForwardStrategyResponseBodyForwardStrategy {
	s.DestinationType = &v
	return s
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) SetForwardId(v string) *CreateForwardStrategyResponseBodyForwardStrategy {
	s.ForwardId = &v
	return s
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) SetName(v string) *CreateForwardStrategyResponseBodyForwardStrategy {
	s.Name = &v
	return s
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) SetPriority(v int64) *CreateForwardStrategyResponseBodyForwardStrategy {
	s.Priority = &v
	return s
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) SetStatus(v string) *CreateForwardStrategyResponseBodyForwardStrategy {
	s.Status = &v
	return s
}

func (s *CreateForwardStrategyResponseBodyForwardStrategy) Validate() error {
	return dara.Validate(s)
}
