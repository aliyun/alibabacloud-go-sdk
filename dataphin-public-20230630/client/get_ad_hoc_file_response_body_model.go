// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAdHocFileResponseBody
	GetCode() *string
	SetFileInfo(v *GetAdHocFileResponseBodyFileInfo) *GetAdHocFileResponseBody
	GetFileInfo() *GetAdHocFileResponseBodyFileInfo
	SetHttpStatusCode(v int32) *GetAdHocFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAdHocFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAdHocFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAdHocFileResponseBody
	GetSuccess() *bool
}

type GetAdHocFileResponseBody struct {
	// example:
	//
	// OK
	Code     *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	FileInfo *GetAdHocFileResponseBodyFileInfo `json:"FileInfo,omitempty" xml:"FileInfo,omitempty" type:"Struct"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAdHocFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdHocFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAdHocFileResponseBody) GetFileInfo() *GetAdHocFileResponseBodyFileInfo {
	return s.FileInfo
}

func (s *GetAdHocFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAdHocFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAdHocFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdHocFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAdHocFileResponseBody) SetCode(v string) *GetAdHocFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetAdHocFileResponseBody) SetFileInfo(v *GetAdHocFileResponseBodyFileInfo) *GetAdHocFileResponseBody {
	s.FileInfo = v
	return s
}

func (s *GetAdHocFileResponseBody) SetHttpStatusCode(v int32) *GetAdHocFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAdHocFileResponseBody) SetMessage(v string) *GetAdHocFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetAdHocFileResponseBody) SetRequestId(v string) *GetAdHocFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdHocFileResponseBody) SetSuccess(v bool) *GetAdHocFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetAdHocFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAdHocFileResponseBodyFileInfo struct {
	// example:
	//
	// select 1;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 12121
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// /xx1/xx2/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// 12121111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 12121
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12132323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetAdHocFileResponseBodyFileInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocFileResponseBodyFileInfo) GoString() string {
	return s.String()
}

func (s *GetAdHocFileResponseBodyFileInfo) GetContent() *string {
	return s.Content
}

func (s *GetAdHocFileResponseBodyFileInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetAdHocFileResponseBodyFileInfo) GetDirectory() *string {
	return s.Directory
}

func (s *GetAdHocFileResponseBodyFileInfo) GetId() *int64 {
	return s.Id
}

func (s *GetAdHocFileResponseBodyFileInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetAdHocFileResponseBodyFileInfo) GetName() *string {
	return s.Name
}

func (s *GetAdHocFileResponseBodyFileInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetAdHocFileResponseBodyFileInfo) SetContent(v string) *GetAdHocFileResponseBodyFileInfo {
	s.Content = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetCreator(v string) *GetAdHocFileResponseBodyFileInfo {
	s.Creator = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetDirectory(v string) *GetAdHocFileResponseBodyFileInfo {
	s.Directory = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetId(v int64) *GetAdHocFileResponseBodyFileInfo {
	s.Id = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetLastModifier(v string) *GetAdHocFileResponseBodyFileInfo {
	s.LastModifier = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetName(v string) *GetAdHocFileResponseBodyFileInfo {
	s.Name = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetProjectId(v int64) *GetAdHocFileResponseBodyFileInfo {
	s.ProjectId = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) Validate() error {
	return dara.Validate(s)
}
