// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRdTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRdTreeResponseBodyData) *GetRdTreeResponseBody
	GetData() *GetRdTreeResponseBodyData
	SetRequestId(v string) *GetRdTreeResponseBody
	GetRequestId() *string
}

type GetRdTreeResponseBody struct {
	// The processing result.
	Data *GetRdTreeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRdTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRdTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRdTreeResponseBody) GetData() *GetRdTreeResponseBodyData {
	return s.Data
}

func (s *GetRdTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRdTreeResponseBody) SetData(v *GetRdTreeResponseBodyData) *GetRdTreeResponseBody {
	s.Data = v
	return s
}

func (s *GetRdTreeResponseBody) SetRequestId(v string) *GetRdTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRdTreeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRdTreeResponseBodyData struct {
	// The subfolder.
	Children []interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The ID of the folder in the resource directory.
	//
	// example:
	//
	// fd-CGA73I****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// Root
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
}

func (s GetRdTreeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRdTreeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRdTreeResponseBodyData) GetChildren() []interface{} {
	return s.Children
}

func (s *GetRdTreeResponseBodyData) GetFolderId() *string {
	return s.FolderId
}

func (s *GetRdTreeResponseBodyData) GetFolderName() *string {
	return s.FolderName
}

func (s *GetRdTreeResponseBodyData) SetChildren(v []interface{}) *GetRdTreeResponseBodyData {
	s.Children = v
	return s
}

func (s *GetRdTreeResponseBodyData) SetFolderId(v string) *GetRdTreeResponseBodyData {
	s.FolderId = &v
	return s
}

func (s *GetRdTreeResponseBodyData) SetFolderName(v string) *GetRdTreeResponseBodyData {
	s.FolderName = &v
	return s
}

func (s *GetRdTreeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
