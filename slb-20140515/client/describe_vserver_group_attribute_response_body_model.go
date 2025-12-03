// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVServerGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v *DescribeVServerGroupAttributeResponseBodyBackendServers) *DescribeVServerGroupAttributeResponseBody
	GetBackendServers() *DescribeVServerGroupAttributeResponseBodyBackendServers
	SetCreateTime(v string) *DescribeVServerGroupAttributeResponseBody
	GetCreateTime() *string
	SetLoadBalancerId(v string) *DescribeVServerGroupAttributeResponseBody
	GetLoadBalancerId() *string
	SetRequestId(v string) *DescribeVServerGroupAttributeResponseBody
	GetRequestId() *string
	SetTags(v *DescribeVServerGroupAttributeResponseBodyTags) *DescribeVServerGroupAttributeResponseBody
	GetTags() *DescribeVServerGroupAttributeResponseBodyTags
	SetVServerGroupId(v string) *DescribeVServerGroupAttributeResponseBody
	GetVServerGroupId() *string
	SetVServerGroupName(v string) *DescribeVServerGroupAttributeResponseBody
	GetVServerGroupName() *string
}

type DescribeVServerGroupAttributeResponseBody struct {
	// The backend servers.
	BackendServers *DescribeVServerGroupAttributeResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The time when the CLB instance was created. The time follows the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2022-08-31T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-jfakd****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags of the backend server.
	Tags *DescribeVServerGroupAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-cige6****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The name of the vServer group.
	//
	// example:
	//
	// Group1
	VServerGroupName *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s DescribeVServerGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponseBody) GetBackendServers() *DescribeVServerGroupAttributeResponseBodyBackendServers {
	return s.BackendServers
}

func (s *DescribeVServerGroupAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeVServerGroupAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeVServerGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVServerGroupAttributeResponseBody) GetTags() *DescribeVServerGroupAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeVServerGroupAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeVServerGroupAttributeResponseBody) GetVServerGroupName() *string {
	return s.VServerGroupName
}

func (s *DescribeVServerGroupAttributeResponseBody) SetBackendServers(v *DescribeVServerGroupAttributeResponseBodyBackendServers) *DescribeVServerGroupAttributeResponseBody {
	s.BackendServers = v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetCreateTime(v string) *DescribeVServerGroupAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetLoadBalancerId(v string) *DescribeVServerGroupAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetRequestId(v string) *DescribeVServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetTags(v *DescribeVServerGroupAttributeResponseBodyTags) *DescribeVServerGroupAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetVServerGroupId(v string) *DescribeVServerGroupAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetVServerGroupName(v string) *DescribeVServerGroupAttributeResponseBody {
	s.VServerGroupName = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
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

type DescribeVServerGroupAttributeResponseBodyBackendServers struct {
	BackendServer []*DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupAttributeResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServers) GetBackendServer() []*DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServers) SetBackendServer(v []*DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) *DescribeVServerGroupAttributeResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServers) Validate() error {
	if s.BackendServer != nil {
		for _, item := range s.BackendServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer struct {
	// The description of the server group.
	//
	// >  This parameter is not returned if the Description parameter is not specified in the request.
	//
	// example:
	//
	// Server Group Description
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
	// vm-233
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.XX.XX.11
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of backend server. Valid values:
	//
	// 	- **ecs**: ECS instance
	//
	// 	- **eni**: ENI
	//
	// 	- **eci**: elastic container instance
	//
	// example:
	//
	// ecs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) GetDescription() *string {
	return s.Description
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) GetPort() *int32 {
	return s.Port
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) GetServerIp() *string {
	return s.ServerIp
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) GetType() *string {
	return s.Type
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetDescription(v string) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetPort(v int32) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetServerId(v string) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetServerIp(v string) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.ServerIp = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetType(v string) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetWeight(v int32) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}

type DescribeVServerGroupAttributeResponseBodyTags struct {
	Tag []*DescribeVServerGroupAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponseBodyTags) GetTag() []*DescribeVServerGroupAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeVServerGroupAttributeResponseBodyTags) SetTag(v []*DescribeVServerGroupAttributeResponseBodyTagsTag) *DescribeVServerGroupAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyTags) Validate() error {
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

type DescribeVServerGroupAttributeResponseBodyTagsTag struct {
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

func (s DescribeVServerGroupAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeVServerGroupAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeVServerGroupAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeVServerGroupAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeVServerGroupAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
