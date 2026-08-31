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
	SetOpUserId(v string) *UpsertQualityScheduleShrinkRequest
	GetOpUserId() *string
	SetUpsertCommandShrink(v string) *UpsertQualityScheduleShrinkRequest
	GetUpsertCommandShrink() *string
}

type UpsertQualityScheduleShrinkRequest struct {
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

func (s UpsertQualityScheduleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityScheduleShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityScheduleShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpsertQualityScheduleShrinkRequest) GetUpsertCommandShrink() *string {
	return s.UpsertCommandShrink
}

func (s *UpsertQualityScheduleShrinkRequest) SetOpTenantId(v int64) *UpsertQualityScheduleShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityScheduleShrinkRequest) SetOpUserId(v string) *UpsertQualityScheduleShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpsertQualityScheduleShrinkRequest) SetUpsertCommandShrink(v string) *UpsertQualityScheduleShrinkRequest {
	s.UpsertCommandShrink = &v
	return s
}

func (s *UpsertQualityScheduleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
