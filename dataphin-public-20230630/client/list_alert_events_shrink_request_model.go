// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListAlertEventsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListAlertEventsShrinkRequest
	GetOpTenantId() *int64
}

type ListAlertEventsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAlertEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertEventsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListAlertEventsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAlertEventsShrinkRequest) SetListQueryShrink(v string) *ListAlertEventsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListAlertEventsShrinkRequest) SetOpTenantId(v int64) *ListAlertEventsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAlertEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
