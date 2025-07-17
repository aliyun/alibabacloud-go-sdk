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
	// Policy type, including RateLimit, ConcurrencyLimit, CircuitBreaker, HttpRewrite, HeaderModify, Cors, Authentication, FlowCopy, Timeout, Retry, IpAccessControl, DirectResponse, Redirect, Fallback, ServiceTls, ServiceLb, ServicePortTls, Waf, JWTAuth, OIDCAuth, ExternalZAuth, AiProxy, ModelRouter, AiStatistics, AiSecurityGuard, AiFallback, ModelMapper, AiTokenRateLimit, AiCache, DynamicRoute
	//
	// This parameter is required.
	//
	// example:
	//
	// Timeout
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// Policy configuration
	//
	// This parameter is required.
	//
	// example:
	//
	// {"unitNum":1,"timeUnit":"s","enable":true}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// Policy description
	//
	// example:
	//
	// timeout policy
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Policy name
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
