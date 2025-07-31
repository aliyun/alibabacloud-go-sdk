// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitBatchTaskShrinkRequest
	GetOpTenantId() *int64
	SetSubmitCommandShrink(v string) *SubmitBatchTaskShrinkRequest
	GetSubmitCommandShrink() *string
}

type SubmitBatchTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SubmitCommandShrink *string `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty"`
}

func (s SubmitBatchTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitBatchTaskShrinkRequest) GetSubmitCommandShrink() *string {
	return s.SubmitCommandShrink
}

func (s *SubmitBatchTaskShrinkRequest) SetOpTenantId(v int64) *SubmitBatchTaskShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitBatchTaskShrinkRequest) SetSubmitCommandShrink(v string) *SubmitBatchTaskShrinkRequest {
	s.SubmitCommandShrink = &v
	return s
}

func (s *SubmitBatchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
