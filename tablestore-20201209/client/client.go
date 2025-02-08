// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeResourceGroupRequest struct {
	// The ID of the new resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek3bgek3kxhyky
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The resource ID, which is the instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// first-ins
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CheckInstancePolicyRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// workshop-bj-ots1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance policy in the JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "Version": "1",
	//
	//     "Statement": [
	//
	//         {
	//
	//             "Action": [
	//
	//                 "ots:*"
	//
	//             ],
	//
	//             "Resource": [
	//
	//                 "acs:ots:*:13791xxxxxxxxxxx:instance/myinstance*"
	//
	//             ],
	//
	//             "Principal": [
	//
	//                 "*"
	//
	//             ],
	//
	//             "Effect": "Allow",
	//
	//             "Condition": {
	//
	//                 "StringEquals": {
	//
	//                     "ots:TLSVersion": [
	//
	//                         "1.2"
	//
	//                     ]
	//
	//                 },
	//
	//                 "IpAddress": {
	//
	//                     "acs:SourceIp": [
	//
	//                         "192.168.0.1",
	//
	//                         "172.16.0.1"
	//
	//                     ]
	//
	//                 }
	//
	//             }
	//
	//         }
	//
	//     ]
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CheckInstancePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckInstancePolicyRequest) GoString() string {
	return s.String()
}

func (s *CheckInstancePolicyRequest) SetInstanceName(v string) *CheckInstancePolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *CheckInstancePolicyRequest) SetPolicy(v string) *CheckInstancePolicyRequest {
	s.Policy = &v
	return s
}

type CheckInstancePolicyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Verification passed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 757E172A-F94B-5E78-8A23-D9068E42F2E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckInstancePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckInstancePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstancePolicyResponseBody) SetCode(v string) *CheckInstancePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CheckInstancePolicyResponseBody) SetMessage(v string) *CheckInstancePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CheckInstancePolicyResponseBody) SetRequestId(v string) *CheckInstancePolicyResponseBody {
	s.RequestId = &v
	return s
}

type CheckInstancePolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstancePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstancePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *CheckInstancePolicyResponse) SetHeaders(v map[string]*string) *CheckInstancePolicyResponse {
	s.Headers = v
	return s
}

func (s *CheckInstancePolicyResponse) SetStatusCode(v int32) *CheckInstancePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstancePolicyResponse) SetBody(v *CheckInstancePolicyResponseBody) *CheckInstancePolicyResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// The type of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// SSD
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// (Deprecated) Specifies whether to enable disaster recovery for the instance.
	//
	// Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	DisableReplication *bool `json:"DisableReplication,omitempty" xml:"DisableReplication,omitempty"`
	// The description of the instance. The instance description must be 3 to 256 characters in length.
	//
	// example:
	//
	// the test instance
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance. Instance naming conventions:
	//
	// 	- The name can contain only letters, digits, and hyphens (-).
	//
	// 	- The name must start with a letter.
	//
	// 	- The name cannot end with a hyphen (-).
	//
	// 	- The name is case-insensitive.
	//
	// 	- The name must be 3 to 16 characters in length.
	//
	// 	- The name cannot contain the following words: ali, ay, ots, taobao, and admin.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// (Deprecated) The network type of the instance. Valid values: NORMAL and VPC_CONSOLE. Default value: NORMAL.
	//
	// example:
	//
	// NORMAL
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The types of the source from which access is allowed. By default, the following source type is allowed:
	//
	// TRUST_PROXY: console
	NetworkSourceACL []*string `json:"NetworkSourceACL,omitempty" xml:"NetworkSourceACL,omitempty" type:"Repeated"`
	// The types of the network from which access is allowed. By default, the following network types are allowed:
	//
	// 	- INTERNET: Internet
	//
	// 	- VPC: virtual private cloud (VPC)
	//
	// 	- CLASSIC: classic network
	NetworkTypeACL []*string `json:"NetworkTypeACL,omitempty" xml:"NetworkTypeACL,omitempty" type:"Repeated"`
	// The instance policy in the JSON format.
	//
	// example:
	//
	// {
	//
	//     "Version": "1",
	//
	//     "Statement": [
	//
	//         {
	//
	//             "Action": [
	//
	//                 "ots:*"
	//
	//             ],
	//
	//             "Resource": [
	//
	//                 "acs:ots:*:13791xxxxxxxxxxx:instance/myinstance*"
	//
	//             ],
	//
	//             "Principal": [
	//
	//                 "*"
	//
	//             ],
	//
	//             "Effect": "Allow",
	//
	//             "Condition": {
	//
	//                 "StringEquals": {
	//
	//                     "ots:TLSVersion": [
	//
	//                         "1.2"
	//
	//                     ]
	//
	//                 },
	//
	//                 "IpAddress": {
	//
	//                     "acs:SourceIp": [
	//
	//                         "192.168.0.1",
	//
	//                         "198.51.100.1"
	//
	//                     ]
	//
	//                 }
	//
	//             }
	//
	//         }
	//
	//     ]
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmxh4em5jnbcd
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*CreateInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetClusterType(v string) *CreateInstanceRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateInstanceRequest) SetDisableReplication(v bool) *CreateInstanceRequest {
	s.DisableReplication = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceDescription(v string) *CreateInstanceRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetNetwork(v string) *CreateInstanceRequest {
	s.Network = &v
	return s
}

func (s *CreateInstanceRequest) SetNetworkSourceACL(v []*string) *CreateInstanceRequest {
	s.NetworkSourceACL = v
	return s
}

func (s *CreateInstanceRequest) SetNetworkTypeACL(v []*string) *CreateInstanceRequest {
	s.NetworkTypeACL = v
	return s
}

func (s *CreateInstanceRequest) SetPolicy(v string) *CreateInstanceRequest {
	s.Policy = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

type CreateInstanceRequestTags struct {
	// The tag key. The tag key can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// p_instance
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0woauavextilfqr61
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

type CreateInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceName(v string) *DeleteInstanceRequest {
	s.InstanceName = &v
	return s
}

type DeleteInstanceResponseBody struct {
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 17827FB1-CC71-559D-BBDF-9D41534322E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteInstancePolicyRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-test-12345
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The version of the instance policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s DeleteInstancePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancePolicyRequest) SetInstanceName(v string) *DeleteInstancePolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteInstancePolicyRequest) SetPolicyVersion(v int64) *DeleteInstancePolicyRequest {
	s.PolicyVersion = &v
	return s
}

type DeleteInstancePolicyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 3104C83E-6E82-57FB-BB88-8C64CCFDEF89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstancePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancePolicyResponseBody) SetCode(v string) *DeleteInstancePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstancePolicyResponseBody) SetMessage(v string) *DeleteInstancePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstancePolicyResponseBody) SetRequestId(v string) *DeleteInstancePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstancePolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstancePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstancePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstancePolicyResponse) SetHeaders(v map[string]*string) *DeleteInstancePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstancePolicyResponse) SetStatusCode(v int32) *DeleteInstancePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstancePolicyResponse) SetBody(v *DeleteInstancePolicyResponseBody) *DeleteInstancePolicyResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// >  #### You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetClientToken(v string) *DescribeRegionsRequest {
	s.ClientToken = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// B37BBA04-D827-55C8-B901-5264B904E8C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The key of the region.
	//
	// example:
	//
	// yunjiankong.cloud.region.cn-hangzhou
	I18nKey *string `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetI18nKey(v string) *DescribeRegionsResponseBodyRegions {
	s.I18nKey = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type GetInstanceRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetInstanceName(v string) *GetInstanceRequest {
	s.InstanceName = &v
	return s
}

type GetInstanceResponseBody struct {
	// The instance alias.
	//
	// example:
	//
	// instance-test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2019-12-23T07:24:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The upper limit for the VCUs of the instance.
	//
	// example:
	//
	// 6
	ElasticVCUUpperLimit *float32 `json:"ElasticVCUUpperLimit,omitempty" xml:"ElasticVCUUpperLimit,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Description of the test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the instance.
	//
	// 	- SSD: high-performance instance
	//
	// 	- HYBRID: capacity instance
	//
	// example:
	//
	// SSD
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The status of the instance.
	//
	// 	- normal: The instance works as expected.
	//
	// 	- forbidden: The instance is disabled.
	//
	// 	- Deleting: The instance is being deleted.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Indicates whether zone-redundant storage (ZRS) is enabled for the instance.
	//
	// 	- true: ZRS is enabled for the instance.
	//
	// 	- false: Locally redundant storage (LRS) is enabled for the instance.
	//
	// example:
	//
	// true
	IsMultiAZ *bool `json:"IsMultiAZ,omitempty" xml:"IsMultiAZ,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- VPC: The instance can be accessed only over the bound virtual private clouds (VPCs).
	//
	// 	- VPC_CONSOLE: The instance can be accessed from the Tablestore console or over the bound VPCs.
	//
	// 	- NORMAL: The instance can be accessed from networks of the custom types.
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The sources of the network from which access is allowed. Valid value:
	//
	// TRUST_PROXY: console
	NetworkSourceACL []*string `json:"NetworkSourceACL,omitempty" xml:"NetworkSourceACL,omitempty" type:"Repeated"`
	// The types of the network from which access is allowed. Valid values:
	//
	// 	- CLASSIC: the classic network
	//
	// 	- INTERNET: the Internet
	//
	// 	- VPC: VPCs
	NetworkTypeACL []*string `json:"NetworkTypeACL,omitempty" xml:"NetworkTypeACL,omitempty" type:"Repeated"`
	// The billing method. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The instance policy.
	//
	// example:
	//
	// {
	//
	//     "Version": "1",
	//
	//     "Statement": [
	//
	//         {
	//
	//             "Action": [
	//
	//                 "ots:*"
	//
	//             ],
	//
	//             "Resource": [
	//
	//                 "acs:ots:*:13791xxxxxxxxxxx:instance/myinstance*"
	//
	//             ],
	//
	//             "Principal": [
	//
	//                 "*"
	//
	//             ],
	//
	//             "Effect": "Allow",
	//
	//             "Condition": {
	//
	//                 "StringEquals": {
	//
	//                     "ots:TLSVersion": [
	//
	//                         "1.2"
	//
	//                     ]
	//
	//                 },
	//
	//                 "IpAddress": {
	//
	//                     "acs:SourceIp": [
	//
	//                         "192.168.0.1",
	//
	//                         "198.51.100.1"
	//
	//                     ]
	//
	//                 }
	//
	//             }
	//
	//         }
	//
	//     ]
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The version of the instance policy.
	//
	// example:
	//
	// 1
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 757E172A-F94B-5E78-8A23-D9068E42F2E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmxh4em5jncda
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ots_standard_public_cn-9lb34u7u5001
	SPInstanceId *string `json:"SPInstanceId,omitempty" xml:"SPInstanceId,omitempty"`
	// The storage type.
	//
	// 	- SSD: high-performance
	//
	// 	- HYBRID: capacity
	//
	// example:
	//
	// HYBRID
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The total number of tables in the instance.
	//
	// example:
	//
	// 100
	TableQuota *int32 `json:"TableQuota,omitempty" xml:"TableQuota,omitempty"`
	// The tags of the instance.
	Tags []*GetInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 16542312566
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The VCU quota.
	//
	// example:
	//
	// 20
	VCUQuota *int32 `json:"VCUQuota,omitempty" xml:"VCUQuota,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetAliasName(v string) *GetInstanceResponseBody {
	s.AliasName = &v
	return s
}

func (s *GetInstanceResponseBody) SetCreateTime(v string) *GetInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetElasticVCUUpperLimit(v float32) *GetInstanceResponseBody {
	s.ElasticVCUUpperLimit = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceDescription(v string) *GetInstanceResponseBody {
	s.InstanceDescription = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceName(v string) *GetInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceSpecification(v string) *GetInstanceResponseBody {
	s.InstanceSpecification = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceStatus(v string) *GetInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceResponseBody) SetIsMultiAZ(v bool) *GetInstanceResponseBody {
	s.IsMultiAZ = &v
	return s
}

func (s *GetInstanceResponseBody) SetNetwork(v string) *GetInstanceResponseBody {
	s.Network = &v
	return s
}

func (s *GetInstanceResponseBody) SetNetworkSourceACL(v []*string) *GetInstanceResponseBody {
	s.NetworkSourceACL = v
	return s
}

func (s *GetInstanceResponseBody) SetNetworkTypeACL(v []*string) *GetInstanceResponseBody {
	s.NetworkTypeACL = v
	return s
}

func (s *GetInstanceResponseBody) SetPaymentType(v string) *GetInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceResponseBody) SetPolicy(v string) *GetInstanceResponseBody {
	s.Policy = &v
	return s
}

func (s *GetInstanceResponseBody) SetPolicyVersion(v int64) *GetInstanceResponseBody {
	s.PolicyVersion = &v
	return s
}

func (s *GetInstanceResponseBody) SetRegionId(v string) *GetInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResourceGroupId(v string) *GetInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSPInstanceId(v string) *GetInstanceResponseBody {
	s.SPInstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetStorageType(v string) *GetInstanceResponseBody {
	s.StorageType = &v
	return s
}

func (s *GetInstanceResponseBody) SetTableQuota(v int32) *GetInstanceResponseBody {
	s.TableQuota = &v
	return s
}

func (s *GetInstanceResponseBody) SetTags(v []*GetInstanceResponseBodyTags) *GetInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBody) SetUserId(v string) *GetInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBody) SetVCUQuota(v int32) *GetInstanceResponseBody {
	s.VCUQuota = &v
	return s
}

type GetInstanceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// (Deprecated) The tag key.
	//
	// example:
	//
	// keyTestA
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// (Deprecated) The tag value.
	//
	// example:
	//
	// 00004a20240452b0
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 333
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyTags) SetKey(v string) *GetInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetTagKey(v string) *GetInstanceResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetTagValue(v string) *GetInstanceResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetValue(v string) *GetInstanceResponseBodyTags {
	s.Value = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// The name of the instance. Fuzzy search is supported.
	//
	// example:
	//
	// instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The names of the instances. This parameter is used to specify multiple instances that you want to query at the same time.
	InstanceNameList []*string `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty" type:"Repeated"`
	// The maximum number of instances that you want to return. Valid values: 0 to 200. If you do not configure this parameter or set this parameter to 0, the default value of 100 is used.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call. Instances are returned in lexicographical order starting from the position that is specified by this parameter. The first time you call the operation, leave this parameter empty.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID. You can query the ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-aek24upgom6p5ri
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// normal
	Status *string                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*ListInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceNameList(v []*string) *ListInstancesRequest {
	s.InstanceNameList = v
	return s
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

type ListInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

type ListInstancesShrinkRequest struct {
	// The name of the instance. Fuzzy search is supported.
	//
	// example:
	//
	// instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The names of the instances. This parameter is used to specify multiple instances that you want to query at the same time.
	InstanceNameListShrink *string `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty"`
	// The maximum number of instances that you want to return. Valid values: 0 to 200. If you do not configure this parameter or set this parameter to 0, the default value of 100 is used.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call. Instances are returned in lexicographical order starting from the position that is specified by this parameter. The first time you call the operation, leave this parameter empty.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID. You can query the ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-aek24upgom6p5ri
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// normal
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) SetInstanceName(v string) *ListInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceNameListShrink(v string) *ListInstancesShrinkRequest {
	s.InstanceNameListShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetMaxResults(v int32) *ListInstancesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetNextToken(v string) *ListInstancesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceGroupId(v string) *ListInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetStatus(v string) *ListInstancesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTagShrink(v string) *ListInstancesShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListInstancesResponseBody struct {
	// The details about the instances.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The token that determines the start position of the next query. If this parameter is empty, all instances that you want to query are returned.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// E734979F-5A44-5993-9CE5-C23103576923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetNextToken(v string) *ListInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	// The instance alias.
	//
	// example:
	//
	// instance-test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2019-04-07T09:19:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// Description of the test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance, which is used to uniquely identify the instance.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the instance.
	//
	// 	- SSD: high-performance instance
	//
	// 	- HYBRID: capacity instance
	//
	// example:
	//
	// HYBRID
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The instance status.
	//
	// 	- normal: The instance works as expected.
	//
	// 	- forbidden: The instance is disabled.
	//
	// 	- deleting: The instance is being deleted.
	//
	// example:
	//
	// normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Indicates whether zone-redundant storage (ZRS) is enabled for the instance.
	//
	// 	- true: ZRS is enabled for the instance.
	//
	// 	- false: Locally redundant storage (LRS) is enabled for the instance.
	//
	// example:
	//
	// true
	IsMultiAZ *bool `json:"IsMultiAZ,omitempty" xml:"IsMultiAZ,omitempty"`
	// The billing method.
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay as you go
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxh4em5jnbcd
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ots_standard_public_cn-g4t3igqjj002
	SPInstanceId *string `json:"SPInstanceId,omitempty" xml:"SPInstanceId,omitempty"`
	// The storage type.
	//
	// 	- SSD: high-performance
	//
	// 	- HYBRID: capacity
	//
	// example:
	//
	// HYBRID
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 13542356466
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The VCU quota.
	//
	// example:
	//
	// 3
	VCUQuota *int32 `json:"VCUQuota,omitempty" xml:"VCUQuota,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetAliasName(v string) *ListInstancesResponseBodyInstances {
	s.AliasName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceDescription(v string) *ListInstancesResponseBodyInstances {
	s.InstanceDescription = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceSpecification(v string) *ListInstancesResponseBodyInstances {
	s.InstanceSpecification = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceStatus(v string) *ListInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetIsMultiAZ(v bool) *ListInstancesResponseBodyInstances {
	s.IsMultiAZ = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPaymentType(v string) *ListInstancesResponseBodyInstances {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceGroupId(v string) *ListInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetSPInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.SPInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStorageType(v string) *ListInstancesResponseBodyInstances {
	s.StorageType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserId(v string) *ListInstancesResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetVCUQuota(v int32) *ListInstancesResponseBodyInstances {
	s.VCUQuota = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The maximum number of tagged resources that you want to return. Valid values: 0 to 200. If you do not specify this parameter or set the parameter to 0, the default value of 100 is used.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. Tagged resources are returned in lexicographical order starting from the position that is specified by this parameter.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs, which are instance names.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the resource. valid value:
	//
	// instance: instance
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*ListTagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v []*string) *ListTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v []*ListTagResourcesRequestTags) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

type ListTagResourcesRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// Owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Tester
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTags) SetKey(v string) *ListTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTags) SetValue(v string) *ListTagResourcesRequestTags {
	s.Value = &v
	return s
}

type ListTagResourcesShrinkRequest struct {
	// The maximum number of tagged resources that you want to return. Valid values: 0 to 200. If you do not specify this parameter or set the parameter to 0, the default value of 100 is used.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. Tagged resources are returned in lexicographical order starting from the position that is specified by this parameter.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs, which are instance names.
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource. valid value:
	//
	// instance: instance
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) SetMaxResults(v int32) *ListTagResourcesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdsShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagsShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The maximum number of tagged resources that are returned for the query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tags.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 31D8120C-AC52-5CA9-BE4A-E4C6316E19AD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetMaxResults(v int32) *ListTagResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The resource ID, which is the instance name.
	//
	// example:
	//
	// instance-test
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ALIYUN::OTS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// Owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Tester
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The resource IDs, which are instance names.
	//
	// This parameter is required.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the resource. valid value:
	//
	// instance: instance
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	//
	// This parameter is required.
	Tags []*TagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceIds(v []*string) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest {
	s.Tags = v
	return s
}

type TagResourcesRequestTags struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// created
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-0007z8j1omiabo5i872l
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTags) SetKey(v string) *TagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTags) SetValue(v string) *TagResourcesRequestTags {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 44FDC379-4443-560E-B652-9F7476D8854F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resources. Default value: false.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The resource IDs, which are instance names.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the resource. valid value:
	//
	// instance: instance
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v []*string) *UntagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v []*string) *UntagResourcesRequest {
	s.TagKeys = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 26B43C99-F5C9-5341-B462-37427C5BF201
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	// The alias of the instance.
	//
	// example:
	//
	// instance-test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// the test instance
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the instance whose information you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// (Deprecated) The network type of the instance. Valid values: NORMAL and VPC_CONSOLE. Default value: NORMAL.
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The new sources of the network from which access is allowed. By default, all sources of networks are allowed. Valid value:
	//
	// TRUST_PROXY: the console
	NetworkSourceACL []*string `json:"NetworkSourceACL,omitempty" xml:"NetworkSourceACL,omitempty" type:"Repeated"`
	// The new types of the network from which access is allowed. By default, all types of networks are allowed. Valid values:
	//
	// 	- INTERNET: the Internet
	//
	// 	- VPC: VPCs
	//
	// 	- CLASSIC: the classic network
	NetworkTypeACL []*string `json:"NetworkTypeACL,omitempty" xml:"NetworkTypeACL,omitempty" type:"Repeated"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetAliasName(v string) *UpdateInstanceRequest {
	s.AliasName = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceDescription(v string) *UpdateInstanceRequest {
	s.InstanceDescription = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetwork(v string) *UpdateInstanceRequest {
	s.Network = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetworkSourceACL(v []*string) *UpdateInstanceRequest {
	s.NetworkSourceACL = v
	return s
}

func (s *UpdateInstanceRequest) SetNetworkTypeACL(v []*string) *UpdateInstanceRequest {
	s.NetworkTypeACL = v
	return s
}

type UpdateInstanceResponseBody struct {
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 3104C83E-6E82-57FB-BB88-8C64CCFDEF89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceElasticVCUUpperLimitRequest struct {
	// The upper limit for the VCUs of the instance.
	//
	// >  Valid values of the upper limit for the VCUs of an instance: **Number of reserved VCUs+0.1 to 2000**. You can upgrade or downgrade configurations to modify the number of reserved VCUs by increments or decrements of 1. You can dynamically modify the upper limit for the VCUs of an instance by increments or decrements of 0.1
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	ElasticVCUUpperLimit *float32 `json:"ElasticVCUUpperLimit,omitempty" xml:"ElasticVCUUpperLimit,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// workshop-bj-ots1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateInstanceElasticVCUUpperLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceElasticVCUUpperLimitRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceElasticVCUUpperLimitRequest) SetElasticVCUUpperLimit(v float32) *UpdateInstanceElasticVCUUpperLimitRequest {
	s.ElasticVCUUpperLimit = &v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitRequest) SetInstanceName(v string) *UpdateInstanceElasticVCUUpperLimitRequest {
	s.InstanceName = &v
	return s
}

type UpdateInstanceElasticVCUUpperLimitResponseBody struct {
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// B37BBA04-D827-55C8-B901-5264B904E8C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateInstanceElasticVCUUpperLimitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceElasticVCUUpperLimitResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceElasticVCUUpperLimitResponseBody) SetRequestId(v string) *UpdateInstanceElasticVCUUpperLimitResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceElasticVCUUpperLimitResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceElasticVCUUpperLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceElasticVCUUpperLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceElasticVCUUpperLimitResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) SetHeaders(v map[string]*string) *UpdateInstanceElasticVCUUpperLimitResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) SetStatusCode(v int32) *UpdateInstanceElasticVCUUpperLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) SetBody(v *UpdateInstanceElasticVCUUpperLimitResponseBody) *UpdateInstanceElasticVCUUpperLimitResponse {
	s.Body = v
	return s
}

type UpdateInstancePolicyRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-test-12345
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance policy in the JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "Version": "1",
	//
	//     "Statement": [
	//
	//         {
	//
	//             "Action": [
	//
	//                 "ots:*"
	//
	//             ],
	//
	//             "Resource": [
	//
	//                 "acs:ots:*:13791xxxxxxxxxxx:instance/myinstance*"
	//
	//             ],
	//
	//             "Principal": [
	//
	//                 "*"
	//
	//             ],
	//
	//             "Effect": "Allow",
	//
	//             "Condition": {
	//
	//                 "StringEquals": {
	//
	//                     "ots:TLSVersion": [
	//
	//                         "1.2"
	//
	//                     ]
	//
	//                 },
	//
	//                 "IpAddress": {
	//
	//                     "acs:SourceIp": [
	//
	//                         "192.168.0.1",
	//
	//                         "172.16.0.1"
	//
	//                     ]
	//
	//                 }
	//
	//             }
	//
	//         }
	//
	//     ]
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The version of the instance policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s UpdateInstancePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstancePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstancePolicyRequest) SetInstanceName(v string) *UpdateInstancePolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstancePolicyRequest) SetPolicy(v string) *UpdateInstancePolicyRequest {
	s.Policy = &v
	return s
}

func (s *UpdateInstancePolicyRequest) SetPolicyVersion(v int64) *UpdateInstancePolicyRequest {
	s.PolicyVersion = &v
	return s
}

type UpdateInstancePolicyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 31D8120C-AC52-5CA9-BE4A-E4C6316E19AD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateInstancePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstancePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstancePolicyResponseBody) SetCode(v string) *UpdateInstancePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstancePolicyResponseBody) SetMessage(v string) *UpdateInstancePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstancePolicyResponseBody) SetRequestId(v string) *UpdateInstancePolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstancePolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstancePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstancePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstancePolicyResponse) SetHeaders(v map[string]*string) *UpdateInstancePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstancePolicyResponse) SetStatusCode(v int32) *UpdateInstancePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstancePolicyResponse) SetBody(v *UpdateInstancePolicyResponseBody) *UpdateInstancePolicyResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("tablestore"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which an instance belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/changeresourcegroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the resource group to which an instance belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the validity of a Resource Access Management (RAM) policy for an instance.
//
// @param request - CheckInstancePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstancePolicyResponse
func (client *Client) CheckInstancePolicyWithOptions(request *CheckInstancePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckInstancePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		body["Policy"] = request.Policy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckInstancePolicy"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/checkinstancepolicy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckInstancePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckInstancePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Checks the validity of a Resource Access Management (RAM) policy for an instance.
//
// @param request - CheckInstancePolicyRequest
//
// @return CheckInstancePolicyResponse
func (client *Client) CheckInstancePolicy(request *CheckInstancePolicyRequest) (_result *CheckInstancePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInstancePolicyResponse{}
	_body, _err := client.CheckInstancePolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
// Description:
//
//   **Before you call this operation, you must understand the billing and pricing of Tablestore. For more information, see [Billing overview](https://help.aliyun.com/document_detail/27291.html).**
//
// 	- Each Alibaba Cloud account can create up to 10 instances. The name of an instance must be unique within the region in which the instance resides.
//
// 	- After you create an instance, you cannot change the type of the instance. Proceed with caution.
//
// @param request - CreateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		body["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.DisableReplication)) {
		body["DisableReplication"] = request.DisableReplication
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceDescription)) {
		body["InstanceDescription"] = request.InstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Network)) {
		body["Network"] = request.Network
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkSourceACL)) {
		body["NetworkSourceACL"] = request.NetworkSourceACL
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkTypeACL)) {
		body["NetworkTypeACL"] = request.NetworkTypeACL
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		body["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/createinstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an instance.
//
// Description:
//
//   **Before you call this operation, you must understand the billing and pricing of Tablestore. For more information, see [Billing overview](https://help.aliyun.com/document_detail/27291.html).**
//
// 	- Each Alibaba Cloud account can create up to 10 instances. The name of an instance must be unique within the region in which the instance resides.
//
// 	- After you create an instance, you cannot change the type of the instance. Proceed with caution.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an instance.
//
// Description:
//
//   Before you delete an instance, make sure that all data tables and time series tables in the instance are deleted and virtual private clouds (VPCs) are unbound from the instance.
//
// 	- To prevent conflicts, do not create an instance that has the same name as the instance that is being deleted.
//
// 	- After an instance is deleted, the instance becomes unavailable and the tables, table data, and related indexes in the instance cannot be recovered. Proceed with caution.
//
// @param request - DeleteInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/deleteinstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an instance.
//
// Description:
//
//   Before you delete an instance, make sure that all data tables and time series tables in the instance are deleted and virtual private clouds (VPCs) are unbound from the instance.
//
// 	- To prevent conflicts, do not create an instance that has the same name as the instance that is being deleted.
//
// 	- After an instance is deleted, the instance becomes unavailable and the tables, table data, and related indexes in the instance cannot be recovered. Proceed with caution.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) policy of an instance.
//
// Description:
//
//   You cannot recover a deleted instance policy. Proceed with caution.
//
// 	- After you delete an instance policy, the access control that is specified by the instance policy becomes invalid. Make sure that your instance is in a secure environment.
//
// @param request - DeleteInstancePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstancePolicyResponse
func (client *Client) DeleteInstancePolicyWithOptions(request *DeleteInstancePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstancePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyVersion)) {
		body["PolicyVersion"] = request.PolicyVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstancePolicy"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/deleteinstancepolicy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteInstancePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteInstancePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a Resource Access Management (RAM) policy of an instance.
//
// Description:
//
//   You cannot recover a deleted instance policy. Proceed with caution.
//
// 	- After you delete an instance policy, the access control that is specified by the instance policy becomes invalid. Make sure that your instance is in a secure environment.
//
// @param request - DeleteInstancePolicyRequest
//
// @return DeleteInstancePolicyResponse
func (client *Client) DeleteInstancePolicy(request *DeleteInstancePolicyRequest) (_result *DeleteInstancePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstancePolicyResponse{}
	_body, _err := client.DeleteInstancePolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries supported regions.
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/region/DescribeRegions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries supported regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instance information.
//
// @param request - GetInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/getinstance"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries instance information.
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instances.
//
// @param tmpReq - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceNameList)) {
		request.InstanceNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceNameList, tea.String("InstanceNameList"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNameListShrink)) {
		query["InstanceNameList"] = request.InstanceNameListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/listinstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries instances.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tagged resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdsShrink)) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/listtagresources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries tagged resources.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to instances.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/tagresources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Adds tags to instances.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// Description:
//
// Removing tags from resources helps simplify resource management, optimize system performance, and mitigate potential security vulnerabilities. After a tag is removed from a resource, the system automatically deletes the tag if the tag is not added to other resources.
//
// @param request - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		body["TagKeys"] = request.TagKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/untagresources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Removes tags from resources.
//
// Description:
//
// Removing tags from resources helps simplify resource management, optimize system performance, and mitigate potential security vulnerabilities. After a tag is removed from a resource, the system automatically deletes the tag if the tag is not added to other resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates instance information.
//
// @param request - UpdateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceDescription)) {
		body["InstanceDescription"] = request.InstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Network)) {
		body["Network"] = request.Network
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkSourceACL)) {
		body["NetworkSourceACL"] = request.NetworkSourceACL
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkTypeACL)) {
		body["NetworkTypeACL"] = request.NetworkTypeACL
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/updateinstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates instance information.
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the upper limit for the VCUs of an instance in VCU mode (formerly reserved mode).
//
// Description:
//
//   **Before you call this operation, you must understand the billing and pricing of Tablestore. For more information, see [Billing overview](https://help.aliyun.com/document_detail/27291.html).**
//
// 	- After you enable scalability for an instance, the default upper limit for the VCUs of the instance is the sum of the scalability and the reserved VCUs.
//
// 	- To use more computing resources when your business grows, you can modify the upper limit for the VCUs of your instance. The new upper limit for the VCUs of your instance immediately takes effect.
//
// @param request - UpdateInstanceElasticVCUUpperLimitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceElasticVCUUpperLimitResponse
func (client *Client) UpdateInstanceElasticVCUUpperLimitWithOptions(request *UpdateInstanceElasticVCUUpperLimitRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceElasticVCUUpperLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticVCUUpperLimit)) {
		body["ElasticVCUUpperLimit"] = request.ElasticVCUUpperLimit
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceElasticVCUUpperLimit"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/updateinstanceelasticvcuupperlimit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceElasticVCUUpperLimitResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceElasticVCUUpperLimitResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the upper limit for the VCUs of an instance in VCU mode (formerly reserved mode).
//
// Description:
//
//   **Before you call this operation, you must understand the billing and pricing of Tablestore. For more information, see [Billing overview](https://help.aliyun.com/document_detail/27291.html).**
//
// 	- After you enable scalability for an instance, the default upper limit for the VCUs of the instance is the sum of the scalability and the reserved VCUs.
//
// 	- To use more computing resources when your business grows, you can modify the upper limit for the VCUs of your instance. The new upper limit for the VCUs of your instance immediately takes effect.
//
// @param request - UpdateInstanceElasticVCUUpperLimitRequest
//
// @return UpdateInstanceElasticVCUUpperLimitResponse
func (client *Client) UpdateInstanceElasticVCUUpperLimit(request *UpdateInstanceElasticVCUUpperLimitRequest) (_result *UpdateInstanceElasticVCUUpperLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceElasticVCUUpperLimitResponse{}
	_body, _err := client.UpdateInstanceElasticVCUUpperLimitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a Resource Access Management (RAM) policy for an instance.
//
// @param request - UpdateInstancePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstancePolicyResponse
func (client *Client) UpdateInstancePolicyWithOptions(request *UpdateInstancePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstancePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		body["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyVersion)) {
		body["PolicyVersion"] = request.PolicyVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstancePolicy"),
		Version:     tea.String("2020-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/openapi/updateinstancepolicy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstancePolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstancePolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies a Resource Access Management (RAM) policy for an instance.
//
// @param request - UpdateInstancePolicyRequest
//
// @return UpdateInstancePolicyResponse
func (client *Client) UpdateInstancePolicy(request *UpdateInstancePolicyRequest) (_result *UpdateInstancePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstancePolicyResponse{}
	_body, _err := client.UpdateInstancePolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
