// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordByIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetailCommandShrink(v string) *GetOperationRecordByIdShrinkRequest
	GetDetailCommandShrink() *string
	SetOpTenantId(v int64) *GetOperationRecordByIdShrinkRequest
	GetOpTenantId() *int64
}

type GetOperationRecordByIdShrinkRequest struct {
	// This parameter is required.
	DetailCommandShrink *string `json:"DetailCommand,omitempty" xml:"DetailCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetOperationRecordByIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRecordByIdShrinkRequest) GetDetailCommandShrink() *string {
	return s.DetailCommandShrink
}

func (s *GetOperationRecordByIdShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetOperationRecordByIdShrinkRequest) SetDetailCommandShrink(v string) *GetOperationRecordByIdShrinkRequest {
	s.DetailCommandShrink = &v
	return s
}

func (s *GetOperationRecordByIdShrinkRequest) SetOpTenantId(v int64) *GetOperationRecordByIdShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetOperationRecordByIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
