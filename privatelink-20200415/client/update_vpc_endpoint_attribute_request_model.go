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
	// The protocol. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **DualStack**
	//
	// >  An endpoint supports dual-stack only if its associated endpoint service and VPC support dual-stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CrossRegionBandwidth *int32  `json:"CrossRegionBandwidth,omitempty" xml:"CrossRegionBandwidth,omitempty"`
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
	// The description of the endpoint.
	//
	// The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The endpoint ID whose attributes you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
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
	// The region ID of the endpoint whose attributes you want to modify. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResetPolicy         *bool   `json:"ResetPolicy,omitempty" xml:"ResetPolicy,omitempty"`
	ZoneAffinityEnabled *bool   `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
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
