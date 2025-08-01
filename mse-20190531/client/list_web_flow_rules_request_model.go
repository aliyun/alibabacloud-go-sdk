// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFlowRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListWebFlowRulesRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *ListWebFlowRulesRequest
	GetAppId() *string
	SetAppName(v string) *ListWebFlowRulesRequest
	GetAppName() *string
	SetNamespace(v string) *ListWebFlowRulesRequest
	GetNamespace() *string
	SetPageIndex(v int32) *ListWebFlowRulesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListWebFlowRulesRequest
	GetPageSize() *int32
	SetResource(v string) *ListWebFlowRulesRequest
	GetResource() *string
	SetResourceSearchKey(v string) *ListWebFlowRulesRequest
	GetResourceSearchKey() *string
}

type ListWebFlowRulesRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// b22qb1****@2f0586be4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
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
	// /flow
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// flow
	ResourceSearchKey *string `json:"ResourceSearchKey,omitempty" xml:"ResourceSearchKey,omitempty"`
}

func (s ListWebFlowRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebFlowRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWebFlowRulesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListWebFlowRulesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListWebFlowRulesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListWebFlowRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListWebFlowRulesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListWebFlowRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWebFlowRulesRequest) GetResource() *string {
	return s.Resource
}

func (s *ListWebFlowRulesRequest) GetResourceSearchKey() *string {
	return s.ResourceSearchKey
}

func (s *ListWebFlowRulesRequest) SetAcceptLanguage(v string) *ListWebFlowRulesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListWebFlowRulesRequest) SetAppId(v string) *ListWebFlowRulesRequest {
	s.AppId = &v
	return s
}

func (s *ListWebFlowRulesRequest) SetAppName(v string) *ListWebFlowRulesRequest {
	s.AppName = &v
	return s
}

func (s *ListWebFlowRulesRequest) SetNamespace(v string) *ListWebFlowRulesRequest {
	s.Namespace = &v
	return s
}

func (s *ListWebFlowRulesRequest) SetPageIndex(v int32) *ListWebFlowRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListWebFlowRulesRequest) SetPageSize(v int32) *ListWebFlowRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWebFlowRulesRequest) SetResource(v string) *ListWebFlowRulesRequest {
	s.Resource = &v
	return s
}

func (s *ListWebFlowRulesRequest) SetResourceSearchKey(v string) *ListWebFlowRulesRequest {
	s.ResourceSearchKey = &v
	return s
}

func (s *ListWebFlowRulesRequest) Validate() error {
	return dara.Validate(s)
}
