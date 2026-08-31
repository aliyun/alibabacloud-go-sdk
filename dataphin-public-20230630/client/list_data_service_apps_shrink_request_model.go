// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAppsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataServiceAppsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataServiceAppsShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListDataServiceAppsShrinkRequest
	GetOpUserId() *string
}

type ListDataServiceAppsShrinkRequest struct {
	// Query the list of all applications under the tenant.
	//
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListDataServiceAppsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAppsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceAppsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataServiceAppsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceAppsShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListDataServiceAppsShrinkRequest) SetListQueryShrink(v string) *ListDataServiceAppsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataServiceAppsShrinkRequest) SetOpTenantId(v int64) *ListDataServiceAppsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceAppsShrinkRequest) SetOpUserId(v string) *ListDataServiceAppsShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListDataServiceAppsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
