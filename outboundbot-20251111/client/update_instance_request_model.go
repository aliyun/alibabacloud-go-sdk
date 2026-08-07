// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *UpdateInstanceRequest
	GetConcurrency() *int32
	SetDescription(v string) *UpdateInstanceRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetName(v string) *UpdateInstanceRequest
	GetName() *string
}

type UpdateInstanceRequest struct {
	// The number of concurrent calls.
	//
	// example:
	//
	// 20
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The instance description.
	//
	// example:
	//
	// Intelligent outbound call instance for telemarketing scenarios
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Intelligent outbound call instance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *UpdateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateInstanceRequest) SetConcurrency(v int32) *UpdateInstanceRequest {
	s.Concurrency = &v
	return s
}

func (s *UpdateInstanceRequest) SetDescription(v string) *UpdateInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetName(v string) *UpdateInstanceRequest {
	s.Name = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
