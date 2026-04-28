// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeltaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v string) *ListDeltaResponseBody
	GetCursor() *string
	SetHasMore(v bool) *ListDeltaResponseBody
	GetHasMore() *bool
	SetItems(v []*ListDeltaResponseBodyItems) *ListDeltaResponseBody
	GetItems() []*ListDeltaResponseBodyItems
}

type ListDeltaResponseBody struct {
	// The cursor of the incremental information.
	//
	// example:
	//
	// 1WQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// Indicates whether more information is returned.
	//
	// example:
	//
	// true
	HasMore *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	// The incremental information returned.
	Items []*ListDeltaResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListDeltaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeltaResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeltaResponseBody) GetCursor() *string {
	return s.Cursor
}

func (s *ListDeltaResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListDeltaResponseBody) GetItems() []*ListDeltaResponseBodyItems {
	return s.Items
}

func (s *ListDeltaResponseBody) SetCursor(v string) *ListDeltaResponseBody {
	s.Cursor = &v
	return s
}

func (s *ListDeltaResponseBody) SetHasMore(v bool) *ListDeltaResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListDeltaResponseBody) SetItems(v []*ListDeltaResponseBodyItems) *ListDeltaResponseBody {
	s.Items = v
	return s
}

func (s *ListDeltaResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeltaResponseBodyItems struct {
	// The information about the file.
	File *File `json:"file,omitempty" xml:"file,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 122fb09598ae66777c7040109a16f49381f6abe2
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The operation that is performed. Valid values: Valid values:
	//
	// 	- create
	//
	// 	- overwrite
	//
	// 	- delete
	//
	// 	- update
	//
	// 	- move
	//
	// 	- trash
	//
	// 	- restore
	//
	// 	- rename
	//
	// example:
	//
	// create
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
}

func (s ListDeltaResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListDeltaResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListDeltaResponseBodyItems) GetFile() *File {
	return s.File
}

func (s *ListDeltaResponseBodyItems) GetFileId() *string {
	return s.FileId
}

func (s *ListDeltaResponseBodyItems) GetOp() *string {
	return s.Op
}

func (s *ListDeltaResponseBodyItems) SetFile(v *File) *ListDeltaResponseBodyItems {
	s.File = v
	return s
}

func (s *ListDeltaResponseBodyItems) SetFileId(v string) *ListDeltaResponseBodyItems {
	s.FileId = &v
	return s
}

func (s *ListDeltaResponseBodyItems) SetOp(v string) *ListDeltaResponseBodyItems {
	s.Op = &v
	return s
}

func (s *ListDeltaResponseBodyItems) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}
