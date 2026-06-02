// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListWorkspaceCodeResponseBodyData) *ListWorkspaceCodeResponseBody
	GetData() *ListWorkspaceCodeResponseBodyData
	SetErrorCode(v string) *ListWorkspaceCodeResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListWorkspaceCodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListWorkspaceCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWorkspaceCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkspaceCodeResponseBody
	GetSuccess() *bool
}

type ListWorkspaceCodeResponseBody struct {
	Data *ListWorkspaceCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// This record is being collected, please wait for a moment.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkspaceCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceCodeResponseBody) GetData() *ListWorkspaceCodeResponseBodyData {
	return s.Data
}

func (s *ListWorkspaceCodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkspaceCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListWorkspaceCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkspaceCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspaceCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkspaceCodeResponseBody) SetData(v *ListWorkspaceCodeResponseBodyData) *ListWorkspaceCodeResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkspaceCodeResponseBody) SetErrorCode(v string) *ListWorkspaceCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkspaceCodeResponseBody) SetHttpStatusCode(v int32) *ListWorkspaceCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWorkspaceCodeResponseBody) SetMessage(v string) *ListWorkspaceCodeResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkspaceCodeResponseBody) SetRequestId(v string) *ListWorkspaceCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceCodeResponseBody) SetSuccess(v bool) *ListWorkspaceCodeResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkspaceCodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkspaceCodeResponseBodyData struct {
	List []*ListWorkspaceCodeResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListWorkspaceCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkspaceCodeResponseBodyData) GetList() []*ListWorkspaceCodeResponseBodyDataList {
	return s.List
}

func (s *ListWorkspaceCodeResponseBodyData) SetList(v []*ListWorkspaceCodeResponseBodyDataList) *ListWorkspaceCodeResponseBodyData {
	s.List = v
	return s
}

func (s *ListWorkspaceCodeResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceCodeResponseBodyDataList struct {
	// example:
	//
	// true
	IsDir *bool `json:"IsDir,omitempty" xml:"IsDir,omitempty"`
	// example:
	//
	// 2026-01-01T10:11:12Z
	Mtime *string `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 59
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// ../abc.py
	Symlink *string `json:"Symlink,omitempty" xml:"Symlink,omitempty"`
}

func (s ListWorkspaceCodeResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceCodeResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListWorkspaceCodeResponseBodyDataList) GetIsDir() *bool {
	return s.IsDir
}

func (s *ListWorkspaceCodeResponseBodyDataList) GetMtime() *string {
	return s.Mtime
}

func (s *ListWorkspaceCodeResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListWorkspaceCodeResponseBodyDataList) GetSize() *int64 {
	return s.Size
}

func (s *ListWorkspaceCodeResponseBodyDataList) GetSymlink() *string {
	return s.Symlink
}

func (s *ListWorkspaceCodeResponseBodyDataList) SetIsDir(v bool) *ListWorkspaceCodeResponseBodyDataList {
	s.IsDir = &v
	return s
}

func (s *ListWorkspaceCodeResponseBodyDataList) SetMtime(v string) *ListWorkspaceCodeResponseBodyDataList {
	s.Mtime = &v
	return s
}

func (s *ListWorkspaceCodeResponseBodyDataList) SetName(v string) *ListWorkspaceCodeResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListWorkspaceCodeResponseBodyDataList) SetSize(v int64) *ListWorkspaceCodeResponseBodyDataList {
	s.Size = &v
	return s
}

func (s *ListWorkspaceCodeResponseBodyDataList) SetSymlink(v string) *ListWorkspaceCodeResponseBodyDataList {
	s.Symlink = &v
	return s
}

func (s *ListWorkspaceCodeResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
