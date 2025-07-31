// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseBatchTaskDependencyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ParseBatchTaskDependencyShrinkRequest
	GetOpTenantId() *int64
	SetParseCommandShrink(v string) *ParseBatchTaskDependencyShrinkRequest
	GetParseCommandShrink() *string
}

type ParseBatchTaskDependencyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	ParseCommandShrink *string `json:"ParseCommand,omitempty" xml:"ParseCommand,omitempty"`
}

func (s ParseBatchTaskDependencyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ParseBatchTaskDependencyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ParseBatchTaskDependencyShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ParseBatchTaskDependencyShrinkRequest) GetParseCommandShrink() *string {
	return s.ParseCommandShrink
}

func (s *ParseBatchTaskDependencyShrinkRequest) SetOpTenantId(v int64) *ParseBatchTaskDependencyShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ParseBatchTaskDependencyShrinkRequest) SetParseCommandShrink(v string) *ParseBatchTaskDependencyShrinkRequest {
	s.ParseCommandShrink = &v
	return s
}

func (s *ParseBatchTaskDependencyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
