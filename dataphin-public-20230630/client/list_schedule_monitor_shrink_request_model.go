// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleMonitorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListScheduleMonitorShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListScheduleMonitorShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListScheduleMonitorShrinkRequest
	GetOpUserId() *string
}

type ListScheduleMonitorShrinkRequest struct {
	// The query conditions for scheduling monitoring and alerting.
	//
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The user ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListScheduleMonitorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListScheduleMonitorShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListScheduleMonitorShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListScheduleMonitorShrinkRequest) SetListQueryShrink(v string) *ListScheduleMonitorShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListScheduleMonitorShrinkRequest) SetOpTenantId(v int64) *ListScheduleMonitorShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListScheduleMonitorShrinkRequest) SetOpUserId(v string) *ListScheduleMonitorShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListScheduleMonitorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
