// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphers(v []*string) *CreateSecurityPolicyRequest
	GetCiphers() []*string
	SetClientToken(v string) *CreateSecurityPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateSecurityPolicyRequest
	GetDryRun() *bool
	SetResourceGroupId(v string) *CreateSecurityPolicyRequest
	GetResourceGroupId() *string
	SetSecurityPolicyName(v string) *CreateSecurityPolicyRequest
	GetSecurityPolicyName() *string
	SetTLSVersions(v []*string) *CreateSecurityPolicyRequest
	GetTLSVersions() []*string
	SetTag(v []*CreateSecurityPolicyRequestTag) *CreateSecurityPolicyRequest
	GetTag() []*CreateSecurityPolicyRequestTag
}

type CreateSecurityPolicyRequest struct {
	// The supported cipher suites.
	//
	// This parameter is required.
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
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
	// 	- **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the security policy.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-secrity
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// The supported Transport Layer Security (TLS) protocol versions.
	//
	// This parameter is required.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
	// The tags.
	Tag []*CreateSecurityPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateSecurityPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyRequest) GetCiphers() []*string {
	return s.Ciphers
}

func (s *CreateSecurityPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSecurityPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateSecurityPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSecurityPolicyRequest) GetSecurityPolicyName() *string {
	return s.SecurityPolicyName
}

func (s *CreateSecurityPolicyRequest) GetTLSVersions() []*string {
	return s.TLSVersions
}

func (s *CreateSecurityPolicyRequest) GetTag() []*CreateSecurityPolicyRequestTag {
	return s.Tag
}

func (s *CreateSecurityPolicyRequest) SetCiphers(v []*string) *CreateSecurityPolicyRequest {
	s.Ciphers = v
	return s
}

func (s *CreateSecurityPolicyRequest) SetClientToken(v string) *CreateSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetDryRun(v bool) *CreateSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetResourceGroupId(v string) *CreateSecurityPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetSecurityPolicyName(v string) *CreateSecurityPolicyRequest {
	s.SecurityPolicyName = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetTLSVersions(v []*string) *CreateSecurityPolicyRequest {
	s.TLSVersions = v
	return s
}

func (s *CreateSecurityPolicyRequest) SetTag(v []*CreateSecurityPolicyRequestTag) *CreateSecurityPolicyRequest {
	s.Tag = v
	return s
}

func (s *CreateSecurityPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type CreateSecurityPolicyRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSecurityPolicyRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateSecurityPolicyRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateSecurityPolicyRequestTag) SetKey(v string) *CreateSecurityPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSecurityPolicyRequestTag) SetValue(v string) *CreateSecurityPolicyRequestTag {
	s.Value = &v
	return s
}

func (s *CreateSecurityPolicyRequestTag) Validate() error {
	return dara.Validate(s)
}
