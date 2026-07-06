// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageDescription(v string) *CreateAgentStorageRequest
	GetAgentStorageDescription() *string
	SetAgentStorageName(v string) *CreateAgentStorageRequest
	GetAgentStorageName() *string
	SetNetwork(v string) *CreateAgentStorageRequest
	GetNetwork() *string
	SetNetworkSourceACL(v []*string) *CreateAgentStorageRequest
	GetNetworkSourceACL() []*string
	SetNetworkTypeACL(v []*string) *CreateAgentStorageRequest
	GetNetworkTypeACL() []*string
	SetPolicy(v string) *CreateAgentStorageRequest
	GetPolicy() *string
	SetResourceGroupId(v string) *CreateAgentStorageRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateAgentStorageRequestTags) *CreateAgentStorageRequest
	GetTags() []*CreateAgentStorageRequestTags
}

type CreateAgentStorageRequest struct {
	// The description of the agent storage. The description must be 3 to 256 characters in length.
	//
	// example:
	//
	// description for agent storage
	AgentStorageDescription *string `json:"AgentStorageDescription,omitempty" xml:"AgentStorageDescription,omitempty"`
	// The name of the agent storage. The naming conventions are as follows:
	//
	// - The name can contain only letters, digits, and hyphens (-).
	//
	// - The name must start with a letter.
	//
	// - The name cannot end with a hyphen (-).
	//
	// - The name is case-insensitive.
	//
	// - The name must be 3 to 16 characters in length.
	//
	// - The name cannot contain the following words: ali, ay, ots, taobao, or admin.
	//
	// This parameter is required.
	//
	// example:
	//
	// first-agent
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// (Deprecated) The network type of the agent storage. Valid values: NORMAL and VPC_CONSOLE. Default value: NORMAL.
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The list of allowed network sources for the agent storage. All sources are allowed by default. Valid values:
	//
	// - TRUST_PROXY: console.
	NetworkSourceACL []*string `json:"NetworkSourceACL,omitempty" xml:"NetworkSourceACL,omitempty" type:"Repeated"`
	// The list of allowed network types for the agent storage. All types are allowed by default. Valid values:
	//
	// - CLASSIC: classic network.
	//
	// - INTERNET: public network.
	//
	// - VPC: VPC network.
	NetworkTypeACL []*string `json:"NetworkTypeACL,omitempty" xml:"NetworkTypeACL,omitempty" type:"Repeated"`
	// The access control policy for the agent storage in JSON format. For the syntax, see https://www.alibabacloud.com/help/en/ram/user-guide/policy-structure-and-syntax.
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
	//                 "acs:ots:*:13791xxxxxxxxxxx:agentstorage/myagentstorage*"
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
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxh4em5jnbcd
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tags []*CreateAgentStorageRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateAgentStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentStorageRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentStorageRequest) GetAgentStorageDescription() *string {
	return s.AgentStorageDescription
}

func (s *CreateAgentStorageRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *CreateAgentStorageRequest) GetNetwork() *string {
	return s.Network
}

func (s *CreateAgentStorageRequest) GetNetworkSourceACL() []*string {
	return s.NetworkSourceACL
}

func (s *CreateAgentStorageRequest) GetNetworkTypeACL() []*string {
	return s.NetworkTypeACL
}

func (s *CreateAgentStorageRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CreateAgentStorageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAgentStorageRequest) GetTags() []*CreateAgentStorageRequestTags {
	return s.Tags
}

func (s *CreateAgentStorageRequest) SetAgentStorageDescription(v string) *CreateAgentStorageRequest {
	s.AgentStorageDescription = &v
	return s
}

func (s *CreateAgentStorageRequest) SetAgentStorageName(v string) *CreateAgentStorageRequest {
	s.AgentStorageName = &v
	return s
}

func (s *CreateAgentStorageRequest) SetNetwork(v string) *CreateAgentStorageRequest {
	s.Network = &v
	return s
}

func (s *CreateAgentStorageRequest) SetNetworkSourceACL(v []*string) *CreateAgentStorageRequest {
	s.NetworkSourceACL = v
	return s
}

func (s *CreateAgentStorageRequest) SetNetworkTypeACL(v []*string) *CreateAgentStorageRequest {
	s.NetworkTypeACL = v
	return s
}

func (s *CreateAgentStorageRequest) SetPolicy(v string) *CreateAgentStorageRequest {
	s.Policy = &v
	return s
}

func (s *CreateAgentStorageRequest) SetResourceGroupId(v string) *CreateAgentStorageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAgentStorageRequest) SetTags(v []*CreateAgentStorageRequestTags) *CreateAgentStorageRequest {
	s.Tags = v
	return s
}

func (s *CreateAgentStorageRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAgentStorageRequestTags struct {
	// The key of the tag. The key can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// created
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The value can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0woauavextilfqr61
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAgentStorageRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentStorageRequestTags) GoString() string {
	return s.String()
}

func (s *CreateAgentStorageRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateAgentStorageRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateAgentStorageRequestTags) SetKey(v string) *CreateAgentStorageRequestTags {
	s.Key = &v
	return s
}

func (s *CreateAgentStorageRequestTags) SetValue(v string) *CreateAgentStorageRequestTags {
	s.Value = &v
	return s
}

func (s *CreateAgentStorageRequestTags) Validate() error {
	return dara.Validate(s)
}
