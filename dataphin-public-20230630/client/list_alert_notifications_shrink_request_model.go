// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertNotificationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListAlertNotificationsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListAlertNotificationsShrinkRequest
	GetOpTenantId() *int64
}

type ListAlertNotificationsShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAlertNotificationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListAlertNotificationsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAlertNotificationsShrinkRequest) SetListQueryShrink(v string) *ListAlertNotificationsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListAlertNotificationsShrinkRequest) SetOpTenantId(v int64) *ListAlertNotificationsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAlertNotificationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
