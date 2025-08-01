// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListFlowRulesRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *ListFlowRulesRequest
	GetAppId() *string
	SetAppName(v string) *ListFlowRulesRequest
	GetAppName() *string
	SetNamespace(v string) *ListFlowRulesRequest
	GetNamespace() *string
	SetPageIndex(v int32) *ListFlowRulesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListFlowRulesRequest
	GetPageSize() *int32
	SetResource(v string) *ListFlowRulesRequest
	GetResource() *string
	SetResourceSearchKey(v string) *ListFlowRulesRequest
	GetResourceSearchKey() *string
}

type ListFlowRulesRequest struct {
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
	// The ID of the application.
	//
	// example:
	//
	// hkhon1po62@c3df23522******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx-demo
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// /a
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The keyword that is used for the search.
	//
	// example:
	//
	// /a
	ResourceSearchKey *string `json:"ResourceSearchKey,omitempty" xml:"ResourceSearchKey,omitempty"`
}

func (s ListFlowRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowRulesRequest) GoString() string {
	return s.String()
}

func (s *ListFlowRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListFlowRulesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListFlowRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListFlowRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListFlowRulesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListFlowRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlowRulesRequest) GetResource() *string {
	return s.Resource
}

func (s *ListFlowRulesRequest) GetResourceSearchKey() *string {
	return s.ResourceSearchKey
}

func (s *ListFlowRulesRequest) SetAcceptLanguage(v string) *ListFlowRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListFlowRulesRequest) SetAppId(v string) *ListFlowRulesRequest {
	s.AppId = &v
	return s
}

func (s *ListFlowRulesRequest) SetAppName(v string) *ListFlowRulesRequest {
	s.AppName = &v
	return s
}

func (s *ListFlowRulesRequest) SetNamespace(v string) *ListFlowRulesRequest {
	s.Namespace = &v
	return s
}

func (s *ListFlowRulesRequest) SetPageIndex(v int32) *ListFlowRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListFlowRulesRequest) SetPageSize(v int32) *ListFlowRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowRulesRequest) SetResource(v string) *ListFlowRulesRequest {
	s.Resource = &v
	return s
}

func (s *ListFlowRulesRequest) SetResourceSearchKey(v string) *ListFlowRulesRequest {
	s.ResourceSearchKey = &v
	return s
}

func (s *ListFlowRulesRequest) Validate() error {
	return dara.Validate(s)
}
