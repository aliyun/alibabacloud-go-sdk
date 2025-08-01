// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateAuthPolicyRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *UpdateAuthPolicyRequest
	GetAppId() *string
	SetAuthRule(v string) *UpdateAuthPolicyRequest
	GetAuthRule() *string
	SetEnable(v string) *UpdateAuthPolicyRequest
	GetEnable() *string
	SetId(v string) *UpdateAuthPolicyRequest
	GetId() *string
	SetK8sNamespace(v string) *UpdateAuthPolicyRequest
	GetK8sNamespace() *string
	SetName(v string) *UpdateAuthPolicyRequest
	GetName() *string
	SetProtocol(v string) *UpdateAuthPolicyRequest
	GetProtocol() *string
	SetRegion(v string) *UpdateAuthPolicyRequest
	GetRegion() *string
	SetSource(v string) *UpdateAuthPolicyRequest
	GetSource() *string
}

type UpdateAuthPolicyRequest struct {
	// The language of the response. Valid values: zh-CN and en-US. Default value: zh-CN. The value zh-CN indicates Chinese, and the value en-US indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application ID.
	//
	// example:
	//
	// abcde@12345
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The content of the service authentication rule.
	//
	// example:
	//
	// [{\\"all\\":true,\\"appIds\\":[\\"f6xqzbao96@4adfaf3c92c947a\\"],\\"black\\":false}]
	AuthRule *string `json:"AuthRule,omitempty" xml:"AuthRule,omitempty"`
	// Specifies whether to enable the rule.
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 432
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the ACK cluster namespace.
	//
	// example:
	//
	// c19c6c500e1ff4d7abc7bed9b8236***
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// demo-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **SPRING_CLOUD**
	//
	// 	- **DUBBO**
	//
	// 	- **istio**
	//
	// example:
	//
	// SPRING_CLOUD
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The source for application access.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateAuthPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthPolicyRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateAuthPolicyRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAuthPolicyRequest) GetAuthRule() *string {
	return s.AuthRule
}

func (s *UpdateAuthPolicyRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateAuthPolicyRequest) GetId() *string {
	return s.Id
}

func (s *UpdateAuthPolicyRequest) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *UpdateAuthPolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAuthPolicyRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateAuthPolicyRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateAuthPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateAuthPolicyRequest) SetAcceptLanguage(v string) *UpdateAuthPolicyRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetAppId(v string) *UpdateAuthPolicyRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetAuthRule(v string) *UpdateAuthPolicyRequest {
	s.AuthRule = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetEnable(v string) *UpdateAuthPolicyRequest {
	s.Enable = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetId(v string) *UpdateAuthPolicyRequest {
	s.Id = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetK8sNamespace(v string) *UpdateAuthPolicyRequest {
	s.K8sNamespace = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetName(v string) *UpdateAuthPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetProtocol(v string) *UpdateAuthPolicyRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetRegion(v string) *UpdateAuthPolicyRequest {
	s.Region = &v
	return s
}

func (s *UpdateAuthPolicyRequest) SetSource(v string) *UpdateAuthPolicyRequest {
	s.Source = &v
	return s
}

func (s *UpdateAuthPolicyRequest) Validate() error {
	return dara.Validate(s)
}
