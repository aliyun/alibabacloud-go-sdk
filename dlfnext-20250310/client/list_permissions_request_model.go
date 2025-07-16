// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v string) *ListPermissionsRequest
	GetDatabase() *string
	SetFunction(v string) *ListPermissionsRequest
	GetFunction() *string
	SetMaxResults(v int32) *ListPermissionsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListPermissionsRequest
	GetPageToken() *string
	SetPrincipal(v string) *ListPermissionsRequest
	GetPrincipal() *string
	SetResourceType(v string) *ListPermissionsRequest
	GetResourceType() *string
	SetTable(v string) *ListPermissionsRequest
	GetTable() *string
	SetView(v string) *ListPermissionsRequest
	GetView() *string
}

type ListPermissionsRequest struct {
	// example:
	//
	// database_name
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	Function *string `json:"function,omitempty" xml:"function,omitempty"`
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:user/user_name
	Principal *string `json:"principal,omitempty" xml:"principal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CATALOG
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// table_name
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	View  *string `json:"view,omitempty" xml:"view,omitempty"`
}

func (s ListPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) GetDatabase() *string {
	return s.Database
}

func (s *ListPermissionsRequest) GetFunction() *string {
	return s.Function
}

func (s *ListPermissionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPermissionsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListPermissionsRequest) GetPrincipal() *string {
	return s.Principal
}

func (s *ListPermissionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPermissionsRequest) GetTable() *string {
	return s.Table
}

func (s *ListPermissionsRequest) GetView() *string {
	return s.View
}

func (s *ListPermissionsRequest) SetDatabase(v string) *ListPermissionsRequest {
	s.Database = &v
	return s
}

func (s *ListPermissionsRequest) SetFunction(v string) *ListPermissionsRequest {
	s.Function = &v
	return s
}

func (s *ListPermissionsRequest) SetMaxResults(v int32) *ListPermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionsRequest) SetPageToken(v string) *ListPermissionsRequest {
	s.PageToken = &v
	return s
}

func (s *ListPermissionsRequest) SetPrincipal(v string) *ListPermissionsRequest {
	s.Principal = &v
	return s
}

func (s *ListPermissionsRequest) SetResourceType(v string) *ListPermissionsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionsRequest) SetTable(v string) *ListPermissionsRequest {
	s.Table = &v
	return s
}

func (s *ListPermissionsRequest) SetView(v string) *ListPermissionsRequest {
	s.View = &v
	return s
}

func (s *ListPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
