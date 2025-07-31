// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddableUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListAddableUsersShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListAddableUsersShrinkRequest
	GetOpTenantId() *int64
}

type ListAddableUsersShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAddableUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddableUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAddableUsersShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListAddableUsersShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAddableUsersShrinkRequest) SetListQueryShrink(v string) *ListAddableUsersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListAddableUsersShrinkRequest) SetOpTenantId(v int64) *ListAddableUsersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAddableUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
