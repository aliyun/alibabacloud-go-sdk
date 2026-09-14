// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *CreatePolicyRequest
	GetClassName() *string
	SetConfig(v string) *CreatePolicyRequest
	GetConfig() *string
	SetDescription(v string) *CreatePolicyRequest
	GetDescription() *string
	SetName(v string) *CreatePolicyRequest
	GetName() *string
}

type CreatePolicyRequest struct {
	// The policy type. Valid values:
	//
	// - RateLimit: Rate limiting. Limits the request rate.
	//
	// - ConcurrencyLimit: Concurrency limiting. Limits the number of concurrent requests.
	//
	// - CircuitBreaker: Circuit breaker. Automatically triggers circuit breaking when backend exceptions occur.
	//
	// - HttpRewrite: HTTP rewrite. Rewrites the request URL or path.
	//
	// - HeaderModify: Header modification. Adds, removes, or modifies HTTP request headers.
	//
	// - Cors: Cross-Origin Resource Sharing (CORS). Controls cross-origin requests.
	//
	// - Authentication: General authentication. A general request authentication policy.
	//
	// - FlowCopy: Traffic mirroring. Copies requests to an additional backend.
	//
	// - Timeout: Timeout. Sets the request timeout period.
	//
	// - Retry: Retry. Automatically retries failed requests.
	//
	// - IpAccessControl: IP access control. Filters requests based on IP address whitelists and blacklists.
	//
	// - DirectResponse: Direct response. Returns a fixed response directly.
	//
	// - Redirect: Redirect. Redirects requests to another address.
	//
	// - Fallback: Fallback. Returns a fallback response when the backend is unavailable.
	//
	// - ServiceTls: Service TLS. Configures TLS for backend services.
	//
	// - ServiceLb: Service load balancing. Configures load balancing for backend services.
	//
	// - ServicePortTls: Service port TLS. Configures TLS for backend service ports.
	//
	// - Waf: Web Application Firewall (WAF). Provides request security protection.
	//
	// - JWTAuth: JWT authentication. Authenticates requests based on JSON Web Tokens (JWT).
	//
	// - OIDCAuth: OIDC authentication. Authenticates requests based on the OpenID Connect (OIDC) protocol.
	//
	// - ExternalZAuth: External authentication. Integrates with an external authentication service.
	//
	// - AiProxy: AI proxy.
	//
	// - ModelRouter: Model router.
	//
	// - AiStatistics: AI statistics.
	//
	// - AiSecurityGuard: AI security guard. Detects the security of AI request and response content.
	//
	// - AiFallback: AI fallback. Falls back to an alternative model when the AI service is unavailable.
	//
	// - ModelMapper: Model mapper.
	//
	// - AiTokenRateLimit: AI token rate limiting. Limits the rate based on token consumption.
	//
	// - AiCache: AI cache. Caches AI response results.
	//
	// - DynamicRoute: Dynamic route.
	//
	// This parameter is required.
	//
	// example:
	//
	// Timeout
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// The policy configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"unitNum":1,"timeUnit":"s","enable":true}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// Timeout policy
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The policy name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetClassName() *string {
	return s.ClassName
}

func (s *CreatePolicyRequest) GetConfig() *string {
	return s.Config
}

func (s *CreatePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreatePolicyRequest) SetClassName(v string) *CreatePolicyRequest {
	s.ClassName = &v
	return s
}

func (s *CreatePolicyRequest) SetConfig(v string) *CreatePolicyRequest {
	s.Config = &v
	return s
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetName(v string) *CreatePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
