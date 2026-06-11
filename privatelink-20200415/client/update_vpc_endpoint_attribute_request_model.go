// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *UpdateVpcEndpointAttributeRequest
	GetAddressIpVersion() *string
	SetClientToken(v string) *UpdateVpcEndpointAttributeRequest
	GetClientToken() *string
	SetCrossRegionBandwidth(v int32) *UpdateVpcEndpointAttributeRequest
	GetCrossRegionBandwidth() *int32
	SetDryRun(v bool) *UpdateVpcEndpointAttributeRequest
	GetDryRun() *bool
	SetEndpointDescription(v string) *UpdateVpcEndpointAttributeRequest
	GetEndpointDescription() *string
	SetEndpointId(v string) *UpdateVpcEndpointAttributeRequest
	GetEndpointId() *string
	SetEndpointName(v string) *UpdateVpcEndpointAttributeRequest
	GetEndpointName() *string
	SetPolicyDocument(v string) *UpdateVpcEndpointAttributeRequest
	GetPolicyDocument() *string
	SetRegionId(v string) *UpdateVpcEndpointAttributeRequest
	GetRegionId() *string
	SetResetPolicy(v bool) *UpdateVpcEndpointAttributeRequest
	GetResetPolicy() *bool
	SetZoneAffinityEnabled(v bool) *UpdateVpcEndpointAttributeRequest
	GetZoneAffinityEnabled() *bool
}

type UpdateVpcEndpointAttributeRequest struct {
	// The protocol version. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **DualStack**: dual-stack.
	//
	// > To enable dual-stack for an endpoint, the associated endpoint service and the endpoint\\"s VPC must also support dual-stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You must generate a unique value for this parameter. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cross-region bandwidth in Mbps. This parameter applies only when the endpoint and its associated endpoint service are in different regions. Valid values:
	//
	// - **Minimum value**: 100.
	//
	// - **Maximum value**: The value is limited by your account quota. For more information, see [Quotas and limits](https://help.aliyun.com/zh/privatelink/quotas-and-limits?spm=a2c4g.11174283.help-menu-search-120462.d_0).
	//
	// > You can specify this parameter only for cross-region endpoints.
	//
	// example:
	//
	// 1000
	CrossRegionBandwidth *int32 `json:"CrossRegionBandwidth,omitempty" xml:"CrossRegionBandwidth,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run. The system checks the required parameters, request format, and service limits. If the request fails the check, the system returns an error message. If the request passes the check, the system returns the `DryRunOperation` error code.
	//
	// - **false*	- (Default): Sends a normal request. If the request passes the check, the system performs the operation and returns a 2xx HTTP status code.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new description for the endpoint.
	//
	// The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The ID of the endpoint to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The new name for the endpoint.
	//
	// The name must be 2 to 128 characters in length, start with a letter or a Chinese character, and can contain digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The RAM access policy. For more information, see [Basic elements of a RAM policy](https://help.aliyun.com/document_detail/93738.html).
	//
	// example:
	//
	// {
	//
	//   "Version": "1",
	//
	//   "Statement": [
	//
	//     {
	//
	//       "Effect": "Allow",
	//
	//       "Action": [
	//
	//         "oss:List*",
	//
	//         "oss:PutObject",
	//
	//         "oss:GetObject"
	//
	//       ],
	//
	//       "Resource": [
	//
	//         "acs:oss:oss-*:*:pvl-policy-test/policy-test.txt"
	//
	//       ],
	//
	//       "Principal": {
	//
	//         "RAM": [
	//
	//           "acs:ram::14199xxxxxx:*"
	//
	//         ]
	//
	//       }
	//
	//     }
	//
	//   ]
	//
	// }
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The region ID of the endpoint. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResetPolicy *bool   `json:"ResetPolicy,omitempty" xml:"ResetPolicy,omitempty"`
	// Specifies whether to enable zone affinity for domain name resolution of the endpoint service. Valid values:
	//
	// - **true**: Enables zone affinity.
	//
	// - **false**: Disables zone affinity.
	//
	// example:
	//
	// false
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s UpdateVpcEndpointAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointAttributeRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *UpdateVpcEndpointAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateVpcEndpointAttributeRequest) GetCrossRegionBandwidth() *int32 {
	return s.CrossRegionBandwidth
}

func (s *UpdateVpcEndpointAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateVpcEndpointAttributeRequest) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *UpdateVpcEndpointAttributeRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateVpcEndpointAttributeRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *UpdateVpcEndpointAttributeRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *UpdateVpcEndpointAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVpcEndpointAttributeRequest) GetResetPolicy() *bool {
	return s.ResetPolicy
}

func (s *UpdateVpcEndpointAttributeRequest) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *UpdateVpcEndpointAttributeRequest) SetAddressIpVersion(v string) *UpdateVpcEndpointAttributeRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetCrossRegionBandwidth(v int32) *UpdateVpcEndpointAttributeRequest {
	s.CrossRegionBandwidth = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetEndpointDescription(v string) *UpdateVpcEndpointAttributeRequest {
	s.EndpointDescription = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetEndpointId(v string) *UpdateVpcEndpointAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetEndpointName(v string) *UpdateVpcEndpointAttributeRequest {
	s.EndpointName = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetPolicyDocument(v string) *UpdateVpcEndpointAttributeRequest {
	s.PolicyDocument = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetResetPolicy(v bool) *UpdateVpcEndpointAttributeRequest {
	s.ResetPolicy = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) SetZoneAffinityEnabled(v bool) *UpdateVpcEndpointAttributeRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *UpdateVpcEndpointAttributeRequest) Validate() error {
	return dara.Validate(s)
}
