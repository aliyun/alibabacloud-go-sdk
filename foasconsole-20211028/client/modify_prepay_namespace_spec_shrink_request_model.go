// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayNamespaceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyPrepayNamespaceSpecShrinkRequest
	GetInstanceId() *string
	SetNamespace(v string) *ModifyPrepayNamespaceSpecShrinkRequest
	GetNamespace() *string
	SetRegion(v string) *ModifyPrepayNamespaceSpecShrinkRequest
	GetRegion() *string
	SetResourceSpecShrink(v string) *ModifyPrepayNamespaceSpecShrinkRequest
	GetResourceSpecShrink() *string
}

type ModifyPrepayNamespaceSpecShrinkRequest struct {
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
	// di-593440219799842
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s ModifyPrepayNamespaceSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) SetInstanceId(v string) *ModifyPrepayNamespaceSpecShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) SetNamespace(v string) *ModifyPrepayNamespaceSpecShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) SetRegion(v string) *ModifyPrepayNamespaceSpecShrinkRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) SetResourceSpecShrink(v string) *ModifyPrepayNamespaceSpecShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
