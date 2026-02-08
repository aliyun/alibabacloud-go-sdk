// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveMaterialDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *MoveMaterialDirectoryRequest
	GetBizId() *string
	SetDirectoryId(v string) *MoveMaterialDirectoryRequest
	GetDirectoryId() *string
	SetParentDirectoryId(v string) *MoveMaterialDirectoryRequest
	GetParentDirectoryId() *string
	SetSortNum(v int32) *MoveMaterialDirectoryRequest
	GetSortNum() *int32
}

type MoveMaterialDirectoryRequest struct {
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
	//
	// example:
	//
	// 58157a0a-769a-4364-bbdc-a0e2cf3d5a2
	ParentDirectoryId *string `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SortNum *int32 `json:"SortNum,omitempty" xml:"SortNum,omitempty"`
}

func (s MoveMaterialDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveMaterialDirectoryRequest) GoString() string {
	return s.String()
}

func (s *MoveMaterialDirectoryRequest) GetBizId() *string {
	return s.BizId
}

func (s *MoveMaterialDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *MoveMaterialDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *MoveMaterialDirectoryRequest) GetSortNum() *int32 {
	return s.SortNum
}

func (s *MoveMaterialDirectoryRequest) SetBizId(v string) *MoveMaterialDirectoryRequest {
	s.BizId = &v
	return s
}

func (s *MoveMaterialDirectoryRequest) SetDirectoryId(v string) *MoveMaterialDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *MoveMaterialDirectoryRequest) SetParentDirectoryId(v string) *MoveMaterialDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *MoveMaterialDirectoryRequest) SetSortNum(v int32) *MoveMaterialDirectoryRequest {
	s.SortNum = &v
	return s
}

func (s *MoveMaterialDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
