// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStoryFile interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddStoryFile
	GetErrorCode() *string
	SetErrorMessage(v string) *AddStoryFile
	GetErrorMessage() *string
	SetFileId(v string) *AddStoryFile
	GetFileId() *string
	SetRevisionId(v string) *AddStoryFile
	GetRevisionId() *string
}

type AddStoryFile struct {
	// Error codes when adding a single file.
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// Error message when adding a single file.
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// File version.
	//
	// example:
	//
	// 642a88dd06e49d9c0a14411ebae606f70edd9a59
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s AddStoryFile) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFile) GoString() string {
	return s.String()
}

func (s *AddStoryFile) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddStoryFile) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddStoryFile) GetFileId() *string {
	return s.FileId
}

func (s *AddStoryFile) GetRevisionId() *string {
	return s.RevisionId
}

func (s *AddStoryFile) SetErrorCode(v string) *AddStoryFile {
	s.ErrorCode = &v
	return s
}

func (s *AddStoryFile) SetErrorMessage(v string) *AddStoryFile {
	s.ErrorMessage = &v
	return s
}

func (s *AddStoryFile) SetFileId(v string) *AddStoryFile {
	s.FileId = &v
	return s
}

func (s *AddStoryFile) SetRevisionId(v string) *AddStoryFile {
	s.RevisionId = &v
	return s
}

func (s *AddStoryFile) Validate() error {
	return dara.Validate(s)
}
