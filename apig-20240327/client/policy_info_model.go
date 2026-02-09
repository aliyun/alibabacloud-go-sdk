// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolicyInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAttachments(v []*Attachment) *PolicyInfo
	GetAttachments() []*Attachment
	SetClassAlias(v string) *PolicyInfo
	GetClassAlias() *string
	SetClassName(v string) *PolicyInfo
	GetClassName() *string
	SetConfig(v string) *PolicyInfo
	GetConfig() *string
	SetDirection(v string) *PolicyInfo
	GetDirection() *string
	SetExecutePriority(v string) *PolicyInfo
	GetExecutePriority() *string
	SetExecuteStage(v string) *PolicyInfo
	GetExecuteStage() *string
	SetName(v string) *PolicyInfo
	GetName() *string
	SetPolicyId(v string) *PolicyInfo
	GetPolicyId() *string
	SetType(v string) *PolicyInfo
	GetType() *string
}

type PolicyInfo struct {
	// The mount information.
	Attachments []*Attachment `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// The policy class alias.
	ClassAlias *string `json:"classAlias,omitempty" xml:"classAlias,omitempty"`
	// The class name supported by the policy. Different policies support different mount points. This parameter is used in combination with AttachResourceType.
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
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// The policy configurations.
	//
	// example:
	//
	// {"enable":false}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The direction of traffic on which the policy takes effect. Valid values:
	//
	// 	- OutBound
	//
	// 	- InBound
	//
	// 	- Both
	//
	// example:
	//
	// InBound
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The execution priority.
	//
	// example:
	//
	// 310
	ExecutePriority *string `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	// The execution phase.
	//
	// Valid values:
	//
	// 	- PluginStatistic
	//
	// 	- PluginAuthorization
	//
	// 	- PluginPre
	//
	// 	- PluginAuthentication
	//
	// 	- PluginDefault
	//
	// 	- PluginPost
	//
	// example:
	//
	// PluginPost
	ExecuteStage *string `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	// The policy name.
	//
	// example:
	//
	// test-policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// p-cq7l5s5lhtgi6qasrdc0
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// The policy type.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PolicyInfo) String() string {
	return dara.Prettify(s)
}

func (s PolicyInfo) GoString() string {
	return s.String()
}

func (s *PolicyInfo) GetAttachments() []*Attachment {
	return s.Attachments
}

func (s *PolicyInfo) GetClassAlias() *string {
	return s.ClassAlias
}

func (s *PolicyInfo) GetClassName() *string {
	return s.ClassName
}

func (s *PolicyInfo) GetConfig() *string {
	return s.Config
}

func (s *PolicyInfo) GetDirection() *string {
	return s.Direction
}

func (s *PolicyInfo) GetExecutePriority() *string {
	return s.ExecutePriority
}

func (s *PolicyInfo) GetExecuteStage() *string {
	return s.ExecuteStage
}

func (s *PolicyInfo) GetName() *string {
	return s.Name
}

func (s *PolicyInfo) GetPolicyId() *string {
	return s.PolicyId
}

func (s *PolicyInfo) GetType() *string {
	return s.Type
}

func (s *PolicyInfo) SetAttachments(v []*Attachment) *PolicyInfo {
	s.Attachments = v
	return s
}

func (s *PolicyInfo) SetClassAlias(v string) *PolicyInfo {
	s.ClassAlias = &v
	return s
}

func (s *PolicyInfo) SetClassName(v string) *PolicyInfo {
	s.ClassName = &v
	return s
}

func (s *PolicyInfo) SetConfig(v string) *PolicyInfo {
	s.Config = &v
	return s
}

func (s *PolicyInfo) SetDirection(v string) *PolicyInfo {
	s.Direction = &v
	return s
}

func (s *PolicyInfo) SetExecutePriority(v string) *PolicyInfo {
	s.ExecutePriority = &v
	return s
}

func (s *PolicyInfo) SetExecuteStage(v string) *PolicyInfo {
	s.ExecuteStage = &v
	return s
}

func (s *PolicyInfo) SetName(v string) *PolicyInfo {
	s.Name = &v
	return s
}

func (s *PolicyInfo) SetPolicyId(v string) *PolicyInfo {
	s.PolicyId = &v
	return s
}

func (s *PolicyInfo) SetType(v string) *PolicyInfo {
	s.Type = &v
	return s
}

func (s *PolicyInfo) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
