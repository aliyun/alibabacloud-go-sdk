// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSentinelBlockFallbackDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *BindSentinelBlockFallbackDefinitionRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *BindSentinelBlockFallbackDefinitionRequest
	GetAppName() *string
	SetFallbackId(v int64) *BindSentinelBlockFallbackDefinitionRequest
	GetFallbackId() *int64
	SetNamespace(v string) *BindSentinelBlockFallbackDefinitionRequest
	GetNamespace() *string
	SetResource(v string) *BindSentinelBlockFallbackDefinitionRequest
	GetResource() *string
	SetTargetType(v string) *BindSentinelBlockFallbackDefinitionRequest
	GetTargetType() *string
}

type BindSentinelBlockFallbackDefinitionRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Behavior ID. 0:the default behavior.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21
	FallbackId *int64 `json:"FallbackId,omitempty" xml:"FallbackId,omitempty"`
	// The microservice namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Interface Name: The resource to which the rule applies. It must match the interface name in the console\\"s interface details.
	//
	// This parameter is required.
	//
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Target rule type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s BindSentinelBlockFallbackDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s BindSentinelBlockFallbackDefinitionRequest) GoString() string {
	return s.String()
}

func (s *BindSentinelBlockFallbackDefinitionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *BindSentinelBlockFallbackDefinitionRequest) GetAppName() *string {
	return s.AppName
}

func (s *BindSentinelBlockFallbackDefinitionRequest) GetFallbackId() *int64 {
	return s.FallbackId
}

func (s *BindSentinelBlockFallbackDefinitionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BindSentinelBlockFallbackDefinitionRequest) GetResource() *string {
	return s.Resource
}

func (s *BindSentinelBlockFallbackDefinitionRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *BindSentinelBlockFallbackDefinitionRequest) SetAcceptLanguage(v string) *BindSentinelBlockFallbackDefinitionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionRequest) SetAppName(v string) *BindSentinelBlockFallbackDefinitionRequest {
	s.AppName = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionRequest) SetFallbackId(v int64) *BindSentinelBlockFallbackDefinitionRequest {
	s.FallbackId = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionRequest) SetNamespace(v string) *BindSentinelBlockFallbackDefinitionRequest {
	s.Namespace = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionRequest) SetResource(v string) *BindSentinelBlockFallbackDefinitionRequest {
	s.Resource = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionRequest) SetTargetType(v string) *BindSentinelBlockFallbackDefinitionRequest {
	s.TargetType = &v
	return s
}

func (s *BindSentinelBlockFallbackDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
