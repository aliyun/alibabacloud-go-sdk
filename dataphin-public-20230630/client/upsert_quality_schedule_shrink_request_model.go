// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityScheduleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityScheduleShrinkRequest
	GetOpTenantId() *int64
	SetUpsertCommandShrink(v string) *UpsertQualityScheduleShrinkRequest
	GetUpsertCommandShrink() *string
}

type UpsertQualityScheduleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommandShrink *string `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty"`
}

func (s UpsertQualityScheduleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityScheduleShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityScheduleShrinkRequest) GetUpsertCommandShrink() *string {
	return s.UpsertCommandShrink
}

func (s *UpsertQualityScheduleShrinkRequest) SetOpTenantId(v int64) *UpsertQualityScheduleShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityScheduleShrinkRequest) SetUpsertCommandShrink(v string) *UpsertQualityScheduleShrinkRequest {
	s.UpsertCommandShrink = &v
	return s
}

func (s *UpsertQualityScheduleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
