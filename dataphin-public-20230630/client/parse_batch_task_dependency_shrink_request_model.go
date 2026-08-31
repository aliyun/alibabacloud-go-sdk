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
	SetOpUserId(v string) *ParseBatchTaskDependencyShrinkRequest
	GetOpUserId() *string
	SetParseCommandShrink(v string) *ParseBatchTaskDependencyShrinkRequest
	GetParseCommandShrink() *string
}

type ParseBatchTaskDependencyShrinkRequest struct {
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
	// The parse request.
	//
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

func (s *ParseBatchTaskDependencyShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ParseBatchTaskDependencyShrinkRequest) GetParseCommandShrink() *string {
	return s.ParseCommandShrink
}

func (s *ParseBatchTaskDependencyShrinkRequest) SetOpTenantId(v int64) *ParseBatchTaskDependencyShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ParseBatchTaskDependencyShrinkRequest) SetOpUserId(v string) *ParseBatchTaskDependencyShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ParseBatchTaskDependencyShrinkRequest) SetParseCommandShrink(v string) *ParseBatchTaskDependencyShrinkRequest {
	s.ParseCommandShrink = &v
	return s
}

func (s *ParseBatchTaskDependencyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
