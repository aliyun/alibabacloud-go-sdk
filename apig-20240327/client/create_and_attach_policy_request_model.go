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
	// The IDs of the resources to be associated with the policy.
	//
	// This parameter is required.
	AttachResourceIds []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	// The supported resource type. Valid values:
	//
	// 	- HttpApi: an HTTP API
	//
	// 	- Operation: an operation in an HTTP API
	//
	// 	- GatewayRoute: a route
	//
	// 	- GatewayService: a service
	//
	// 	- GatewayServicePort: a service port
	//
	// 	- Domain: a domain name
	//
	// 	- Gateway: an instance
	//
	// This parameter is required.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The class name supported by the policy. Different policies support different resources. This parameter is used in combination with AttachResourceType.
	//
	// 	- RateLimit: throttles traffic. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- ConcurrencyLimit: controls concurrency. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- CircuitBreaker: breaks circuits and downgrades traffic. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- HttpRewrite: rewrites HTTP traffic. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- HeaderModify: modifies headers. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- Cors: supports CORS. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- FlowCopy: replicates traffic. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- Timeout: times out requests. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- Retry: retries requests. Supported: HttpApi, Operation, and GatewayRoute.
	//
	// 	- IpAccessControl: implements IP address-based access control. Supported: HttpApi, Operation, GatewayRoute, Domain, and Gateway.
	//
	// 	- DirectResponse: mocks responses. Supported: Operation and GatewayRoute.
	//
	// 	- Redirect: redirects traffic. Supported: GatewayRoute.
	//
	// 	- Fallback: implements fallback. Supported: Operation and GatewayRoute.
	//
	// 	- ServiceTls: implements TLS authentication. Supported: GatewayService.
	//
	// 	- ServiceLb: balances loads. Supported: GatewayService.
	//
	// 	- ServicePortTls: implements service port TLS authentication. Supported: GatewayServicePort.
	//
	// 	- Waf: implements WAF protection. Supported: GatewayRoute and Gateway.
	//
	// 	- JWTAuth: implements global JWT authentication. Supported: Gateway.
	//
	// 	- OIDCAuth: implements global OIDC authentication. Supported: Gateway.
	//
	// 	- ExternalZAuth: implements custom authentication. Supported: Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// IpAccessControl
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// The policy configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"enable\\":false}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The policy description.
	//
	// example:
	//
	// This is the policy description.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qasrdc0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The policy name.
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
