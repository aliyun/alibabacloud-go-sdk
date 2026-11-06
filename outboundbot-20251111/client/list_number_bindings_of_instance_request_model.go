// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNumberBindingsOfInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNumberBindingsOfInstanceRequest
	GetInstanceId() *string
	SetTenantId(v string) *ListNumberBindingsOfInstanceRequest
	GetTenantId() *string
}

type ListNumberBindingsOfInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListNumberBindingsOfInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNumberBindingsOfInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListNumberBindingsOfInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNumberBindingsOfInstanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListNumberBindingsOfInstanceRequest) SetInstanceId(v string) *ListNumberBindingsOfInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNumberBindingsOfInstanceRequest) SetTenantId(v string) *ListNumberBindingsOfInstanceRequest {
	s.TenantId = &v
	return s
}

func (s *ListNumberBindingsOfInstanceRequest) Validate() error {
	return dara.Validate(s)
}
