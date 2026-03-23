// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagMetaAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetaParentId(v string) *ListTagMetaAssetRequest
	GetMetaParentId() *string
	SetMetaType(v string) *ListTagMetaAssetRequest
	GetMetaType() *string
	SetPageNumber(v int32) *ListTagMetaAssetRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagMetaAssetRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListTagMetaAssetRequest
	GetSearchKey() *string
	SetTagName(v string) *ListTagMetaAssetRequest
	GetTagName() *string
	SetTid(v int64) *ListTagMetaAssetRequest
	GetTid() *int64
}

type ListTagMetaAssetRequest struct {
	// example:
	//
	// 123456
	MetaParentId *string `json:"MetaParentId,omitempty" xml:"MetaParentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// META_DATABASE
	MetaType *string `json:"MetaType,omitempty" xml:"MetaType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sys::DMS-DA::cn-hangzhou::space:abcde
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTagMetaAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagMetaAssetRequest) GoString() string {
	return s.String()
}

func (s *ListTagMetaAssetRequest) GetMetaParentId() *string {
	return s.MetaParentId
}

func (s *ListTagMetaAssetRequest) GetMetaType() *string {
	return s.MetaType
}

func (s *ListTagMetaAssetRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagMetaAssetRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagMetaAssetRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListTagMetaAssetRequest) GetTagName() *string {
	return s.TagName
}

func (s *ListTagMetaAssetRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTagMetaAssetRequest) SetMetaParentId(v string) *ListTagMetaAssetRequest {
	s.MetaParentId = &v
	return s
}

func (s *ListTagMetaAssetRequest) SetMetaType(v string) *ListTagMetaAssetRequest {
	s.MetaType = &v
	return s
}

func (s *ListTagMetaAssetRequest) SetPageNumber(v int32) *ListTagMetaAssetRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagMetaAssetRequest) SetPageSize(v int32) *ListTagMetaAssetRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagMetaAssetRequest) SetSearchKey(v string) *ListTagMetaAssetRequest {
	s.SearchKey = &v
	return s
}

func (s *ListTagMetaAssetRequest) SetTagName(v string) *ListTagMetaAssetRequest {
	s.TagName = &v
	return s
}

func (s *ListTagMetaAssetRequest) SetTid(v int64) *ListTagMetaAssetRequest {
	s.Tid = &v
	return s
}

func (s *ListTagMetaAssetRequest) Validate() error {
	return dara.Validate(s)
}
