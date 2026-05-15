// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFileTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileInfos(v []*BatchUpdateFileTagRequestFileInfos) *BatchUpdateFileTagRequest
	GetFileInfos() []*BatchUpdateFileTagRequestFileInfos
	SetUpdateMode(v string) *BatchUpdateFileTagRequest
	GetUpdateMode() *string
}

type BatchUpdateFileTagRequest struct {
	// This parameter is required.
	FileInfos []*BatchUpdateFileTagRequestFileInfos `json:"FileInfos,omitempty" xml:"FileInfos,omitempty" type:"Repeated"`
	// example:
	//
	// OVERWRITE
	UpdateMode *string `json:"UpdateMode,omitempty" xml:"UpdateMode,omitempty"`
}

func (s BatchUpdateFileTagRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileTagRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileTagRequest) GetFileInfos() []*BatchUpdateFileTagRequestFileInfos {
	return s.FileInfos
}

func (s *BatchUpdateFileTagRequest) GetUpdateMode() *string {
	return s.UpdateMode
}

func (s *BatchUpdateFileTagRequest) SetFileInfos(v []*BatchUpdateFileTagRequestFileInfos) *BatchUpdateFileTagRequest {
	s.FileInfos = v
	return s
}

func (s *BatchUpdateFileTagRequest) SetUpdateMode(v string) *BatchUpdateFileTagRequest {
	s.UpdateMode = &v
	return s
}

func (s *BatchUpdateFileTagRequest) Validate() error {
	if s.FileInfos != nil {
		for _, item := range s.FileInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchUpdateFileTagRequestFileInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// file_3d5319366e2c46309f4c11cfbeacd5fd_10045951
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s BatchUpdateFileTagRequestFileInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileTagRequestFileInfos) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileTagRequestFileInfos) GetFileId() *string {
	return s.FileId
}

func (s *BatchUpdateFileTagRequestFileInfos) GetTags() []*string {
	return s.Tags
}

func (s *BatchUpdateFileTagRequestFileInfos) SetFileId(v string) *BatchUpdateFileTagRequestFileInfos {
	s.FileId = &v
	return s
}

func (s *BatchUpdateFileTagRequestFileInfos) SetTags(v []*string) *BatchUpdateFileTagRequestFileInfos {
	s.Tags = v
	return s
}

func (s *BatchUpdateFileTagRequestFileInfos) Validate() error {
	return dara.Validate(s)
}
