// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveMaterialFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *MoveMaterialFileRequest
	GetBizId() *string
	SetDirectoryId(v string) *MoveMaterialFileRequest
	GetDirectoryId() *string
	SetFileIds(v []*string) *MoveMaterialFileRequest
	GetFileIds() []*string
}

type MoveMaterialFileRequest struct {
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
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
}

func (s MoveMaterialFileRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveMaterialFileRequest) GoString() string {
	return s.String()
}

func (s *MoveMaterialFileRequest) GetBizId() *string {
	return s.BizId
}

func (s *MoveMaterialFileRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *MoveMaterialFileRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *MoveMaterialFileRequest) SetBizId(v string) *MoveMaterialFileRequest {
	s.BizId = &v
	return s
}

func (s *MoveMaterialFileRequest) SetDirectoryId(v string) *MoveMaterialFileRequest {
	s.DirectoryId = &v
	return s
}

func (s *MoveMaterialFileRequest) SetFileIds(v []*string) *MoveMaterialFileRequest {
	s.FileIds = v
	return s
}

func (s *MoveMaterialFileRequest) Validate() error {
	return dara.Validate(s)
}
