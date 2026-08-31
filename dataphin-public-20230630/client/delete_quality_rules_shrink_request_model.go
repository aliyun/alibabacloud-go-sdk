// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteQualityRulesShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteQualityRulesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *DeleteQualityRulesShrinkRequest
	GetOpUserId() *string
}

type DeleteQualityRulesShrinkRequest struct {
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

func (s DeleteQualityRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityRulesShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteQualityRulesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteQualityRulesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteQualityRulesShrinkRequest) SetDeleteCommandShrink(v string) *DeleteQualityRulesShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteQualityRulesShrinkRequest) SetOpTenantId(v int64) *DeleteQualityRulesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualityRulesShrinkRequest) SetOpUserId(v string) *DeleteQualityRulesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteQualityRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
