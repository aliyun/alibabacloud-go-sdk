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
	SetRegionId(v string) *DeleteSecurityPolicyRequest
	GetRegionId() *string
	SetSecurityPolicyId(v string) *DeleteSecurityPolicyRequest
	GetSecurityPolicyId() *string
}

type DeleteSecurityPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a value for this parameter on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run only. Valid values:
	//
	// - **true**: sends a check request and does not delete the custom security policy. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the check fails, the system returns an error. If the check succeeds, the system returns the `DryRunOperation` error code.
	//
	// - **false*	- (default): sends a regular request. If the check succeeds, the system returns an HTTP 2xx status code and performs the operation.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the Network Load Balancer (NLB) instance is deployed.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to obtain region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the TLS security policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-wuytp9pa******
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

func (s *DeleteSecurityPolicyRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *DeleteSecurityPolicyRequest) SetRegionId(v string) *DeleteSecurityPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetSecurityPolicyId(v string) *DeleteSecurityPolicyRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) Validate() error {
	return dara.Validate(s)
}
