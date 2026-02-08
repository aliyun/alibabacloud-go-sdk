// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialFileStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ModifyMaterialFileStatusShrinkRequest
	GetBizId() *string
	SetFileIdsShrink(v string) *ModifyMaterialFileStatusShrinkRequest
	GetFileIdsShrink() *string
	SetStatus(v string) *ModifyMaterialFileStatusShrinkRequest
	GetStatus() *string
}

type ModifyMaterialFileStatusShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WS12345678
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyMaterialFileStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialFileStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaterialFileStatusShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *ModifyMaterialFileStatusShrinkRequest) GetFileIdsShrink() *string {
	return s.FileIdsShrink
}

func (s *ModifyMaterialFileStatusShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyMaterialFileStatusShrinkRequest) SetBizId(v string) *ModifyMaterialFileStatusShrinkRequest {
	s.BizId = &v
	return s
}

func (s *ModifyMaterialFileStatusShrinkRequest) SetFileIdsShrink(v string) *ModifyMaterialFileStatusShrinkRequest {
	s.FileIdsShrink = &v
	return s
}

func (s *ModifyMaterialFileStatusShrinkRequest) SetStatus(v string) *ModifyMaterialFileStatusShrinkRequest {
	s.Status = &v
	return s
}

func (s *ModifyMaterialFileStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
