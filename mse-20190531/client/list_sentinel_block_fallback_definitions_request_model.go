// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSentinelBlockFallbackDefinitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListSentinelBlockFallbackDefinitionsRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ListSentinelBlockFallbackDefinitionsRequest
	GetAppName() *string
	SetClassificationSet(v []*int32) *ListSentinelBlockFallbackDefinitionsRequest
	GetClassificationSet() []*int32
	SetNamespace(v string) *ListSentinelBlockFallbackDefinitionsRequest
	GetNamespace() *string
}

type ListSentinelBlockFallbackDefinitionsRequest struct {
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
	ClassificationSet []*int32 `json:"ClassificationSet,omitempty" xml:"ClassificationSet,omitempty" type:"Repeated"`
	// The name of the Microservices namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListSentinelBlockFallbackDefinitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSentinelBlockFallbackDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) GetClassificationSet() []*int32 {
	return s.ClassificationSet
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) SetAcceptLanguage(v string) *ListSentinelBlockFallbackDefinitionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) SetAppName(v string) *ListSentinelBlockFallbackDefinitionsRequest {
	s.AppName = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) SetClassificationSet(v []*int32) *ListSentinelBlockFallbackDefinitionsRequest {
	s.ClassificationSet = v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) SetNamespace(v string) *ListSentinelBlockFallbackDefinitionsRequest {
	s.Namespace = &v
	return s
}

func (s *ListSentinelBlockFallbackDefinitionsRequest) Validate() error {
	return dara.Validate(s)
}
