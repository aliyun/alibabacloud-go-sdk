// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *CreateInstanceRequest
	GetClusterType() *string
	SetDisableReplication(v bool) *CreateInstanceRequest
	GetDisableReplication() *bool
	SetInstanceDescription(v string) *CreateInstanceRequest
	GetInstanceDescription() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetNetwork(v string) *CreateInstanceRequest
	GetNetwork() *string
	SetNetworkSourceACL(v []*string) *CreateInstanceRequest
	GetNetworkSourceACL() []*string
	SetNetworkTypeACL(v []*string) *CreateInstanceRequest
	GetNetworkTypeACL() []*string
	SetPolicy(v string) *CreateInstanceRequest
	GetPolicy() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest
	GetTags() []*CreateInstanceRequestTags
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
	// Deprecated
	//
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
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateInstanceRequest) GetDisableReplication() *bool {
	return s.DisableReplication
}

func (s *CreateInstanceRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetNetwork() *string {
	return s.Network
}

func (s *CreateInstanceRequest) GetNetworkSourceACL() []*string {
	return s.NetworkSourceACL
}

func (s *CreateInstanceRequest) GetNetworkTypeACL() []*string {
	return s.NetworkTypeACL
}

func (s *CreateInstanceRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetTags() []*CreateInstanceRequestTags {
	return s.Tags
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

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
