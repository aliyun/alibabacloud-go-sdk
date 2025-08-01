// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIsolationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListIsolationRulesRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *ListIsolationRulesRequest
	GetAppId() *string
	SetAppName(v string) *ListIsolationRulesRequest
	GetAppName() *string
	SetNamespace(v string) *ListIsolationRulesRequest
	GetNamespace() *string
	SetPageIndex(v int32) *ListIsolationRulesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListIsolationRulesRequest
	GetPageSize() *int32
	SetResource(v string) *ListIsolationRulesRequest
	GetResource() *string
	SetResourceSearchKey(v string) *ListIsolationRulesRequest
	GetResourceSearchKey() *string
}

type ListIsolationRulesRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// hkhon1po62@c3df23522bXXXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// a
	ResourceSearchKey *string `json:"ResourceSearchKey,omitempty" xml:"ResourceSearchKey,omitempty"`
}

func (s ListIsolationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIsolationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListIsolationRulesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListIsolationRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListIsolationRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIsolationRulesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListIsolationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIsolationRulesRequest) GetResource() *string {
	return s.Resource
}

func (s *ListIsolationRulesRequest) GetResourceSearchKey() *string {
	return s.ResourceSearchKey
}

func (s *ListIsolationRulesRequest) SetAcceptLanguage(v string) *ListIsolationRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListIsolationRulesRequest) SetAppId(v string) *ListIsolationRulesRequest {
	s.AppId = &v
	return s
}

func (s *ListIsolationRulesRequest) SetAppName(v string) *ListIsolationRulesRequest {
	s.AppName = &v
	return s
}

func (s *ListIsolationRulesRequest) SetNamespace(v string) *ListIsolationRulesRequest {
	s.Namespace = &v
	return s
}

func (s *ListIsolationRulesRequest) SetPageIndex(v int32) *ListIsolationRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListIsolationRulesRequest) SetPageSize(v int32) *ListIsolationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListIsolationRulesRequest) SetResource(v string) *ListIsolationRulesRequest {
	s.Resource = &v
	return s
}

func (s *ListIsolationRulesRequest) SetResourceSearchKey(v string) *ListIsolationRulesRequest {
	s.ResourceSearchKey = &v
	return s
}

func (s *ListIsolationRulesRequest) Validate() error {
	return dara.Validate(s)
}
