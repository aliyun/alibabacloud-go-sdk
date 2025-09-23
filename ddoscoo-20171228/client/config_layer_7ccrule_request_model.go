// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7CCRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAct(v string) *ConfigLayer7CCRuleRequest
	GetAct() *string
	SetCount(v int32) *ConfigLayer7CCRuleRequest
	GetCount() *int32
	SetDomain(v string) *ConfigLayer7CCRuleRequest
	GetDomain() *string
	SetInterval(v int32) *ConfigLayer7CCRuleRequest
	GetInterval() *int32
	SetMode(v string) *ConfigLayer7CCRuleRequest
	GetMode() *string
	SetName(v string) *ConfigLayer7CCRuleRequest
	GetName() *string
	SetResourceGroupId(v string) *ConfigLayer7CCRuleRequest
	GetResourceGroupId() *string
	SetTtl(v int32) *ConfigLayer7CCRuleRequest
	GetTtl() *int32
	SetUri(v string) *ConfigLayer7CCRuleRequest
	GetUri() *string
}

type ConfigLayer7CCRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// close
	Act *string `json:"Act,omitempty" xml:"Act,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// match
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testCcRule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /a/b/c
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ConfigLayer7CCRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCRuleRequest) GetAct() *string {
	return s.Act
}

func (s *ConfigLayer7CCRuleRequest) GetCount() *int32 {
	return s.Count
}

func (s *ConfigLayer7CCRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigLayer7CCRuleRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *ConfigLayer7CCRuleRequest) GetMode() *string {
	return s.Mode
}

func (s *ConfigLayer7CCRuleRequest) GetName() *string {
	return s.Name
}

func (s *ConfigLayer7CCRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigLayer7CCRuleRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *ConfigLayer7CCRuleRequest) GetUri() *string {
	return s.Uri
}

func (s *ConfigLayer7CCRuleRequest) SetAct(v string) *ConfigLayer7CCRuleRequest {
	s.Act = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetCount(v int32) *ConfigLayer7CCRuleRequest {
	s.Count = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetDomain(v string) *ConfigLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetInterval(v int32) *ConfigLayer7CCRuleRequest {
	s.Interval = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetMode(v string) *ConfigLayer7CCRuleRequest {
	s.Mode = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetName(v string) *ConfigLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetResourceGroupId(v string) *ConfigLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetTtl(v int32) *ConfigLayer7CCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetUri(v string) *ConfigLayer7CCRuleRequest {
	s.Uri = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) Validate() error {
	return dara.Validate(s)
}
