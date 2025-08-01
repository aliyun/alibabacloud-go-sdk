// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteFlowRulesShrinkRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteFlowRulesShrinkRequest
	GetAppName() *string
	SetIdsShrink(v string) *DeleteFlowRulesShrinkRequest
	GetIdsShrink() *string
	SetNamespace(v string) *DeleteFlowRulesShrinkRequest
	GetNamespace() *string
}

type DeleteFlowRulesShrinkRequest struct {
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
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The IDs of the rules to be deleted.
	//
	// example:
	//
	// [1,2]
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

func (s DeleteFlowRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRulesShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteFlowRulesShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteFlowRulesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DeleteFlowRulesShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteFlowRulesShrinkRequest) SetAcceptLanguage(v string) *DeleteFlowRulesShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteFlowRulesShrinkRequest) SetAppName(v string) *DeleteFlowRulesShrinkRequest {
	s.AppName = &v
	return s
}

func (s *DeleteFlowRulesShrinkRequest) SetIdsShrink(v string) *DeleteFlowRulesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DeleteFlowRulesShrinkRequest) SetNamespace(v string) *DeleteFlowRulesShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteFlowRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
