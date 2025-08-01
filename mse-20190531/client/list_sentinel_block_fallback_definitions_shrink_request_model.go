// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSentinelBlockFallbackDefinitionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListSentinelBlockFallbackDefinitionsShrinkRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ListSentinelBlockFallbackDefinitionsShrinkRequest
	GetAppName() *string
	SetClassificationSetShrink(v string) *ListSentinelBlockFallbackDefinitionsShrinkRequest
	GetClassificationSetShrink() *string
	SetNamespace(v string) *ListSentinelBlockFallbackDefinitionsShrinkRequest
	GetNamespace() *string
}

type ListSentinelBlockFallbackDefinitionsShrinkRequest struct {
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
	// Behavior Classification Set.
	ClassificationSetShrink *string `json:"ClassificationSet,omitempty" xml:"ClassificationSet,omitempty"`
	// The name of the Microservices namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListSentinelBlockFallbackDefinitionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSentinelBlockFallbackDefinitionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) GetClassificationSetShrink() *string {
	return s.ClassificationSetShrink
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) SetAcceptLanguage(v string) *ListSentinelBlockFallbackDefinitionsShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) SetAppName(v string) *ListSentinelBlockFallbackDefinitionsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) SetClassificationSetShrink(v string) *ListSentinelBlockFallbackDefinitionsShrinkRequest {
	s.ClassificationSetShrink = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) SetNamespace(v string) *ListSentinelBlockFallbackDefinitionsShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
