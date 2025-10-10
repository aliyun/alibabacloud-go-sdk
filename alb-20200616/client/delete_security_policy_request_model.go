// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSecurityPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteSecurityPolicyRequest
	GetDryRun() *bool
	SetSecurityPolicyId(v string) *DeleteSecurityPolicyRequest
	GetSecurityPolicyId() *string
}

type DeleteSecurityPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
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
}

func (s DeleteSecurityPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSecurityPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteSecurityPolicyRequest) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *DeleteSecurityPolicyRequest) SetClientToken(v string) *DeleteSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetDryRun(v bool) *DeleteSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetSecurityPolicyId(v string) *DeleteSecurityPolicyRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) Validate() error {
	return dara.Validate(s)
}
