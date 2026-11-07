// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTagPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ApplyTagPoliciesRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *ApplyTagPoliciesRequest
	GetAppId() *string
	SetAppName(v string) *ApplyTagPoliciesRequest
	GetAppName() *string
	SetEnable(v bool) *ApplyTagPoliciesRequest
	GetEnable() *bool
	SetNamespace(v string) *ApplyTagPoliciesRequest
	GetNamespace() *string
	SetNamespaceId(v string) *ApplyTagPoliciesRequest
	GetNamespaceId() *string
	SetRegion(v string) *ApplyTagPoliciesRequest
	GetRegion() *string
	SetRules(v map[string]*RulesValue) *ApplyTagPoliciesRequest
	GetRules() map[string]*RulesValue
}

type ApplyTagPoliciesRequest struct {
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Deprecated
	//
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// h5pohqu3gd@xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether the rule is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The MSE namespace to which the application belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The namespace ID. This parameter is optional.
	//
	// example:
	//
	// 12233****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The routing rules.
	Rules map[string]*RulesValue `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ApplyTagPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyTagPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ApplyTagPoliciesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ApplyTagPoliciesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ApplyTagPoliciesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ApplyTagPoliciesRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ApplyTagPoliciesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ApplyTagPoliciesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ApplyTagPoliciesRequest) GetRegion() *string {
	return s.Region
}

func (s *ApplyTagPoliciesRequest) GetRules() map[string]*RulesValue {
	return s.Rules
}

func (s *ApplyTagPoliciesRequest) SetAcceptLanguage(v string) *ApplyTagPoliciesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ApplyTagPoliciesRequest) SetAppId(v string) *ApplyTagPoliciesRequest {
	s.AppId = &v
	return s
}

func (s *ApplyTagPoliciesRequest) SetAppName(v string) *ApplyTagPoliciesRequest {
	s.AppName = &v
	return s
}

func (s *ApplyTagPoliciesRequest) SetEnable(v bool) *ApplyTagPoliciesRequest {
	s.Enable = &v
	return s
}

func (s *ApplyTagPoliciesRequest) SetNamespace(v string) *ApplyTagPoliciesRequest {
	s.Namespace = &v
	return s
}

func (s *ApplyTagPoliciesRequest) SetNamespaceId(v string) *ApplyTagPoliciesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ApplyTagPoliciesRequest) SetRegion(v string) *ApplyTagPoliciesRequest {
	s.Region = &v
	return s
}

func (s *ApplyTagPoliciesRequest) SetRules(v map[string]*RulesValue) *ApplyTagPoliciesRequest {
	s.Rules = v
	return s
}

func (s *ApplyTagPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
