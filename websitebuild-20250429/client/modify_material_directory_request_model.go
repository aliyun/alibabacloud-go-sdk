// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ModifyMaterialDirectoryRequest
	GetBizId() *string
	SetDirectoryId(v string) *ModifyMaterialDirectoryRequest
	GetDirectoryId() *string
	SetName(v string) *ModifyMaterialDirectoryRequest
	GetName() *string
}

type ModifyMaterialDirectoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyMaterialDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaterialDirectoryRequest) GetBizId() *string {
	return s.BizId
}

func (s *ModifyMaterialDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ModifyMaterialDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *ModifyMaterialDirectoryRequest) SetBizId(v string) *ModifyMaterialDirectoryRequest {
	s.BizId = &v
	return s
}

func (s *ModifyMaterialDirectoryRequest) SetDirectoryId(v string) *ModifyMaterialDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *ModifyMaterialDirectoryRequest) SetName(v string) *ModifyMaterialDirectoryRequest {
	s.Name = &v
	return s
}

func (s *ModifyMaterialDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
