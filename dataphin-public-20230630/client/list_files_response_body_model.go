// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFilesResponseBody
	GetCode() *string
	SetFileList(v []*ListFilesResponseBodyFileList) *ListFilesResponseBody
	GetFileList() []*ListFilesResponseBodyFileList
	SetHttpStatusCode(v int32) *ListFilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFilesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFilesResponseBody
	GetSuccess() *bool
}

type ListFilesResponseBody struct {
	// example:
	//
	// OK
	Code     *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	FileList []*ListFilesResponseBodyFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFilesResponseBody) GetFileList() []*ListFilesResponseBodyFileList {
	return s.FileList
}

func (s *ListFilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFilesResponseBody) SetCode(v string) *ListFilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFilesResponseBody) SetFileList(v []*ListFilesResponseBodyFileList) *ListFilesResponseBody {
	s.FileList = v
	return s
}

func (s *ListFilesResponseBody) SetHttpStatusCode(v int32) *ListFilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFilesResponseBody) SetMessage(v string) *ListFilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFilesResponseBody) SetRequestId(v string) *ListFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFilesResponseBody) SetSuccess(v bool) *ListFilesResponseBody {
	s.Success = &v
	return s
}

func (s *ListFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFilesResponseBodyFileList struct {
	// example:
	//
	// tempCode
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// select 1;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1212111
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// /xx/x
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// directory
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// 1717483193830
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1717483193830
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 111231112
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1212111
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 312112121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListFilesResponseBodyFileList) String() string {
	return dara.Prettify(s)
}

func (s ListFilesResponseBodyFileList) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBodyFileList) GetCategory() *string {
	return s.Category
}

func (s *ListFilesResponseBodyFileList) GetContent() *string {
	return s.Content
}

func (s *ListFilesResponseBodyFileList) GetCreator() *string {
	return s.Creator
}

func (s *ListFilesResponseBodyFileList) GetDirectory() *string {
	return s.Directory
}

func (s *ListFilesResponseBodyFileList) GetFileType() *string {
	return s.FileType
}

func (s *ListFilesResponseBodyFileList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListFilesResponseBodyFileList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListFilesResponseBodyFileList) GetId() *int64 {
	return s.Id
}

func (s *ListFilesResponseBodyFileList) GetLastModifier() *string {
	return s.LastModifier
}

func (s *ListFilesResponseBodyFileList) GetName() *string {
	return s.Name
}

func (s *ListFilesResponseBodyFileList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListFilesResponseBodyFileList) SetCategory(v string) *ListFilesResponseBodyFileList {
	s.Category = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetContent(v string) *ListFilesResponseBodyFileList {
	s.Content = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetCreator(v string) *ListFilesResponseBodyFileList {
	s.Creator = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetDirectory(v string) *ListFilesResponseBodyFileList {
	s.Directory = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetFileType(v string) *ListFilesResponseBodyFileList {
	s.FileType = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetGmtCreate(v int64) *ListFilesResponseBodyFileList {
	s.GmtCreate = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetGmtModified(v int64) *ListFilesResponseBodyFileList {
	s.GmtModified = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetId(v int64) *ListFilesResponseBodyFileList {
	s.Id = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetLastModifier(v string) *ListFilesResponseBodyFileList {
	s.LastModifier = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetName(v string) *ListFilesResponseBodyFileList {
	s.Name = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetProjectId(v int64) *ListFilesResponseBodyFileList {
	s.ProjectId = &v
	return s
}

func (s *ListFilesResponseBodyFileList) Validate() error {
	return dara.Validate(s)
}
