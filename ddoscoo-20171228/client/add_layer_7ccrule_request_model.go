// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLayer7CCRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAct(v string) *AddLayer7CCRuleRequest
	GetAct() *string
	SetCount(v int32) *AddLayer7CCRuleRequest
	GetCount() *int32
	SetDomain(v string) *AddLayer7CCRuleRequest
	GetDomain() *string
	SetInterval(v int32) *AddLayer7CCRuleRequest
	GetInterval() *int32
	SetMode(v string) *AddLayer7CCRuleRequest
	GetMode() *string
	SetName(v string) *AddLayer7CCRuleRequest
	GetName() *string
	SetResourceGroupId(v string) *AddLayer7CCRuleRequest
	GetResourceGroupId() *string
	SetTtl(v int32) *AddLayer7CCRuleRequest
	GetTtl() *int32
	SetUri(v string) *AddLayer7CCRuleRequest
	GetUri() *string
}

type AddLayer7CCRuleRequest struct {
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

func (s AddLayer7CCRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *AddLayer7CCRuleRequest) GetAct() *string {
	return s.Act
}

func (s *AddLayer7CCRuleRequest) GetCount() *int32 {
	return s.Count
}

func (s *AddLayer7CCRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddLayer7CCRuleRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *AddLayer7CCRuleRequest) GetMode() *string {
	return s.Mode
}

func (s *AddLayer7CCRuleRequest) GetName() *string {
	return s.Name
}

func (s *AddLayer7CCRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddLayer7CCRuleRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *AddLayer7CCRuleRequest) GetUri() *string {
	return s.Uri
}

func (s *AddLayer7CCRuleRequest) SetAct(v string) *AddLayer7CCRuleRequest {
	s.Act = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetCount(v int32) *AddLayer7CCRuleRequest {
	s.Count = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetDomain(v string) *AddLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetInterval(v int32) *AddLayer7CCRuleRequest {
	s.Interval = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetMode(v string) *AddLayer7CCRuleRequest {
	s.Mode = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetName(v string) *AddLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetResourceGroupId(v string) *AddLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetTtl(v int32) *AddLayer7CCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetUri(v string) *AddLayer7CCRuleRequest {
	s.Uri = &v
	return s
}

func (s *AddLayer7CCRuleRequest) Validate() error {
	return dara.Validate(s)
}
