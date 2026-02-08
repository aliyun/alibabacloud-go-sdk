// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveMaterialFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *MoveMaterialFileShrinkRequest
	GetBizId() *string
	SetDirectoryId(v string) *MoveMaterialFileShrinkRequest
	GetDirectoryId() *string
	SetFileIdsShrink(v string) *MoveMaterialFileShrinkRequest
	GetFileIdsShrink() *string
}

type MoveMaterialFileShrinkRequest struct {
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
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
}

func (s MoveMaterialFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveMaterialFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *MoveMaterialFileShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *MoveMaterialFileShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *MoveMaterialFileShrinkRequest) GetFileIdsShrink() *string {
	return s.FileIdsShrink
}

func (s *MoveMaterialFileShrinkRequest) SetBizId(v string) *MoveMaterialFileShrinkRequest {
	s.BizId = &v
	return s
}

func (s *MoveMaterialFileShrinkRequest) SetDirectoryId(v string) *MoveMaterialFileShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *MoveMaterialFileShrinkRequest) SetFileIdsShrink(v string) *MoveMaterialFileShrinkRequest {
	s.FileIdsShrink = &v
	return s
}

func (s *MoveMaterialFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
