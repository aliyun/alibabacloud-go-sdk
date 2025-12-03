// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMasterSlaveServerGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody
	GetCreateTime() *string
	SetLoadBalancerId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody
	GetLoadBalancerId() *string
	SetMasterSlaveBackendServers(v *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) *DescribeMasterSlaveServerGroupAttributeResponseBody
	GetMasterSlaveBackendServers() *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers
	SetMasterSlaveServerGroupId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody
	GetMasterSlaveServerGroupId() *string
	SetMasterSlaveServerGroupName(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody
	GetMasterSlaveServerGroupName() *string
	SetRequestId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody
	GetRequestId() *string
	SetTags(v *DescribeMasterSlaveServerGroupAttributeResponseBodyTags) *DescribeMasterSlaveServerGroupAttributeResponseBody
	GetTags() *DescribeMasterSlaveServerGroupAttributeResponseBodyTags
}

type DescribeMasterSlaveServerGroupAttributeResponseBody struct {
	// The time when the CLB instance was created. The time follows the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2022-12-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the associated CLB instance.
	//
	// example:
	//
	// lb-14fadafw4343a******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// A list of backend servers in the primary/secondary server group.
	MasterSlaveBackendServers *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers `json:"MasterSlaveBackendServers,omitempty" xml:"MasterSlaveBackendServers,omitempty" type:"Struct"`
	// The ID of the primary/secondary server group.
	//
	// example:
	//
	// rsp-cige6******
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	// The name of the primary/secondary server group.
	//
	// example:
	//
	// Group1
	MasterSlaveServerGroupName *string `json:"MasterSlaveServerGroupName,omitempty" xml:"MasterSlaveServerGroupName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag list.
	Tags *DescribeMasterSlaveServerGroupAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) GetMasterSlaveBackendServers() *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers {
	return s.MasterSlaveBackendServers
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) GetMasterSlaveServerGroupName() *string {
	return s.MasterSlaveServerGroupName
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) GetTags() *DescribeMasterSlaveServerGroupAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetCreateTime(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetLoadBalancerId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetMasterSlaveBackendServers(v *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.MasterSlaveBackendServers = v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetMasterSlaveServerGroupId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetMasterSlaveServerGroupName(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.MasterSlaveServerGroupName = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetRequestId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetTags(v *DescribeMasterSlaveServerGroupAttributeResponseBodyTags) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) Validate() error {
	if s.MasterSlaveBackendServers != nil {
		if err := s.MasterSlaveBackendServers.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers struct {
	MasterSlaveBackendServer []*DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer `json:"MasterSlaveBackendServer,omitempty" xml:"MasterSlaveBackendServer,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) GetMasterSlaveBackendServer() []*DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	return s.MasterSlaveBackendServer
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) SetMasterSlaveBackendServer(v []*DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers {
	s.MasterSlaveBackendServer = v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) Validate() error {
	if s.MasterSlaveBackendServer != nil {
		for _, item := range s.MasterSlaveBackendServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer struct {
	// The description of the primary/secondary server group.
	//
	// example:
	//
	// Primary and secondary server group description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port that is used by the backend server.
	//
	// example:
	//
	// 90
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server.
	//
	// example:
	//
	// eni-hhshhs****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of backend server. Valid values: **Master*	- and **Slave**.
	//
	// example:
	//
	// Slave
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **ecs*	- (default): Elastic Compute Service (ECS) instance
	//
	// 	- **eni**: elastic network interface (ENI)
	//
	// 	- **eci**: elastic container instance
	//
	// example:
	//
	// eni
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetDescription() *string {
	return s.Description
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetServerType() *string {
	return s.ServerType
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetType() *string {
	return s.Type
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetDescription(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Description = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetPort(v int32) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Port = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetServerId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetServerType(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.ServerType = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetType(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Type = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetWeight(v int32) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Weight = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) Validate() error {
	return dara.Validate(s)
}

type DescribeMasterSlaveServerGroupAttributeResponseBodyTags struct {
	Tag []*DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyTags) GetTag() []*DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyTags) SetTag(v []*DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag) *DescribeMasterSlaveServerGroupAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyTags) Validate() error {
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

type DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag struct {
	// The tag key. Valid values of N: **1*	- to **20**. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value. Valid values of N: **1*	- to **20**. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
