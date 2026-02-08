// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryMaterialFileDetailRequest
	GetBizId() *string
	SetFileId(v string) *QueryMaterialFileDetailRequest
	GetFileId() *string
}

type QueryMaterialFileDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s QueryMaterialFileDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileDetailRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryMaterialFileDetailRequest) GetFileId() *string {
	return s.FileId
}

func (s *QueryMaterialFileDetailRequest) SetBizId(v string) *QueryMaterialFileDetailRequest {
	s.BizId = &v
	return s
}

func (s *QueryMaterialFileDetailRequest) SetFileId(v string) *QueryMaterialFileDetailRequest {
	s.FileId = &v
	return s
}

func (s *QueryMaterialFileDetailRequest) Validate() error {
	return dara.Validate(s)
}
