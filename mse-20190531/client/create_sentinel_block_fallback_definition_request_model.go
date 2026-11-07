// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSentinelBlockFallbackDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetAppId() *string
	SetAppName(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetAppName() *string
	SetFallbackBehavior(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetFallbackBehavior() *string
	SetLanguage(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetLanguage() *string
	SetName(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetName() *string
	SetNamespace(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetNamespace() *string
	SetRegionId(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetRegionId() *string
	SetResourceClassification(v int32) *CreateSentinelBlockFallbackDefinitionRequest
	GetResourceClassification() *int32
	SetScenario(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetScenario() *string
	SetSource(v string) *CreateSentinelBlockFallbackDefinitionRequest
	GetSource() *string
}

type CreateSentinelBlockFallbackDefinitionRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh is the default value, which indicates Chinese. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application ID, which can be omitted.
	//
	// example:
	//
	// abcde@12345
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The behavior details.
	//
	// example:
	//
	// {"webFallbackMode":0,"webRespMessage":"{}","webRespStatusCode":429,"webRespContentType":0}
	FallbackBehavior *string `json:"FallbackBehavior,omitempty" xml:"FallbackBehavior,omitempty"`
	// The programming language, such as Java or Go.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The behavior name.
	//
	// example:
	//
	// MyFallback
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Microservices Engine (MSE) namespace.
	//
	// example:
	//
	// test
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The target resource type.
	//
	// example:
	//
	// 1 Web
	//
	// 2 Rpc
	ResourceClassification *int32 `json:"ResourceClassification,omitempty" xml:"ResourceClassification,omitempty"`
	// The scenario, uniformly fill in default.
	//
	// example:
	//
	// default
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// Redundant parameter, currently ignored.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateSentinelBlockFallbackDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSentinelBlockFallbackDefinitionRequest) GoString() string {
	return s.String()
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetFallbackBehavior() *string {
	return s.FallbackBehavior
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetName() *string {
	return s.Name
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetResourceClassification() *int32 {
	return s.ResourceClassification
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetScenario() *string {
	return s.Scenario
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) GetSource() *string {
	return s.Source
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetAcceptLanguage(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetAppId(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.AppId = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetAppName(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.AppName = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetFallbackBehavior(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.FallbackBehavior = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetLanguage(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.Language = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetName(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.Name = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetNamespace(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.Namespace = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetRegionId(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetResourceClassification(v int32) *CreateSentinelBlockFallbackDefinitionRequest {
	s.ResourceClassification = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetScenario(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.Scenario = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) SetSource(v string) *CreateSentinelBlockFallbackDefinitionRequest {
	s.Source = &v
	return s
}

func (s *CreateSentinelBlockFallbackDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
