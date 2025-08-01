// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteFlowRulesRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteFlowRulesRequest
	GetAppName() *string
	SetIds(v []*int64) *DeleteFlowRulesRequest
	GetIds() []*int64
	SetNamespace(v string) *DeleteFlowRulesRequest
	GetNamespace() *string
}

type DeleteFlowRulesRequest struct {
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
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The microservice namespace to which the application belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteFlowRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteFlowRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteFlowRulesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DeleteFlowRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteFlowRulesRequest) SetAcceptLanguage(v string) *DeleteFlowRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteFlowRulesRequest) SetAppName(v string) *DeleteFlowRulesRequest {
	s.AppName = &v
	return s
}

func (s *DeleteFlowRulesRequest) SetIds(v []*int64) *DeleteFlowRulesRequest {
	s.Ids = v
	return s
}

func (s *DeleteFlowRulesRequest) SetNamespace(v string) *DeleteFlowRulesRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteFlowRulesRequest) Validate() error {
	return dara.Validate(s)
}
