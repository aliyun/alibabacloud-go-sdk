// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListCommandShrink(v string) *ListOperationRecordShrinkRequest
	GetListCommandShrink() *string
	SetOpTenantId(v int64) *ListOperationRecordShrinkRequest
	GetOpTenantId() *int64
}

type ListOperationRecordShrinkRequest struct {
	// The query command.
	//
	// This parameter is required.
	ListCommandShrink *string `json:"ListCommand,omitempty" xml:"ListCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListOperationRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOperationRecordShrinkRequest) GetListCommandShrink() *string {
	return s.ListCommandShrink
}

func (s *ListOperationRecordShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListOperationRecordShrinkRequest) SetListCommandShrink(v string) *ListOperationRecordShrinkRequest {
	s.ListCommandShrink = &v
	return s
}

func (s *ListOperationRecordShrinkRequest) SetOpTenantId(v int64) *ListOperationRecordShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListOperationRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
