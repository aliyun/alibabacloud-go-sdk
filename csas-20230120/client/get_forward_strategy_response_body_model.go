// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetForwardStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardStrategy(v *GetForwardStrategyResponseBodyForwardStrategy) *GetForwardStrategyResponseBody
	GetForwardStrategy() *GetForwardStrategyResponseBodyForwardStrategy
	SetRequestId(v string) *GetForwardStrategyResponseBody
	GetRequestId() *string
}

type GetForwardStrategyResponseBody struct {
	// The forwarding rule.
	ForwardStrategy *GetForwardStrategyResponseBodyForwardStrategy `json:"ForwardStrategy,omitempty" xml:"ForwardStrategy,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetForwardStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetForwardStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetForwardStrategyResponseBody) GetForwardStrategy() *GetForwardStrategyResponseBodyForwardStrategy {
	return s.ForwardStrategy
}

func (s *GetForwardStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetForwardStrategyResponseBody) SetForwardStrategy(v *GetForwardStrategyResponseBodyForwardStrategy) *GetForwardStrategyResponseBody {
	s.ForwardStrategy = v
	return s
}

func (s *GetForwardStrategyResponseBody) SetRequestId(v string) *GetForwardStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetForwardStrategyResponseBody) Validate() error {
	if s.ForwardStrategy != nil {
		if err := s.ForwardStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetForwardStrategyResponseBodyForwardStrategy struct {
	// The policy description.
	//
	// example:
	//
	// solemn_index
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The target instance ID.
	//
	// example:
	//
	// connector-bb95f515b6818623
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
	// fs-b87a2f8e863bf02c
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// dynamic_route_name_eb55d3a3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy priority. The value 1 indicates the highest priority, and the value 100 indicates the lowest priority.
	//
	// example:
	//
	// 1
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

func (s GetForwardStrategyResponseBodyForwardStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetForwardStrategyResponseBodyForwardStrategy) GoString() string {
	return s.String()
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) GetDescription() *string {
	return s.Description
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) GetDestinationId() *string {
	return s.DestinationId
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) GetDestinationType() *string {
	return s.DestinationType
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) GetForwardId() *string {
	return s.ForwardId
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) GetName() *string {
	return s.Name
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) GetPriority() *int64 {
	return s.Priority
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) GetStatus() *string {
	return s.Status
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) SetDescription(v string) *GetForwardStrategyResponseBodyForwardStrategy {
	s.Description = &v
	return s
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) SetDestinationId(v string) *GetForwardStrategyResponseBodyForwardStrategy {
	s.DestinationId = &v
	return s
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) SetDestinationType(v string) *GetForwardStrategyResponseBodyForwardStrategy {
	s.DestinationType = &v
	return s
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) SetForwardId(v string) *GetForwardStrategyResponseBodyForwardStrategy {
	s.ForwardId = &v
	return s
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) SetName(v string) *GetForwardStrategyResponseBodyForwardStrategy {
	s.Name = &v
	return s
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) SetPriority(v int64) *GetForwardStrategyResponseBodyForwardStrategy {
	s.Priority = &v
	return s
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) SetStatus(v string) *GetForwardStrategyResponseBodyForwardStrategy {
	s.Status = &v
	return s
}

func (s *GetForwardStrategyResponseBodyForwardStrategy) Validate() error {
	return dara.Validate(s)
}
