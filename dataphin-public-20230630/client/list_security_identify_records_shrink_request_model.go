// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityIdentifyRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListSecurityIdentifyRecordsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListSecurityIdentifyRecordsShrinkRequest
	GetOpTenantId() *int64
}

type ListSecurityIdentifyRecordsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListSecurityIdentifyRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyRecordsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListSecurityIdentifyRecordsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListSecurityIdentifyRecordsShrinkRequest) SetListQueryShrink(v string) *ListSecurityIdentifyRecordsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListSecurityIdentifyRecordsShrinkRequest) SetOpTenantId(v int64) *ListSecurityIdentifyRecordsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
