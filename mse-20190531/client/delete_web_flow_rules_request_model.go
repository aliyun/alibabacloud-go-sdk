// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFlowRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteWebFlowRulesRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteWebFlowRulesRequest
	GetAppName() *string
	SetIds(v string) *DeleteWebFlowRulesRequest
	GetIds() *string
	SetNamespace(v string) *DeleteWebFlowRulesRequest
	GetNamespace() *string
}

type DeleteWebFlowRulesRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [1, 2, 3]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteWebFlowRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFlowRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebFlowRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteWebFlowRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteWebFlowRulesRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteWebFlowRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteWebFlowRulesRequest) SetAcceptLanguage(v string) *DeleteWebFlowRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteWebFlowRulesRequest) SetAppName(v string) *DeleteWebFlowRulesRequest {
	s.AppName = &v
	return s
}

func (s *DeleteWebFlowRulesRequest) SetIds(v string) *DeleteWebFlowRulesRequest {
	s.Ids = &v
	return s
}

func (s *DeleteWebFlowRulesRequest) SetNamespace(v string) *DeleteWebFlowRulesRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteWebFlowRulesRequest) Validate() error {
	return dara.Validate(s)
}
