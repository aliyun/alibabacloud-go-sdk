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
	// List of attachment point IDs.
	//
	// This parameter is required.
	AttachResourceIds []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	// Types of attachment points supported by the policy.
	//
	// - HttpApi: HttpApi.
	//
	// - Operation: Operation of HttpApi.
	//
	// - GatewayRoute: Gateway route.
	//
	// - GatewayService: Gateway service.
	//
	// - GatewayServicePort: Gateway service port.
	//
	// - Domain: Gateway domain.
	//
	// - Gateway: Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The class name types supported by the policy. Different policies support different attachment points, to be used in conjunction with `attachResourceType`.
	//
	// - RateLimit: Traffic control, supports HttpApi, Operation, GatewayRoute.
	//
	// - ConcurrencyLimit: Concurrency control, supports HttpApi, Operation, GatewayRoute.
	//
	// - CircuitBreaker: Circuit breaking and degradation, supports HttpApi, Operation, GatewayRoute.
	//
	// - HttpRewrite: HTTP rewrite, supports HttpApi, Operation, GatewayRoute.
	//
	// - HeaderModify: Header modification, supports HttpApi, Operation, GatewayRoute.
	//
	// - Cors: Cross-origin, supports HttpApi, Operation, GatewayRoute.
	//
	// - FlowCopy: Traffic replication, supports HttpApi, Operation, GatewayRoute.
	//
	// - Timeout: Timeout, supports HttpApi, Operation, GatewayRoute.
	//
	// - Retry: Retry, supports HttpApi, Operation, GatewayRoute.
	//
	// - IpAccessControl: IP access control, supports HttpApi, Operation, GatewayRoute, Domain, Gateway.
	//
	// - DirectResponse: Mock, supports Operation, GatewayRoute.
	//
	// - Redirect: Redirection, supports GatewayRoute.
	//
	// - Fallback: Fallback, supports Operation, GatewayRoute.
	//
	// - ServiceTls: Service TLS authentication, supports GatewayService.
	//
	// - ServiceLb: Service load balancing, supports GatewayService.
	//
	// - ServicePortTls: Service port TLS authentication, supports GatewayServicePort.
	//
	// - Waf: WAF protection, supports GatewayRoute, Gateway.
	//
	// - JWTAuth: JWT global authentication, supports Gateway.
	//
	// - OIDCAuth: OIDC global authentication, supports Gateway.
	//
	// - ExternalZAuth: Custom authorization, supports Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// IpAccessControl
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// Configuration information.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"enable\\":false}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// Policy description.
	//
	// example:
	//
	// 这是策略描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Environment ID.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Gateway ID.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qasrdc0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Policy name.
	//
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
