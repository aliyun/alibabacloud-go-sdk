// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityWatchAlertShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityWatchAlertShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpsertQualityWatchAlertShrinkRequest
	GetOpUserId() *string
	SetUpsertCommandShrink(v string) *UpsertQualityWatchAlertShrinkRequest
	GetUpsertCommandShrink() *string
}

type UpsertQualityWatchAlertShrinkRequest struct {
	// The tenant ID.
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
	// The update instruction.
	//
	// This parameter is required.
	UpsertCommandShrink *string `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty"`
}

func (s UpsertQualityWatchAlertShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchAlertShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchAlertShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityWatchAlertShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpsertQualityWatchAlertShrinkRequest) GetUpsertCommandShrink() *string {
	return s.UpsertCommandShrink
}

func (s *UpsertQualityWatchAlertShrinkRequest) SetOpTenantId(v int64) *UpsertQualityWatchAlertShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityWatchAlertShrinkRequest) SetOpUserId(v string) *UpsertQualityWatchAlertShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpsertQualityWatchAlertShrinkRequest) SetUpsertCommandShrink(v string) *UpsertQualityWatchAlertShrinkRequest {
	s.UpsertCommandShrink = &v
	return s
}

func (s *UpsertQualityWatchAlertShrinkRequest) Validate() error {
	return dara.Validate(s)
}
