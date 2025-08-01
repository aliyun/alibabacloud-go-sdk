// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIsolationRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteIsolationRulesShrinkRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteIsolationRulesShrinkRequest
	GetAppName() *string
	SetIdsShrink(v string) *DeleteIsolationRulesShrinkRequest
	GetIdsShrink() *string
	SetNamespace(v string) *DeleteIsolationRulesShrinkRequest
	GetNamespace() *string
}

type DeleteIsolationRulesShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteIsolationRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIsolationRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteIsolationRulesShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteIsolationRulesShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteIsolationRulesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DeleteIsolationRulesShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteIsolationRulesShrinkRequest) SetAcceptLanguage(v string) *DeleteIsolationRulesShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteIsolationRulesShrinkRequest) SetAppName(v string) *DeleteIsolationRulesShrinkRequest {
	s.AppName = &v
	return s
}

func (s *DeleteIsolationRulesShrinkRequest) SetIdsShrink(v string) *DeleteIsolationRulesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DeleteIsolationRulesShrinkRequest) SetNamespace(v string) *DeleteIsolationRulesShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteIsolationRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
