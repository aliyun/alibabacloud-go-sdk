// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityRuleShrinkRequest
	GetOpTenantId() *int64
	SetUpsertCommandShrink(v string) *UpsertQualityRuleShrinkRequest
	GetUpsertCommandShrink() *string
}

type UpsertQualityRuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommandShrink *string `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty"`
}

func (s UpsertQualityRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityRuleShrinkRequest) GetUpsertCommandShrink() *string {
	return s.UpsertCommandShrink
}

func (s *UpsertQualityRuleShrinkRequest) SetOpTenantId(v int64) *UpsertQualityRuleShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityRuleShrinkRequest) SetUpsertCommandShrink(v string) *UpsertQualityRuleShrinkRequest {
	s.UpsertCommandShrink = &v
	return s
}

func (s *UpsertQualityRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
