// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityPolicyAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphers(v []*string) *UpdateSecurityPolicyAttributeRequest
	GetCiphers() []*string
	SetClientToken(v string) *UpdateSecurityPolicyAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateSecurityPolicyAttributeRequest
	GetDryRun() *bool
	SetSecurityPolicyId(v string) *UpdateSecurityPolicyAttributeRequest
	GetSecurityPolicyId() *string
	SetSecurityPolicyName(v string) *UpdateSecurityPolicyAttributeRequest
	GetSecurityPolicyName() *string
	SetTLSVersions(v []*string) *UpdateSecurityPolicyAttributeRequest
	GetTLSVersions() []*string
}

type UpdateSecurityPolicyAttributeRequest struct {
	// The supported cipher suites.
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The security policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// spy-n0kn923****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The name of the security policy.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test-secrity
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// The supported TLS protocol versions.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s UpdateSecurityPolicyAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeRequest) GetCiphers() []*string {
	return s.Ciphers
}

func (s *UpdateSecurityPolicyAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSecurityPolicyAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateSecurityPolicyAttributeRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *UpdateSecurityPolicyAttributeRequest) GetSecurityPolicyName() *string {
	return s.SecurityPolicyName
}

func (s *UpdateSecurityPolicyAttributeRequest) GetTLSVersions() []*string {
	return s.TLSVersions
}

func (s *UpdateSecurityPolicyAttributeRequest) SetCiphers(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.Ciphers = v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetClientToken(v string) *UpdateSecurityPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetDryRun(v bool) *UpdateSecurityPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyId(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyName(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyName = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetTLSVersions(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.TLSVersions = v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) Validate() error {
	return dara.Validate(s)
}
