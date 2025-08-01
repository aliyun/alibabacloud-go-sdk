// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCircuitBreakerRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteCircuitBreakerRulesRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteCircuitBreakerRulesRequest
	GetAppName() *string
	SetIds(v []*int64) *DeleteCircuitBreakerRulesRequest
	GetIds() []*int64
	SetNamespace(v string) *DeleteCircuitBreakerRulesRequest
	GetNamespace() *string
}

type DeleteCircuitBreakerRulesRequest struct {
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

func (s DeleteCircuitBreakerRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCircuitBreakerRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCircuitBreakerRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteCircuitBreakerRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteCircuitBreakerRulesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DeleteCircuitBreakerRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteCircuitBreakerRulesRequest) SetAcceptLanguage(v string) *DeleteCircuitBreakerRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteCircuitBreakerRulesRequest) SetAppName(v string) *DeleteCircuitBreakerRulesRequest {
	s.AppName = &v
	return s
}

func (s *DeleteCircuitBreakerRulesRequest) SetIds(v []*int64) *DeleteCircuitBreakerRulesRequest {
	s.Ids = v
	return s
}

func (s *DeleteCircuitBreakerRulesRequest) SetNamespace(v string) *DeleteCircuitBreakerRulesRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteCircuitBreakerRulesRequest) Validate() error {
	return dara.Validate(s)
}
