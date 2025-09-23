// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayer7RuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CreateLayer7RuleRequest
	GetDomain() *string
	SetInstanceIds(v []*string) *CreateLayer7RuleRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *CreateLayer7RuleRequest
	GetResourceGroupId() *string
	SetRsType(v int32) *CreateLayer7RuleRequest
	GetRsType() *int32
	SetRules(v string) *CreateLayer7RuleRequest
	GetRules() *string
}

type CreateLayer7RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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
	// This parameter is required.
	//
	// example:
	//
	// [{"ProxyRules":[{"ProxyPort":443,"RealServers":["1.1.1.1:443"]}],"ProxyType":"https"},{"ProxyRules":[{"ProxyPort":80,"RealServers":["1.1.1.1:80"]}],"ProxyType":"http"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreateLayer7RuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateLayer7RuleRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateLayer7RuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLayer7RuleRequest) GetRsType() *int32 {
	return s.RsType
}

func (s *CreateLayer7RuleRequest) GetRules() *string {
	return s.Rules
}

func (s *CreateLayer7RuleRequest) SetDomain(v string) *CreateLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetInstanceIds(v []*string) *CreateLayer7RuleRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateLayer7RuleRequest) SetResourceGroupId(v string) *CreateLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetRsType(v int32) *CreateLayer7RuleRequest {
	s.RsType = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetRules(v string) *CreateLayer7RuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateLayer7RuleRequest) Validate() error {
	return dara.Validate(s)
}
