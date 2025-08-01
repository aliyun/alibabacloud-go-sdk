// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCircuitBreakerRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteCircuitBreakerRulesShrinkRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteCircuitBreakerRulesShrinkRequest
	GetAppName() *string
	SetIdsShrink(v string) *DeleteCircuitBreakerRulesShrinkRequest
	GetIdsShrink() *string
	SetNamespace(v string) *DeleteCircuitBreakerRulesShrinkRequest
	GetNamespace() *string
}

type DeleteCircuitBreakerRulesShrinkRequest struct {
	// The language of the response. Valid values: zh and en. Default value: zh. The value zh indicates Chinese, and the value en indicates English.
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
	// The IDs of the rules that you want to delete.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The microservice namespace to which the application belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteCircuitBreakerRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCircuitBreakerRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) SetAcceptLanguage(v string) *DeleteCircuitBreakerRulesShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) SetAppName(v string) *DeleteCircuitBreakerRulesShrinkRequest {
	s.AppName = &v
	return s
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) SetIdsShrink(v string) *DeleteCircuitBreakerRulesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) SetNamespace(v string) *DeleteCircuitBreakerRulesShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteCircuitBreakerRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
