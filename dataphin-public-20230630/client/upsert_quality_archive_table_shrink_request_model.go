// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityArchiveTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityArchiveTableShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpsertQualityArchiveTableShrinkRequest
	GetOpUserId() *string
	SetUpsertCommandShrink(v string) *UpsertQualityArchiveTableShrinkRequest
	GetUpsertCommandShrink() *string
}

type UpsertQualityArchiveTableShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The upsert command.
	//
	// This parameter is required.
	UpsertCommandShrink *string `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty"`
}

func (s UpsertQualityArchiveTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityArchiveTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityArchiveTableShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityArchiveTableShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpsertQualityArchiveTableShrinkRequest) GetUpsertCommandShrink() *string {
	return s.UpsertCommandShrink
}

func (s *UpsertQualityArchiveTableShrinkRequest) SetOpTenantId(v int64) *UpsertQualityArchiveTableShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityArchiveTableShrinkRequest) SetOpUserId(v string) *UpsertQualityArchiveTableShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpsertQualityArchiveTableShrinkRequest) SetUpsertCommandShrink(v string) *UpsertQualityArchiveTableShrinkRequest {
	s.UpsertCommandShrink = &v
	return s
}

func (s *UpsertQualityArchiveTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
