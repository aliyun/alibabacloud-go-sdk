// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v []*string) *ListChunksRequest
	GetFields() []*string
	SetFileId(v string) *ListChunksRequest
	GetFileId() *string
	SetFiled(v string) *ListChunksRequest
	GetFiled() *string
	SetIndexId(v string) *ListChunksRequest
	GetIndexId() *string
	SetPageNum(v int32) *ListChunksRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListChunksRequest
	GetPageSize() *int32
}

type ListChunksRequest struct {
	// An array of field names. This parameter is used to filter non-private fields (prefixed with_underscores) in the Metadata parameter returned by this operation. By default, this parameter is left empty, which means all non-private fields in the Metadata parameter are returned. If you only want specified non-private fields, such as title, set this parameter to title.
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	FileId *string   `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The primary key ID of the document. This parameter is not required for structured knowledge base, but is required for unstructured knowledge base. To view the ID, you can click the ID icon next to the file name on the [Data Management](https://bailian.console.aliyun.com/#/data-center) page. You can filter returned chunks by the document ID. This parameter is left empty by default.
	//
	// example:
	//
	// file_5f03dfea56da4050ab68d61871fc4cb3_10151493
	Filed *string `json:"Filed,omitempty" xml:"Filed,omitempty"`
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// otoru9en4v
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The number of the pages to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of chunks to display on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChunksRequest) GoString() string {
	return s.String()
}

func (s *ListChunksRequest) GetFields() []*string {
	return s.Fields
}

func (s *ListChunksRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListChunksRequest) GetFiled() *string {
	return s.Filed
}

func (s *ListChunksRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *ListChunksRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListChunksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChunksRequest) SetFields(v []*string) *ListChunksRequest {
	s.Fields = v
	return s
}

func (s *ListChunksRequest) SetFileId(v string) *ListChunksRequest {
	s.FileId = &v
	return s
}

func (s *ListChunksRequest) SetFiled(v string) *ListChunksRequest {
	s.Filed = &v
	return s
}

func (s *ListChunksRequest) SetIndexId(v string) *ListChunksRequest {
	s.IndexId = &v
	return s
}

func (s *ListChunksRequest) SetPageNum(v int32) *ListChunksRequest {
	s.PageNum = &v
	return s
}

func (s *ListChunksRequest) SetPageSize(v int32) *ListChunksRequest {
	s.PageSize = &v
	return s
}

func (s *ListChunksRequest) Validate() error {
	return dara.Validate(s)
}
