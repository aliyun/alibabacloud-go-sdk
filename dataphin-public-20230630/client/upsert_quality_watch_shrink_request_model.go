// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityWatchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityWatchShrinkRequest
	GetOpTenantId() *int64
	SetUpsertCommandShrink(v string) *UpsertQualityWatchShrinkRequest
	GetUpsertCommandShrink() *string
}

type UpsertQualityWatchShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommandShrink *string `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty"`
}

func (s UpsertQualityWatchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityWatchShrinkRequest) GetUpsertCommandShrink() *string {
	return s.UpsertCommandShrink
}

func (s *UpsertQualityWatchShrinkRequest) SetOpTenantId(v int64) *UpsertQualityWatchShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityWatchShrinkRequest) SetUpsertCommandShrink(v string) *UpsertQualityWatchShrinkRequest {
	s.UpsertCommandShrink = &v
	return s
}

func (s *UpsertQualityWatchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
