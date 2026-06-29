// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageDescription(v string) *GetAgentStorageResponseBody
	GetAgentStorageDescription() *string
	SetAgentStorageName(v string) *GetAgentStorageResponseBody
	GetAgentStorageName() *string
	SetAgentStorageSpecification(v string) *GetAgentStorageResponseBody
	GetAgentStorageSpecification() *string
	SetAgentStorageStatus(v string) *GetAgentStorageResponseBody
	GetAgentStorageStatus() *string
	SetAliasName(v string) *GetAgentStorageResponseBody
	GetAliasName() *string
	SetCreateTime(v string) *GetAgentStorageResponseBody
	GetCreateTime() *string
	SetNetworkSourceACL(v []*string) *GetAgentStorageResponseBody
	GetNetworkSourceACL() []*string
	SetNetworkTypeACL(v []*string) *GetAgentStorageResponseBody
	GetNetworkTypeACL() []*string
	SetPolicy(v string) *GetAgentStorageResponseBody
	GetPolicy() *string
	SetPolicyVersion(v int64) *GetAgentStorageResponseBody
	GetPolicyVersion() *int64
	SetRegionId(v string) *GetAgentStorageResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetAgentStorageResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetAgentStorageResponseBody
	GetResourceGroupId() *string
	SetTags(v []*GetAgentStorageResponseBodyTags) *GetAgentStorageResponseBody
	GetTags() []*GetAgentStorageResponseBodyTags
	SetUserId(v string) *GetAgentStorageResponseBody
	GetUserId() *string
}

type GetAgentStorageResponseBody struct {
	// agent storage description
	//
	// example:
	//
	// description for agent storage
	AgentStorageDescription *string `json:"AgentStorageDescription,omitempty" xml:"AgentStorageDescription,omitempty"`
	// agent storage name
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The specification of the agent storage.
	//
	// example:
	//
	// SSD
	AgentStorageSpecification *string `json:"AgentStorageSpecification,omitempty" xml:"AgentStorageSpecification,omitempty"`
	// The status of the agent storage. Valid values:
	//
	// - normal: Normal.
	//
	// - forbidden: Disabled.
	//
	// - deleting: Being released.
	//
	// example:
	//
	// normal
	AgentStorageStatus *string `json:"AgentStorageStatus,omitempty" xml:"AgentStorageStatus,omitempty"`
	// The alias of the agent storage.
	//
	// example:
	//
	// agent-test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The time when the agent storage was created.
	//
	// example:
	//
	// 2019-12-23T07:24:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The list of network sources allowed for the agent storage. TRUST_PROXY: console.
	NetworkSourceACL []*string `json:"NetworkSourceACL,omitempty" xml:"NetworkSourceACL,omitempty" type:"Repeated"`
	// The list of network types allowed for the agent storage. CLASSIC: classic network. INTERNET: Internet. VPC: VPC network.
	NetworkTypeACL []*string `json:"NetworkTypeACL,omitempty" xml:"NetworkTypeACL,omitempty" type:"Repeated"`
	// The access control policy of the agent storage.
	//
	// example:
	//
	// drop
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The version of the agent storage policy.
	//
	// example:
	//
	// 1
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The region ID of the agent storage.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3104C83E-6E82-57FB-BB88-8C64CCFDEF89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxh4em5jncda
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the agent storage.
	Tags []*GetAgentStorageResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 16542312566
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAgentStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStorageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentStorageResponseBody) GetAgentStorageDescription() *string {
	return s.AgentStorageDescription
}

func (s *GetAgentStorageResponseBody) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *GetAgentStorageResponseBody) GetAgentStorageSpecification() *string {
	return s.AgentStorageSpecification
}

func (s *GetAgentStorageResponseBody) GetAgentStorageStatus() *string {
	return s.AgentStorageStatus
}

func (s *GetAgentStorageResponseBody) GetAliasName() *string {
	return s.AliasName
}

func (s *GetAgentStorageResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAgentStorageResponseBody) GetNetworkSourceACL() []*string {
	return s.NetworkSourceACL
}

func (s *GetAgentStorageResponseBody) GetNetworkTypeACL() []*string {
	return s.NetworkTypeACL
}

func (s *GetAgentStorageResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GetAgentStorageResponseBody) GetPolicyVersion() *int64 {
	return s.PolicyVersion
}

func (s *GetAgentStorageResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAgentStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentStorageResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetAgentStorageResponseBody) GetTags() []*GetAgentStorageResponseBodyTags {
	return s.Tags
}

func (s *GetAgentStorageResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetAgentStorageResponseBody) SetAgentStorageDescription(v string) *GetAgentStorageResponseBody {
	s.AgentStorageDescription = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetAgentStorageName(v string) *GetAgentStorageResponseBody {
	s.AgentStorageName = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetAgentStorageSpecification(v string) *GetAgentStorageResponseBody {
	s.AgentStorageSpecification = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetAgentStorageStatus(v string) *GetAgentStorageResponseBody {
	s.AgentStorageStatus = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetAliasName(v string) *GetAgentStorageResponseBody {
	s.AliasName = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetCreateTime(v string) *GetAgentStorageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetNetworkSourceACL(v []*string) *GetAgentStorageResponseBody {
	s.NetworkSourceACL = v
	return s
}

func (s *GetAgentStorageResponseBody) SetNetworkTypeACL(v []*string) *GetAgentStorageResponseBody {
	s.NetworkTypeACL = v
	return s
}

func (s *GetAgentStorageResponseBody) SetPolicy(v string) *GetAgentStorageResponseBody {
	s.Policy = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetPolicyVersion(v int64) *GetAgentStorageResponseBody {
	s.PolicyVersion = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetRegionId(v string) *GetAgentStorageResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetRequestId(v string) *GetAgentStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetResourceGroupId(v string) *GetAgentStorageResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetAgentStorageResponseBody) SetTags(v []*GetAgentStorageResponseBodyTags) *GetAgentStorageResponseBody {
	s.Tags = v
	return s
}

func (s *GetAgentStorageResponseBody) SetUserId(v string) *GetAgentStorageResponseBody {
	s.UserId = &v
	return s
}

func (s *GetAgentStorageResponseBody) Validate() error {
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

type GetAgentStorageResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// Owner
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
	// Tester
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAgentStorageResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStorageResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetAgentStorageResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetAgentStorageResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAgentStorageResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAgentStorageResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetAgentStorageResponseBodyTags) SetKey(v string) *GetAgentStorageResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetAgentStorageResponseBodyTags) SetTagKey(v string) *GetAgentStorageResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetAgentStorageResponseBodyTags) SetTagValue(v string) *GetAgentStorageResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetAgentStorageResponseBodyTags) SetValue(v string) *GetAgentStorageResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetAgentStorageResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
