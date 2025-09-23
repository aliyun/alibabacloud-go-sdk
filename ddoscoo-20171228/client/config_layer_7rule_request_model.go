// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7RuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ConfigLayer7RuleRequest
	GetDomain() *string
	SetInstanceIds(v []*string) *ConfigLayer7RuleRequest
	GetInstanceIds() []*string
	SetProxyTypeList(v string) *ConfigLayer7RuleRequest
	GetProxyTypeList() *string
	SetProxyTypes(v []*string) *ConfigLayer7RuleRequest
	GetProxyTypes() []*string
	SetRealServers(v []*string) *ConfigLayer7RuleRequest
	GetRealServers() []*string
	SetResourceGroupId(v string) *ConfigLayer7RuleRequest
	GetResourceGroupId() *string
	SetRsType(v int32) *ConfigLayer7RuleRequest
	GetRsType() *int32
}

type ConfigLayer7RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXXX
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// [{"ProxyPorts":[80,8080],"ProxyType":"http"},{"ProxyPorts":[443],"ProxyType":"https"}]rts\\":[443],\\"ProxyType\\":\\"https\\"}]
	ProxyTypeList *string `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty"`
	// example:
	//
	// [{"ProxyPorts":[80,8080],"ProxyType":"http"},{"ProxyPorts":[443],"ProxyType":"https"}]rts\\":[443],\\"ProxyType\\":\\"https\\"}]
	ProxyTypes []*string `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s ConfigLayer7RuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7RuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigLayer7RuleRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ConfigLayer7RuleRequest) GetProxyTypeList() *string {
	return s.ProxyTypeList
}

func (s *ConfigLayer7RuleRequest) GetProxyTypes() []*string {
	return s.ProxyTypes
}

func (s *ConfigLayer7RuleRequest) GetRealServers() []*string {
	return s.RealServers
}

func (s *ConfigLayer7RuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigLayer7RuleRequest) GetRsType() *int32 {
	return s.RsType
}

func (s *ConfigLayer7RuleRequest) SetDomain(v string) *ConfigLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetInstanceIds(v []*string) *ConfigLayer7RuleRequest {
	s.InstanceIds = v
	return s
}

func (s *ConfigLayer7RuleRequest) SetProxyTypeList(v string) *ConfigLayer7RuleRequest {
	s.ProxyTypeList = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetProxyTypes(v []*string) *ConfigLayer7RuleRequest {
	s.ProxyTypes = v
	return s
}

func (s *ConfigLayer7RuleRequest) SetRealServers(v []*string) *ConfigLayer7RuleRequest {
	s.RealServers = v
	return s
}

func (s *ConfigLayer7RuleRequest) SetResourceGroupId(v string) *ConfigLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetRsType(v int32) *ConfigLayer7RuleRequest {
	s.RsType = &v
	return s
}

func (s *ConfigLayer7RuleRequest) Validate() error {
	return dara.Validate(s)
}
