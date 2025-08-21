// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCCRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAct(v string) *ModifyWebCCRuleRequest
	GetAct() *string
	SetCount(v int32) *ModifyWebCCRuleRequest
	GetCount() *int32
	SetDomain(v string) *ModifyWebCCRuleRequest
	GetDomain() *string
	SetInterval(v int32) *ModifyWebCCRuleRequest
	GetInterval() *int32
	SetMode(v string) *ModifyWebCCRuleRequest
	GetMode() *string
	SetName(v string) *ModifyWebCCRuleRequest
	GetName() *string
	SetResourceGroupId(v string) *ModifyWebCCRuleRequest
	GetResourceGroupId() *string
	SetTtl(v int32) *ModifyWebCCRuleRequest
	GetTtl() *int32
	SetUri(v string) *ModifyWebCCRuleRequest
	GetUri() *string
}

type ModifyWebCCRuleRequest struct {
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
	// 3
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
	// 30
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prefix
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testrule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The blocking duration. Valid values: **60*	- to **86400**. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The check path.
	//
	// >  You cannot modify the Uniform Resource Identifier (URI). The domain name of the website, the check path, and the rule name uniquely identify a rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// /abc
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ModifyWebCCRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCCRuleRequest) GetAct() *string {
	return s.Act
}

func (s *ModifyWebCCRuleRequest) GetCount() *int32 {
	return s.Count
}

func (s *ModifyWebCCRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebCCRuleRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyWebCCRuleRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyWebCCRuleRequest) GetName() *string {
	return s.Name
}

func (s *ModifyWebCCRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebCCRuleRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *ModifyWebCCRuleRequest) GetUri() *string {
	return s.Uri
}

func (s *ModifyWebCCRuleRequest) SetAct(v string) *ModifyWebCCRuleRequest {
	s.Act = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetCount(v int32) *ModifyWebCCRuleRequest {
	s.Count = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetDomain(v string) *ModifyWebCCRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetInterval(v int32) *ModifyWebCCRuleRequest {
	s.Interval = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetMode(v string) *ModifyWebCCRuleRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetName(v string) *ModifyWebCCRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetResourceGroupId(v string) *ModifyWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetTtl(v int32) *ModifyWebCCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetUri(v string) *ModifyWebCCRuleRequest {
	s.Uri = &v
	return s
}

func (s *ModifyWebCCRuleRequest) Validate() error {
	return dara.Validate(s)
}
