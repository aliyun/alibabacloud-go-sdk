// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityTemplateShrinkRequest
	GetOpTenantId() *int64
	SetUpsertCommandShrink(v string) *UpsertQualityTemplateShrinkRequest
	GetUpsertCommandShrink() *string
}

type UpsertQualityTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommandShrink *string `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty"`
}

func (s UpsertQualityTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityTemplateShrinkRequest) GetUpsertCommandShrink() *string {
	return s.UpsertCommandShrink
}

func (s *UpsertQualityTemplateShrinkRequest) SetOpTenantId(v int64) *UpsertQualityTemplateShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityTemplateShrinkRequest) SetUpsertCommandShrink(v string) *UpsertQualityTemplateShrinkRequest {
	s.UpsertCommandShrink = &v
	return s
}

func (s *UpsertQualityTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
