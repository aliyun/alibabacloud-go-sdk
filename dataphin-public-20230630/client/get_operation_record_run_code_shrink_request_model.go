// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordRunCodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeCommandShrink(v string) *GetOperationRecordRunCodeShrinkRequest
	GetCodeCommandShrink() *string
	SetOpTenantId(v int64) *GetOperationRecordRunCodeShrinkRequest
	GetOpTenantId() *int64
}

type GetOperationRecordRunCodeShrinkRequest struct {
	// The query command.
	//
	// This parameter is required.
	CodeCommandShrink *string `json:"CodeCommand,omitempty" xml:"CodeCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetOperationRecordRunCodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordRunCodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRecordRunCodeShrinkRequest) GetCodeCommandShrink() *string {
	return s.CodeCommandShrink
}

func (s *GetOperationRecordRunCodeShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetOperationRecordRunCodeShrinkRequest) SetCodeCommandShrink(v string) *GetOperationRecordRunCodeShrinkRequest {
	s.CodeCommandShrink = &v
	return s
}

func (s *GetOperationRecordRunCodeShrinkRequest) SetOpTenantId(v int64) *GetOperationRecordRunCodeShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetOperationRecordRunCodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
