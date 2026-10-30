// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventHouseRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v int32) *UpdateEventHouseRuntimeRequest
	GetCu() *int32
	SetName(v string) *UpdateEventHouseRuntimeRequest
	GetName() *string
}

type UpdateEventHouseRuntimeRequest struct {
	// The number of compute units (CUs) for the EventHouse runtime. The value must be greater than 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The name of the EventHouse runtime. If you do not specify this parameter, the default runtime is used. Typically, you do not need to specify this parameter in the initial phase.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateEventHouseRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventHouseRuntimeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventHouseRuntimeRequest) GetCu() *int32 {
	return s.Cu
}

func (s *UpdateEventHouseRuntimeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEventHouseRuntimeRequest) SetCu(v int32) *UpdateEventHouseRuntimeRequest {
	s.Cu = &v
	return s
}

func (s *UpdateEventHouseRuntimeRequest) SetName(v string) *UpdateEventHouseRuntimeRequest {
	s.Name = &v
	return s
}

func (s *UpdateEventHouseRuntimeRequest) Validate() error {
	return dara.Validate(s)
}
