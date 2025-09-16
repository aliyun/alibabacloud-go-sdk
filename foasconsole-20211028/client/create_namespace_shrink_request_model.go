// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *CreateNamespaceShrinkRequest
	GetHa() *bool
	SetInstanceId(v string) *CreateNamespaceShrinkRequest
	GetInstanceId() *string
	SetNamespace(v string) *CreateNamespaceShrinkRequest
	GetNamespace() *string
	SetRegion(v string) *CreateNamespaceShrinkRequest
	GetRegion() *string
	SetResourceSpecShrink(v string) *CreateNamespaceShrinkRequest
	GetResourceSpecShrink() *string
}

type CreateNamespaceShrinkRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// di-593440390152545
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s CreateNamespaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceShrinkRequest) GetHa() *bool {
	return s.Ha
}

func (s *CreateNamespaceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNamespaceShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateNamespaceShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateNamespaceShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *CreateNamespaceShrinkRequest) SetHa(v bool) *CreateNamespaceShrinkRequest {
	s.Ha = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetInstanceId(v string) *CreateNamespaceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetNamespace(v string) *CreateNamespaceShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetRegion(v string) *CreateNamespaceShrinkRequest {
	s.Region = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetResourceSpecShrink(v string) *CreateNamespaceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
