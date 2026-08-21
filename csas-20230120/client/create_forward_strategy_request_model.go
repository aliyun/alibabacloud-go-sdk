// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateForwardStrategyRequest
	GetDescription() *string
	SetDestinationId(v string) *CreateForwardStrategyRequest
	GetDestinationId() *string
	SetDestinationType(v string) *CreateForwardStrategyRequest
	GetDestinationType() *string
	SetName(v string) *CreateForwardStrategyRequest
	GetName() *string
	SetPriority(v int32) *CreateForwardStrategyRequest
	GetPriority() *int32
	SetStatus(v string) *CreateForwardStrategyRequest
	GetStatus() *string
}

type CreateForwardStrategyRequest struct {
	// The description. The description must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), hyphens (-), and spaces. The description can also contain Chinese characters.
	//
	// example:
	//
	// This is an internal access policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-4178bc59bec56df1
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The destination type. Valid values:
	//
	// - **Connector**: connector.
	//
	// This parameter is required.
	//
	// example:
	//
	// Connector
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The name. The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name can also contain Chinese characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// jogg-K8sapi
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy priority. A value of 1 indicates the highest priority. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The policy status. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateForwardStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardStrategyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateForwardStrategyRequest) GetDestinationId() *string {
	return s.DestinationId
}

func (s *CreateForwardStrategyRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *CreateForwardStrategyRequest) GetName() *string {
	return s.Name
}

func (s *CreateForwardStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateForwardStrategyRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateForwardStrategyRequest) SetDescription(v string) *CreateForwardStrategyRequest {
	s.Description = &v
	return s
}

func (s *CreateForwardStrategyRequest) SetDestinationId(v string) *CreateForwardStrategyRequest {
	s.DestinationId = &v
	return s
}

func (s *CreateForwardStrategyRequest) SetDestinationType(v string) *CreateForwardStrategyRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateForwardStrategyRequest) SetName(v string) *CreateForwardStrategyRequest {
	s.Name = &v
	return s
}

func (s *CreateForwardStrategyRequest) SetPriority(v int32) *CreateForwardStrategyRequest {
	s.Priority = &v
	return s
}

func (s *CreateForwardStrategyRequest) SetStatus(v string) *CreateForwardStrategyRequest {
	s.Status = &v
	return s
}

func (s *CreateForwardStrategyRequest) Validate() error {
	return dara.Validate(s)
}
