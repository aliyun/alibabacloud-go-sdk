// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *CreateInstanceRequest
	GetConcurrency() *int32
	SetDescription(v string) *CreateInstanceRequest
	GetDescription() *string
	SetName(v string) *CreateInstanceRequest
	GetName() *string
	SetServiceMode(v string) *CreateInstanceRequest
	GetServiceMode() *string
}

type CreateInstanceRequest struct {
	// The number of concurrent calls.
	//
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The instance description.
	//
	// example:
	//
	// Intelligent outbound call instance for telemarketing scenarios
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Intelligent outbound call instance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service mode.
	//
	// example:
	//
	// STANDARD
	ServiceMode *string `json:"ServiceMode,omitempty" xml:"ServiceMode,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *CreateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateInstanceRequest) GetServiceMode() *string {
	return s.ServiceMode
}

func (s *CreateInstanceRequest) SetConcurrency(v int32) *CreateInstanceRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetServiceMode(v string) *CreateInstanceRequest {
	s.ServiceMode = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
