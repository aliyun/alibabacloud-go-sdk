// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteQualityTemplatesShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteQualityTemplatesShrinkRequest
	GetOpTenantId() *int64
}

type DeleteQualityTemplatesShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteQualityTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityTemplatesShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteQualityTemplatesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteQualityTemplatesShrinkRequest) SetDeleteCommandShrink(v string) *DeleteQualityTemplatesShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteQualityTemplatesShrinkRequest) SetOpTenantId(v int64) *DeleteQualityTemplatesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualityTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
