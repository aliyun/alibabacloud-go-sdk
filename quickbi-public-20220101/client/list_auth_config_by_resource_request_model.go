// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthConfigByResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHasAuth(v bool) *ListAuthConfigByResourceRequest
	GetHasAuth() *bool
	SetKeyword(v string) *ListAuthConfigByResourceRequest
	GetKeyword() *string
	SetResourceId(v string) *ListAuthConfigByResourceRequest
	GetResourceId() *string
	SetShareToType(v int32) *ListAuthConfigByResourceRequest
	GetShareToType() *int32
	SetWorksType(v string) *ListAuthConfigByResourceRequest
	GetWorksType() *string
}

type ListAuthConfigByResourceRequest struct {
	// Specifies whether to query only records that have permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	HasAuth *bool `json:"HasAuth,omitempty" xml:"HasAuth,omitempty"`
	// The search keyword for the authorized object.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The ID of the resource to authorize.
	//
	// This parameter is required.
	//
	// example:
	//
	// al*************7ufv
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The query granularity. Valid values:
	//
	// - 0: user
	//
	// - 1: user group
	//
	// 	Notice: Authorization records for organizations and workspaces are returned at the user granularity. Datasets and data sources can only be authorized to users.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ShareToType *int32 `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
	// The type of the authorization resource. Valid values:
	//
	// - dashboard: dashboard
	//
	// - screen: data screen
	//
	// - report: electron
	//
	// - ANALYSIS: ad hoc analysis
	//
	// - dashboardOfflineQuery: self-service data retrieval
	//
	// - dataForm: data form
	//
	// - cube: dataset
	//
	// - datasource: data source
	//
	// This parameter is required.
	//
	// example:
	//
	// REPORT
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
}

func (s ListAuthConfigByResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthConfigByResourceRequest) GoString() string {
	return s.String()
}

func (s *ListAuthConfigByResourceRequest) GetHasAuth() *bool {
	return s.HasAuth
}

func (s *ListAuthConfigByResourceRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAuthConfigByResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAuthConfigByResourceRequest) GetShareToType() *int32 {
	return s.ShareToType
}

func (s *ListAuthConfigByResourceRequest) GetWorksType() *string {
	return s.WorksType
}

func (s *ListAuthConfigByResourceRequest) SetHasAuth(v bool) *ListAuthConfigByResourceRequest {
	s.HasAuth = &v
	return s
}

func (s *ListAuthConfigByResourceRequest) SetKeyword(v string) *ListAuthConfigByResourceRequest {
	s.Keyword = &v
	return s
}

func (s *ListAuthConfigByResourceRequest) SetResourceId(v string) *ListAuthConfigByResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *ListAuthConfigByResourceRequest) SetShareToType(v int32) *ListAuthConfigByResourceRequest {
	s.ShareToType = &v
	return s
}

func (s *ListAuthConfigByResourceRequest) SetWorksType(v string) *ListAuthConfigByResourceRequest {
	s.WorksType = &v
	return s
}

func (s *ListAuthConfigByResourceRequest) Validate() error {
	return dara.Validate(s)
}
