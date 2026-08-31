// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualitySchedulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteQualitySchedulesShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteQualitySchedulesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *DeleteQualitySchedulesShrinkRequest
	GetOpUserId() *string
}

type DeleteQualitySchedulesShrinkRequest struct {
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

func (s DeleteQualitySchedulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualitySchedulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualitySchedulesShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteQualitySchedulesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteQualitySchedulesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteQualitySchedulesShrinkRequest) SetDeleteCommandShrink(v string) *DeleteQualitySchedulesShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteQualitySchedulesShrinkRequest) SetOpTenantId(v int64) *DeleteQualitySchedulesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualitySchedulesShrinkRequest) SetOpUserId(v string) *DeleteQualitySchedulesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteQualitySchedulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
