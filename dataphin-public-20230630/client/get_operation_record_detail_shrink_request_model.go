// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetOperationRecordDetailShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetOperationRecordDetailShrinkRequest
	GetOpUserId() *string
	SetRecordDetailCommandShrink(v string) *GetOperationRecordDetailShrinkRequest
	GetRecordDetailCommandShrink() *string
}

type GetOperationRecordDetailShrinkRequest struct {
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
	// The query command.
	//
	// This parameter is required.
	RecordDetailCommandShrink *string `json:"RecordDetailCommand,omitempty" xml:"RecordDetailCommand,omitempty"`
}

func (s GetOperationRecordDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRecordDetailShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetOperationRecordDetailShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetOperationRecordDetailShrinkRequest) GetRecordDetailCommandShrink() *string {
	return s.RecordDetailCommandShrink
}

func (s *GetOperationRecordDetailShrinkRequest) SetOpTenantId(v int64) *GetOperationRecordDetailShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetOperationRecordDetailShrinkRequest) SetOpUserId(v string) *GetOperationRecordDetailShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *GetOperationRecordDetailShrinkRequest) SetRecordDetailCommandShrink(v string) *GetOperationRecordDetailShrinkRequest {
	s.RecordDetailCommandShrink = &v
	return s
}

func (s *GetOperationRecordDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
