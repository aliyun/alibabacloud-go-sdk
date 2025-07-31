// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedDataServiceApiDetailsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListAuthorizedDataServiceApiDetailsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListAuthorizedDataServiceApiDetailsShrinkRequest
	GetOpTenantId() *int64
}

type ListAuthorizedDataServiceApiDetailsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListAuthorizedDataServiceApiDetailsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAuthorizedDataServiceApiDetailsShrinkRequest) SetListQueryShrink(v string) *ListAuthorizedDataServiceApiDetailsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsShrinkRequest) SetOpTenantId(v int64) *ListAuthorizedDataServiceApiDetailsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
