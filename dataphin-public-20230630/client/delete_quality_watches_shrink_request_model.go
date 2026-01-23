// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityWatchesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteQualityWatchesShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteQualityWatchesShrinkRequest
	GetOpTenantId() *int64
}

type DeleteQualityWatchesShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteQualityWatchesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityWatchesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityWatchesShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteQualityWatchesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteQualityWatchesShrinkRequest) SetDeleteCommandShrink(v string) *DeleteQualityWatchesShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteQualityWatchesShrinkRequest) SetOpTenantId(v int64) *DeleteQualityWatchesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualityWatchesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
