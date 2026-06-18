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
	// An array of field names used to filter non-private fields (those not prefixed with an underscore _) in the Metadata field returned by this operation. By default, Fields is empty, and all non-private fields in Metadata are returned. To return only specific non-private fields in Metadata, such as title, pass title in this parameter.
	//
	// Default value: empty.
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// <props="china">
	//
	// The file ID, which is the `FileId` returned by the **AddFile*	- operation. This field is not required for data query or image Q&A knowledge bases. This field is required for document search or audio/video search knowledge bases. You can also obtain the file ID by clicking the ID icon next to the file name on the Files tab of [Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center). You can use the file ID to filter the returned chunks. Default value: empty.
	//
	//
	//
	//
	// <props="intl">
	//
	// The file ID, which is the `FileId` returned by the **AddFile*	- operation. This field is not required for data query or image Q&A knowledge bases. This field is required for document search knowledge bases. You can also obtain the file ID by clicking the ID icon next to the file name on the Files tab of
	//
	// [Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). You can use the file ID to filter the returned chunks. Default value: empty.
	//
	// .
	//
	// example:
	//
	// file_5f03dfea56da4050ab68d61871fc4cb3_xxxxxxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file ID field in the legacy Model Studio SDK. The usage and default value are identical to those of the `FileId` field. If you are using the following versions (or later) of the Model Studio SDK, use the `FileId` field instead. If you are using the SWIFT Model Studio SDK, continue to use this field.
	//
	// - Java (async): 1.0.18
	//
	// - Java: 1.10.2
	//
	// - TypeScript: 1.10.2
	//
	// - Go: 1.10.2
	//
	// - PHP: 1.10.2
	//
	// - Python: 1.10.2
	//
	// - C#: 1.10.2
	//
	// - C++: 1.10.17
	//
	// > **How to check the Model Studio SDK version**: Visit the <props="china">[Model Studio SDK center](https://api.aliyun.com/api-tools/sdk/bailian?version=2023-12-29)<props="intl">[Model Studio SDK center](https://api.alibabacloud.com/api-tools/sdk/bailian?version=2023-12-29), click "**Install**" in the left-side navigation pane, set the API version to "**2023-12-29**", select your programming language, and then click "**History Versions**" to view the version.
	//
	// example:
	//
	// file_5f03dfea56da4050ab68d61871fc4cb3_xxxxxxxx
	Filed *string `json:"Filed,omitempty" xml:"Filed,omitempty"`
	// The knowledge base ID, which is the `Data.Id` returned by the **CreateIndex*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// otoru9xxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The page number to query. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of text chunks to display per page in a paged query. Maximum value: 100. Default value: 10.
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
