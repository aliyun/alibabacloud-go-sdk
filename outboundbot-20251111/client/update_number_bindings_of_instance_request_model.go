// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNumberBindingsOfInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateNumberBindingsOfInstanceRequest
	GetInstanceId() *string
	SetNumbers(v []*string) *UpdateNumberBindingsOfInstanceRequest
	GetNumbers() []*string
	SetTenantId(v string) *UpdateNumberBindingsOfInstanceRequest
	GetTenantId() *string
}

type UpdateNumberBindingsOfInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4eee9bf8-1319-468f-ac82-83c50ae389f8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of phone numbers.
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s UpdateNumberBindingsOfInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNumberBindingsOfInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNumberBindingsOfInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNumberBindingsOfInstanceRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *UpdateNumberBindingsOfInstanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateNumberBindingsOfInstanceRequest) SetInstanceId(v string) *UpdateNumberBindingsOfInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceRequest) SetNumbers(v []*string) *UpdateNumberBindingsOfInstanceRequest {
	s.Numbers = v
	return s
}

func (s *UpdateNumberBindingsOfInstanceRequest) SetTenantId(v string) *UpdateNumberBindingsOfInstanceRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceRequest) Validate() error {
	return dara.Validate(s)
}
