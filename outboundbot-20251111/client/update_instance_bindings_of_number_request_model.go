// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceBindingsOfNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *UpdateInstanceBindingsOfNumberRequest
	GetInstanceIds() []*string
	SetNumber(v string) *UpdateInstanceBindingsOfNumberRequest
	GetNumber() *string
	SetTenantId(v string) *UpdateInstanceBindingsOfNumberRequest
	GetTenantId() *string
}

type UpdateInstanceBindingsOfNumberRequest struct {
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The bound number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 057100001234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s UpdateInstanceBindingsOfNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceBindingsOfNumberRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceBindingsOfNumberRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *UpdateInstanceBindingsOfNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *UpdateInstanceBindingsOfNumberRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateInstanceBindingsOfNumberRequest) SetInstanceIds(v []*string) *UpdateInstanceBindingsOfNumberRequest {
	s.InstanceIds = v
	return s
}

func (s *UpdateInstanceBindingsOfNumberRequest) SetNumber(v string) *UpdateInstanceBindingsOfNumberRequest {
	s.Number = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberRequest) SetTenantId(v string) *UpdateInstanceBindingsOfNumberRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberRequest) Validate() error {
	return dara.Validate(s)
}
