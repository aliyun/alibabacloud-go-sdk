// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafPolicyDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindDomains(v string) *ModifyDcdnWafPolicyDomainsRequest
	GetBindDomains() *string
	SetMethod(v int32) *ModifyDcdnWafPolicyDomainsRequest
	GetMethod() *int32
	SetPolicyId(v int64) *ModifyDcdnWafPolicyDomainsRequest
	GetPolicyId() *int64
	SetUnbindDomains(v string) *ModifyDcdnWafPolicyDomainsRequest
	GetUnbindDomains() *string
}

type ModifyDcdnWafPolicyDomainsRequest struct {
	// The domain names that you want to bind to the protection policy. You can specify up to 50 domain names. Separate multiple domain names with commas (,).
	//
	// > You can configure either **BindDomains*	- or **UnbindDomains**.
	//
	// example:
	//
	// example.com,example2.com
	BindDomains *string `json:"BindDomains,omitempty" xml:"BindDomains,omitempty"`
	// The association method. Valid values:
	//
	// 	- 0: replace.
	//
	// 	- 1: add.
	//
	// 	- Default value: 0.
	//
	// >
	//
	// 	- This parameter takes effect only when you specify **BindDomains**. If you have associated a domain name indicated by **BindDomains*	- with the default protection policy, the `Policy.DefaultAndCustom.BindToSameDomain` error is returned.
	//
	// 	- You can only replace accelerated domain names that are bound to the default protection policy.
	//
	// example:
	//
	// 0
	Method *int32 `json:"Method,omitempty" xml:"Method,omitempty"`
	// The ID of the protection policy. You can specify only one ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The domain names that you want to unbind from the protection policy. You can specify up to 50 domain names. Separate multiple domain names with commas (,).
	//
	// > You can configure either **BindDomains*	- or **UnbindDomains**.
	//
	// example:
	//
	// example3.com
	UnbindDomains *string `json:"UnbindDomains,omitempty" xml:"UnbindDomains,omitempty"`
}

func (s ModifyDcdnWafPolicyDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafPolicyDomainsRequest) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafPolicyDomainsRequest) GetBindDomains() *string {
	return s.BindDomains
}

func (s *ModifyDcdnWafPolicyDomainsRequest) GetMethod() *int32 {
	return s.Method
}

func (s *ModifyDcdnWafPolicyDomainsRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ModifyDcdnWafPolicyDomainsRequest) GetUnbindDomains() *string {
	return s.UnbindDomains
}

func (s *ModifyDcdnWafPolicyDomainsRequest) SetBindDomains(v string) *ModifyDcdnWafPolicyDomainsRequest {
	s.BindDomains = &v
	return s
}

func (s *ModifyDcdnWafPolicyDomainsRequest) SetMethod(v int32) *ModifyDcdnWafPolicyDomainsRequest {
	s.Method = &v
	return s
}

func (s *ModifyDcdnWafPolicyDomainsRequest) SetPolicyId(v int64) *ModifyDcdnWafPolicyDomainsRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyDcdnWafPolicyDomainsRequest) SetUnbindDomains(v string) *ModifyDcdnWafPolicyDomainsRequest {
	s.UnbindDomains = &v
	return s
}

func (s *ModifyDcdnWafPolicyDomainsRequest) Validate() error {
	return dara.Validate(s)
}
