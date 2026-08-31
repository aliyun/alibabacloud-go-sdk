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
	SetOpUserId(v string) *DeleteQualityTemplatesShrinkRequest
	GetOpUserId() *string
}

type DeleteQualityTemplatesShrinkRequest struct {
	// The delete instruction.
	//
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
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

func (s *DeleteQualityTemplatesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteQualityTemplatesShrinkRequest) SetDeleteCommandShrink(v string) *DeleteQualityTemplatesShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteQualityTemplatesShrinkRequest) SetOpTenantId(v int64) *DeleteQualityTemplatesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualityTemplatesShrinkRequest) SetOpUserId(v string) *DeleteQualityTemplatesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteQualityTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
