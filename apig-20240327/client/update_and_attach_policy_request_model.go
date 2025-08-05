// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndAttachPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceIds(v []*string) *UpdateAndAttachPolicyRequest
	GetAttachResourceIds() []*string
	SetAttachResourceType(v string) *UpdateAndAttachPolicyRequest
	GetAttachResourceType() *string
	SetConfig(v string) *UpdateAndAttachPolicyRequest
	GetConfig() *string
	SetDescription(v string) *UpdateAndAttachPolicyRequest
	GetDescription() *string
	SetEnvironmentId(v string) *UpdateAndAttachPolicyRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *UpdateAndAttachPolicyRequest
	GetGatewayId() *string
	SetName(v string) *UpdateAndAttachPolicyRequest
	GetName() *string
}

type UpdateAndAttachPolicyRequest struct {
	// This parameter is required.
	AttachResourceIds []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\\"enable\\":false}
	Config      *string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// env-cq2avtllhtgja4dk5djg
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// gw-cq2avtllhtgja4dk5djg
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateAndAttachPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndAttachPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAndAttachPolicyRequest) GetAttachResourceIds() []*string {
	return s.AttachResourceIds
}

func (s *UpdateAndAttachPolicyRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *UpdateAndAttachPolicyRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateAndAttachPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAndAttachPolicyRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateAndAttachPolicyRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateAndAttachPolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAndAttachPolicyRequest) SetAttachResourceIds(v []*string) *UpdateAndAttachPolicyRequest {
	s.AttachResourceIds = v
	return s
}

func (s *UpdateAndAttachPolicyRequest) SetAttachResourceType(v string) *UpdateAndAttachPolicyRequest {
	s.AttachResourceType = &v
	return s
}

func (s *UpdateAndAttachPolicyRequest) SetConfig(v string) *UpdateAndAttachPolicyRequest {
	s.Config = &v
	return s
}

func (s *UpdateAndAttachPolicyRequest) SetDescription(v string) *UpdateAndAttachPolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdateAndAttachPolicyRequest) SetEnvironmentId(v string) *UpdateAndAttachPolicyRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateAndAttachPolicyRequest) SetGatewayId(v string) *UpdateAndAttachPolicyRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateAndAttachPolicyRequest) SetName(v string) *UpdateAndAttachPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateAndAttachPolicyRequest) Validate() error {
	return dara.Validate(s)
}
