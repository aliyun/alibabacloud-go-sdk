// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ModifyMaterialFileRequest
	GetBizId() *string
	SetFileId(v string) *ModifyMaterialFileRequest
	GetFileId() *string
	SetName(v string) *ModifyMaterialFileRequest
	GetName() *string
}

type ModifyMaterialFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WS12345678
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyMaterialFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaterialFileRequest) GetBizId() *string {
	return s.BizId
}

func (s *ModifyMaterialFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *ModifyMaterialFileRequest) GetName() *string {
	return s.Name
}

func (s *ModifyMaterialFileRequest) SetBizId(v string) *ModifyMaterialFileRequest {
	s.BizId = &v
	return s
}

func (s *ModifyMaterialFileRequest) SetFileId(v string) *ModifyMaterialFileRequest {
	s.FileId = &v
	return s
}

func (s *ModifyMaterialFileRequest) SetName(v string) *ModifyMaterialFileRequest {
	s.Name = &v
	return s
}

func (s *ModifyMaterialFileRequest) Validate() error {
	return dara.Validate(s)
}
