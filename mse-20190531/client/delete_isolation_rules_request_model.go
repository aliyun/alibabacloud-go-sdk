// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIsolationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteIsolationRulesRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteIsolationRulesRequest
	GetAppName() *string
	SetIds(v []*int64) *DeleteIsolationRulesRequest
	GetIds() []*int64
	SetNamespace(v string) *DeleteIsolationRulesRequest
	GetNamespace() *string
}

type DeleteIsolationRulesRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Ids     []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteIsolationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIsolationRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteIsolationRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteIsolationRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteIsolationRulesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DeleteIsolationRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteIsolationRulesRequest) SetAcceptLanguage(v string) *DeleteIsolationRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteIsolationRulesRequest) SetAppName(v string) *DeleteIsolationRulesRequest {
	s.AppName = &v
	return s
}

func (s *DeleteIsolationRulesRequest) SetIds(v []*int64) *DeleteIsolationRulesRequest {
	s.Ids = v
	return s
}

func (s *DeleteIsolationRulesRequest) SetNamespace(v string) *DeleteIsolationRulesRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteIsolationRulesRequest) Validate() error {
	return dara.Validate(s)
}
