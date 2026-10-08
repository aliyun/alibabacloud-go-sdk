// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchHandoverAssetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHandoverCommandShrink(v string) *BatchHandoverAssetShrinkRequest
	GetHandoverCommandShrink() *string
	SetOpTenantId(v int64) *BatchHandoverAssetShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *BatchHandoverAssetShrinkRequest
	GetOpUserId() *string
}

type BatchHandoverAssetShrinkRequest struct {
	// This parameter is required.
	HandoverCommandShrink *string `json:"HandoverCommand,omitempty" xml:"HandoverCommand,omitempty"`
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

func (s BatchHandoverAssetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchHandoverAssetShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchHandoverAssetShrinkRequest) GetHandoverCommandShrink() *string {
	return s.HandoverCommandShrink
}

func (s *BatchHandoverAssetShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *BatchHandoverAssetShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *BatchHandoverAssetShrinkRequest) SetHandoverCommandShrink(v string) *BatchHandoverAssetShrinkRequest {
	s.HandoverCommandShrink = &v
	return s
}

func (s *BatchHandoverAssetShrinkRequest) SetOpTenantId(v int64) *BatchHandoverAssetShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *BatchHandoverAssetShrinkRequest) SetOpUserId(v string) *BatchHandoverAssetShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *BatchHandoverAssetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
