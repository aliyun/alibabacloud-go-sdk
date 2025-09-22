// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *FolderItem) *AddFolderResponseBody
	GetData() *FolderItem
	SetRequestId(v string) *AddFolderResponseBody
	GetRequestId() *string
}

type AddFolderResponseBody struct {
	Data *FolderItem `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// E38E561C-B996-0E19-8DBC-A218AAE17FBA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFolderResponseBody) GoString() string {
	return s.String()
}

func (s *AddFolderResponseBody) GetData() *FolderItem {
	return s.Data
}

func (s *AddFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFolderResponseBody) SetData(v *FolderItem) *AddFolderResponseBody {
	s.Data = v
	return s
}

func (s *AddFolderResponseBody) SetRequestId(v string) *AddFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
