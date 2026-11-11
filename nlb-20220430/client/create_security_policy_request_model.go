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
	SetRegionId(v string) *CreateSecurityPolicyRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateSecurityPolicyRequest
	GetResourceGroupId() *string
	SetSecurityPolicyName(v string) *CreateSecurityPolicyRequest
	GetSecurityPolicyName() *string
	SetTag(v []*CreateSecurityPolicyRequestTag) *CreateSecurityPolicyRequest
	GetTag() []*CreateSecurityPolicyRequestTag
	SetTlsVersions(v []*string) *CreateSecurityPolicyRequest
	GetTlsVersions() []*string
}

type CreateSecurityPolicyRequest struct {
	// The supported cipher suites, which depend on the TLSVersion value. You can add up to 32 cipher suites.
	//
	// Supported by TLSv1.0 and TLSv1.1:
	//
	// - **ECDHE-ECDSA-AES128-SHA**
	//
	// - **ECDHE-ECDSA-AES256-SHA**
	//
	// - **ECDHE-RSA-AES128-SHA**
	//
	// - **ECDHE-RSA-AES256-SHA**
	//
	// - **AES128-SHA**
	//
	// - **AES256-SHA**
	//
	// - **DES-CBC3-SHA**
	//
	// Supported by TLSv1.2:
	//
	// -   **ECDHE-ECDSA-AES128-SHA**
	//
	// -   **ECDHE-ECDSA-AES256-SHA**
	//
	// -   **ECDHE-RSA-AES128-SHA**
	//
	// -   **ECDHE-RSA-AES256-SHA**
	//
	// -   **AES128-SHA**
	//
	// -   **AES256-SHA**
	//
	// -   **DES-CBC3-SHA**
	//
	// -   **ECDHE-ECDSA-AES128-GCM-SHA256**
	//
	// -   **ECDHE-ECDSA-AES256-GCM-SHA384**
	//
	// -   **ECDHE-ECDSA-AES128-SHA256**
	//
	// -   **ECDHE-ECDSA-AES256-SHA384**
	//
	// -   **ECDHE-RSA-AES128-GCM-SHA256**
	//
	// -   **ECDHE-RSA-AES256-GCM-SHA384**
	//
	// -   **ECDHE-RSA-AES128-SHA256**
	//
	// -   **ECDHE-RSA-AES256-SHA384**
	//
	// -   **AES128-GCM-SHA256**
	//
	// -   **AES256-GCM-SHA384**
	//
	// -   **AES128-SHA256**
	//
	// -   **AES256-SHA256**
	//
	// -   **ECDHE-ECDSA-CHACHA20-POLY1305**
	//
	// -   **ECDHE-RSA-CHACHA20-POLY1305**
	//
	// Supported by TLSv1.3:
	//
	// - **TLS_AES_128_GCM_SHA256**
	//
	// - **TLS_AES_256_GCM_SHA384**
	//
	// - **TLS_CHACHA20_POLY1305_SHA256**
	//
	// - **TLS_AES_128_CCM_SHA256**
	//
	// - **TLS_AES_128_CCM_8_SHA256**
	//
	// This parameter is required.
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without creating the custom security policy. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the precheck, an error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the Network Load Balancer (NLB) instance.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Policy Name of the security policy.
	//
	// Policy Name must be 1 to 200 characters in length, and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// TLSCipherPolicy
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// The list of tags.
	//
	// if can be null:
	// true
	Tag []*CreateSecurityPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The supported TLS versions. Valid values: **TLSv1.0**, **TLSv1.1**, **TLSv1.2**, and **TLSv1.3**.
	//
	// This parameter is required.
	TlsVersions []*string `json:"TlsVersions,omitempty" xml:"TlsVersions,omitempty" type:"Repeated"`
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

func (s *CreateSecurityPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSecurityPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSecurityPolicyRequest) GetSecurityPolicyName() *string {
	return s.SecurityPolicyName
}

func (s *CreateSecurityPolicyRequest) GetTag() []*CreateSecurityPolicyRequestTag {
	return s.Tag
}

func (s *CreateSecurityPolicyRequest) GetTlsVersions() []*string {
	return s.TlsVersions
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

func (s *CreateSecurityPolicyRequest) SetRegionId(v string) *CreateSecurityPolicyRequest {
	s.RegionId = &v
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

func (s *CreateSecurityPolicyRequest) SetTag(v []*CreateSecurityPolicyRequestTag) *CreateSecurityPolicyRequest {
	s.Tag = v
	return s
}

func (s *CreateSecurityPolicyRequest) SetTlsVersions(v []*string) *CreateSecurityPolicyRequest {
	s.TlsVersions = v
	return s
}

func (s *CreateSecurityPolicyRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSecurityPolicyRequestTag struct {
	// The tag key. The key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. Supported characters include letters, digits, underscores (_), periods (.), colons (:), forward slashes (/), equal signs (=), plus signs (+), hyphens (-), and at signs (@).
	//
	// You can add up to 20 tags in a single call.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. Supported characters include letters, digits, underscores (_), periods (.), colons (:), forward slashes (/), equal signs (=), plus signs (+), hyphens (-), and at signs (@).
	//
	// You can add up to 20 tags in a single call.
	//
	// example:
	//
	// ValueTest
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
