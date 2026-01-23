// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityIdentifyResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListSecurityIdentifyResultsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListSecurityIdentifyResultsShrinkRequest
	GetOpTenantId() *int64
}

type ListSecurityIdentifyResultsShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListSecurityIdentifyResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyResultsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListSecurityIdentifyResultsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListSecurityIdentifyResultsShrinkRequest) SetListQueryShrink(v string) *ListSecurityIdentifyResultsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListSecurityIdentifyResultsShrinkRequest) SetOpTenantId(v int64) *ListSecurityIdentifyResultsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListSecurityIdentifyResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
