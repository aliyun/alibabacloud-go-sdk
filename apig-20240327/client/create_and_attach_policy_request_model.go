// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndAttachPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceIds(v []*string) *CreateAndAttachPolicyRequest
	GetAttachResourceIds() []*string
	SetAttachResourceType(v string) *CreateAndAttachPolicyRequest
	GetAttachResourceType() *string
	SetClassName(v string) *CreateAndAttachPolicyRequest
	GetClassName() *string
	SetConfig(v string) *CreateAndAttachPolicyRequest
	GetConfig() *string
	SetDescription(v string) *CreateAndAttachPolicyRequest
	GetDescription() *string
	SetEnvironmentId(v string) *CreateAndAttachPolicyRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *CreateAndAttachPolicyRequest
	GetGatewayId() *string
	SetName(v string) *CreateAndAttachPolicyRequest
	GetName() *string
}

type CreateAndAttachPolicyRequest struct {
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
	// IpAccessControl
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\\"enable\\":false}
	Config      *string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// gw-cq7l5s5lhtgi6qasrdc0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAndAttachPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndAttachPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAndAttachPolicyRequest) GetAttachResourceIds() []*string {
	return s.AttachResourceIds
}

func (s *CreateAndAttachPolicyRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *CreateAndAttachPolicyRequest) GetClassName() *string {
	return s.ClassName
}

func (s *CreateAndAttachPolicyRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateAndAttachPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAndAttachPolicyRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateAndAttachPolicyRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateAndAttachPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateAndAttachPolicyRequest) SetAttachResourceIds(v []*string) *CreateAndAttachPolicyRequest {
	s.AttachResourceIds = v
	return s
}

func (s *CreateAndAttachPolicyRequest) SetAttachResourceType(v string) *CreateAndAttachPolicyRequest {
	s.AttachResourceType = &v
	return s
}

func (s *CreateAndAttachPolicyRequest) SetClassName(v string) *CreateAndAttachPolicyRequest {
	s.ClassName = &v
	return s
}

func (s *CreateAndAttachPolicyRequest) SetConfig(v string) *CreateAndAttachPolicyRequest {
	s.Config = &v
	return s
}

func (s *CreateAndAttachPolicyRequest) SetDescription(v string) *CreateAndAttachPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateAndAttachPolicyRequest) SetEnvironmentId(v string) *CreateAndAttachPolicyRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateAndAttachPolicyRequest) SetGatewayId(v string) *CreateAndAttachPolicyRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateAndAttachPolicyRequest) SetName(v string) *CreateAndAttachPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateAndAttachPolicyRequest) Validate() error {
	return dara.Validate(s)
}
