// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityWatchTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitQualityWatchTasksShrinkRequest
	GetOpTenantId() *int64
	SetSubmitCommandShrink(v string) *SubmitQualityWatchTasksShrinkRequest
	GetSubmitCommandShrink() *string
}

type SubmitQualityWatchTasksShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SubmitCommandShrink *string `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty"`
}

func (s SubmitQualityWatchTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityWatchTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitQualityWatchTasksShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitQualityWatchTasksShrinkRequest) GetSubmitCommandShrink() *string {
	return s.SubmitCommandShrink
}

func (s *SubmitQualityWatchTasksShrinkRequest) SetOpTenantId(v int64) *SubmitQualityWatchTasksShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitQualityWatchTasksShrinkRequest) SetSubmitCommandShrink(v string) *SubmitQualityWatchTasksShrinkRequest {
	s.SubmitCommandShrink = &v
	return s
}

func (s *SubmitQualityWatchTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
