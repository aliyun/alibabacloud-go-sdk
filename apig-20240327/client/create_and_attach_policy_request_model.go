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
	// The list of target resource IDs to attach.
	//
	// This parameter is required.
	AttachResourceIds []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	// The type of the target resource to attach.
	//
	// This parameter is required.
	//
	// example:
	//
	// GatewayRoute
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The policy type.
	//
	// This parameter is required.
	//
	// example:
	//
	// AiFallback
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// The policy configuration content (JSON string).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"serviceConfigs":[...]}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The policy description.
	//
	// example:
	//
	// Fallback when primary route fails
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-test
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The policy name.
	//
	// This parameter is required when className is set to IpAccessControl, JWTAuth, OIDCAuth, or ExternalZAuth, and must be unique within the same gateway instance (gatewayId) under the current account. If the name conflicts with an existing policy, a PolicyExisted error is returned. When retrying after receiving this error, you must use a different name. Submitting the same name repeatedly will always fail.
	//
	// For other className values (such as RateLimit, Timeout, Retry, ServiceTls, and AiProxy), name is optional and used only for display purposes.
	//
	// example:
	//
	// my-fallback-policy
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
