// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteMaterialDirectoryRequest
	GetBizId() *string
	SetDirectoryId(v string) *DeleteMaterialDirectoryRequest
	GetDirectoryId() *string
}

type DeleteMaterialDirectoryRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s DeleteMaterialDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaterialDirectoryRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteMaterialDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteMaterialDirectoryRequest) SetBizId(v string) *DeleteMaterialDirectoryRequest {
	s.BizId = &v
	return s
}

func (s *DeleteMaterialDirectoryRequest) SetDirectoryId(v string) *DeleteMaterialDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteMaterialDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
