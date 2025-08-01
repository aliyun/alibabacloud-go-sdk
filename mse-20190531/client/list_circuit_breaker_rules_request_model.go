// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCircuitBreakerRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListCircuitBreakerRulesRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *ListCircuitBreakerRulesRequest
	GetAppId() *string
	SetAppName(v string) *ListCircuitBreakerRulesRequest
	GetAppName() *string
	SetNamespace(v string) *ListCircuitBreakerRulesRequest
	GetNamespace() *string
	SetPageIndex(v int32) *ListCircuitBreakerRulesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListCircuitBreakerRulesRequest
	GetPageSize() *int32
	SetResource(v string) *ListCircuitBreakerRulesRequest
	GetResource() *string
	SetResourceSearchKey(v string) *ListCircuitBreakerRulesRequest
	GetResourceSearchKey() *string
}

type ListCircuitBreakerRulesRequest struct {
	// The language of the response. Valid values: zh-CN and en-US. Default value: zh-CN. The value zh-CN indicates Chinese, and the value en-US indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The microservice namespace to which the application belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The start page of the returned pages. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page. Default value: 6.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is used for exact match of circuit breaking rules.
	//
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// This parameter is used for fuzzy match of circuit breaking rules.
	//
	// example:
	//
	// a
	ResourceSearchKey *string `json:"ResourceSearchKey,omitempty" xml:"ResourceSearchKey,omitempty"`
}

func (s ListCircuitBreakerRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCircuitBreakerRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCircuitBreakerRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListCircuitBreakerRulesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListCircuitBreakerRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListCircuitBreakerRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListCircuitBreakerRulesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListCircuitBreakerRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCircuitBreakerRulesRequest) GetResource() *string {
	return s.Resource
}

func (s *ListCircuitBreakerRulesRequest) GetResourceSearchKey() *string {
	return s.ResourceSearchKey
}

func (s *ListCircuitBreakerRulesRequest) SetAcceptLanguage(v string) *ListCircuitBreakerRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCircuitBreakerRulesRequest) SetAppId(v string) *ListCircuitBreakerRulesRequest {
	s.AppId = &v
	return s
}

func (s *ListCircuitBreakerRulesRequest) SetAppName(v string) *ListCircuitBreakerRulesRequest {
	s.AppName = &v
	return s
}

func (s *ListCircuitBreakerRulesRequest) SetNamespace(v string) *ListCircuitBreakerRulesRequest {
	s.Namespace = &v
	return s
}

func (s *ListCircuitBreakerRulesRequest) SetPageIndex(v int32) *ListCircuitBreakerRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListCircuitBreakerRulesRequest) SetPageSize(v int32) *ListCircuitBreakerRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCircuitBreakerRulesRequest) SetResource(v string) *ListCircuitBreakerRulesRequest {
	s.Resource = &v
	return s
}

func (s *ListCircuitBreakerRulesRequest) SetResourceSearchKey(v string) *ListCircuitBreakerRulesRequest {
	s.ResourceSearchKey = &v
	return s
}

func (s *ListCircuitBreakerRulesRequest) Validate() error {
	return dara.Validate(s)
}
