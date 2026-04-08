// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoldersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListFoldersResponseBodyData) *ListFoldersResponseBody
	GetData() *ListFoldersResponseBodyData
	SetErrorCode(v string) *ListFoldersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListFoldersResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListFoldersResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListFoldersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFoldersResponseBody
	GetSuccess() *bool
}

type ListFoldersResponseBody struct {
	// The path of the folder.
	Data *ListFoldersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The information about the folders.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 0000-ABCD-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of folders.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFoldersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoldersResponseBody) GetData() *ListFoldersResponseBodyData {
	return s.Data
}

func (s *ListFoldersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListFoldersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListFoldersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFoldersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFoldersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFoldersResponseBody) SetData(v *ListFoldersResponseBodyData) *ListFoldersResponseBody {
	s.Data = v
	return s
}

func (s *ListFoldersResponseBody) SetErrorCode(v string) *ListFoldersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListFoldersResponseBody) SetErrorMessage(v string) *ListFoldersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListFoldersResponseBody) SetHttpStatusCode(v int32) *ListFoldersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFoldersResponseBody) SetRequestId(v string) *ListFoldersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFoldersResponseBody) SetSuccess(v bool) *ListFoldersResponseBody {
	s.Success = &v
	return s
}

func (s *ListFoldersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFoldersResponseBodyData struct {
	Folders []*ListFoldersResponseBodyDataFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Repeated"`
	// The folder ID.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// http://100.67.165.184/business/api/folders/list
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ListFolders
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFoldersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFoldersResponseBodyData) GetFolders() []*ListFoldersResponseBodyDataFolders {
	return s.Folders
}

func (s *ListFoldersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFoldersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFoldersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFoldersResponseBodyData) SetFolders(v []*ListFoldersResponseBodyDataFolders) *ListFoldersResponseBodyData {
	s.Folders = v
	return s
}

func (s *ListFoldersResponseBodyData) SetPageNumber(v int32) *ListFoldersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersResponseBodyData) SetPageSize(v int32) *ListFoldersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFoldersResponseBodyData) SetTotalCount(v int32) *ListFoldersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFoldersResponseBodyData) Validate() error {
	if s.Folders != nil {
		for _, item := range s.Folders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFoldersResponseBodyDataFolders struct {
	// example:
	//
	// 2735c2****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	FolderPath *string `json:"FolderPath,omitempty" xml:"FolderPath,omitempty"`
}

func (s ListFoldersResponseBodyDataFolders) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersResponseBodyDataFolders) GoString() string {
	return s.String()
}

func (s *ListFoldersResponseBodyDataFolders) GetFolderId() *string {
	return s.FolderId
}

func (s *ListFoldersResponseBodyDataFolders) GetFolderPath() *string {
	return s.FolderPath
}

func (s *ListFoldersResponseBodyDataFolders) SetFolderId(v string) *ListFoldersResponseBodyDataFolders {
	s.FolderId = &v
	return s
}

func (s *ListFoldersResponseBodyDataFolders) SetFolderPath(v string) *ListFoldersResponseBodyDataFolders {
	s.FolderPath = &v
	return s
}

func (s *ListFoldersResponseBodyDataFolders) Validate() error {
	return dara.Validate(s)
}
