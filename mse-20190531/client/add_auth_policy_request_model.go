// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddAuthPolicyRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *AddAuthPolicyRequest
	GetAppId() *string
	SetAuthRule(v string) *AddAuthPolicyRequest
	GetAuthRule() *string
	SetAuthType(v int32) *AddAuthPolicyRequest
	GetAuthType() *int32
	SetEnable(v string) *AddAuthPolicyRequest
	GetEnable() *string
	SetK8sNamespace(v string) *AddAuthPolicyRequest
	GetK8sNamespace() *string
	SetName(v string) *AddAuthPolicyRequest
	GetName() *string
	SetNamespace(v string) *AddAuthPolicyRequest
	GetNamespace() *string
	SetProtocol(v string) *AddAuthPolicyRequest
	GetProtocol() *string
	SetRegion(v string) *AddAuthPolicyRequest
	GetRegion() *string
	SetSource(v string) *AddAuthPolicyRequest
	GetSource() *string
}

type AddAuthPolicyRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// e9clba2xlc@***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"all":true,"black":false,"appIds":["dez4qt7weh@***"]}]
	AuthRule *string `json:"AuthRule,omitempty" xml:"AuthRule,omitempty"`
	// example:
	//
	// 0
	AuthType *int32 `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// c19c6c500e1ff4d7abc7***
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SPRING_CLOUD
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s AddAuthPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuthPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddAuthPolicyRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddAuthPolicyRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddAuthPolicyRequest) GetAuthRule() *string {
	return s.AuthRule
}

func (s *AddAuthPolicyRequest) GetAuthType() *int32 {
	return s.AuthType
}

func (s *AddAuthPolicyRequest) GetEnable() *string {
	return s.Enable
}

func (s *AddAuthPolicyRequest) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *AddAuthPolicyRequest) GetName() *string {
	return s.Name
}

func (s *AddAuthPolicyRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *AddAuthPolicyRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *AddAuthPolicyRequest) GetRegion() *string {
	return s.Region
}

func (s *AddAuthPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *AddAuthPolicyRequest) SetAcceptLanguage(v string) *AddAuthPolicyRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddAuthPolicyRequest) SetAppId(v string) *AddAuthPolicyRequest {
	s.AppId = &v
	return s
}

func (s *AddAuthPolicyRequest) SetAuthRule(v string) *AddAuthPolicyRequest {
	s.AuthRule = &v
	return s
}

func (s *AddAuthPolicyRequest) SetAuthType(v int32) *AddAuthPolicyRequest {
	s.AuthType = &v
	return s
}

func (s *AddAuthPolicyRequest) SetEnable(v string) *AddAuthPolicyRequest {
	s.Enable = &v
	return s
}

func (s *AddAuthPolicyRequest) SetK8sNamespace(v string) *AddAuthPolicyRequest {
	s.K8sNamespace = &v
	return s
}

func (s *AddAuthPolicyRequest) SetName(v string) *AddAuthPolicyRequest {
	s.Name = &v
	return s
}

func (s *AddAuthPolicyRequest) SetNamespace(v string) *AddAuthPolicyRequest {
	s.Namespace = &v
	return s
}

func (s *AddAuthPolicyRequest) SetProtocol(v string) *AddAuthPolicyRequest {
	s.Protocol = &v
	return s
}

func (s *AddAuthPolicyRequest) SetRegion(v string) *AddAuthPolicyRequest {
	s.Region = &v
	return s
}

func (s *AddAuthPolicyRequest) SetSource(v string) *AddAuthPolicyRequest {
	s.Source = &v
	return s
}

func (s *AddAuthPolicyRequest) Validate() error {
	return dara.Validate(s)
}
