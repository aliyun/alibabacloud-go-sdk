// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseAlbumFileOperationResult interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *BaseAlbumFileOperationResult
	GetErrorCode() *string
	SetErrorMessage(v string) *BaseAlbumFileOperationResult
	GetErrorMessage() *string
	SetFile(v *CommonFileItem) *BaseAlbumFileOperationResult
	GetFile() *CommonFileItem
	SetIsSucceed(v bool) *BaseAlbumFileOperationResult
	GetIsSucceed() *bool
}

type BaseAlbumFileOperationResult struct {
	ErrorCode    *string         `json:"error_code,omitempty" xml:"error_code,omitempty"`
	ErrorMessage *string         `json:"error_message,omitempty" xml:"error_message,omitempty"`
	File         *CommonFileItem `json:"file,omitempty" xml:"file,omitempty"`
	IsSucceed    *bool           `json:"is_succeed,omitempty" xml:"is_succeed,omitempty"`
}

func (s BaseAlbumFileOperationResult) String() string {
	return dara.Prettify(s)
}

func (s BaseAlbumFileOperationResult) GoString() string {
	return s.String()
}

func (s *BaseAlbumFileOperationResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BaseAlbumFileOperationResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BaseAlbumFileOperationResult) GetFile() *CommonFileItem {
	return s.File
}

func (s *BaseAlbumFileOperationResult) GetIsSucceed() *bool {
	return s.IsSucceed
}

func (s *BaseAlbumFileOperationResult) SetErrorCode(v string) *BaseAlbumFileOperationResult {
	s.ErrorCode = &v
	return s
}

func (s *BaseAlbumFileOperationResult) SetErrorMessage(v string) *BaseAlbumFileOperationResult {
	s.ErrorMessage = &v
	return s
}

func (s *BaseAlbumFileOperationResult) SetFile(v *CommonFileItem) *BaseAlbumFileOperationResult {
	s.File = v
	return s
}

func (s *BaseAlbumFileOperationResult) SetIsSucceed(v bool) *BaseAlbumFileOperationResult {
	s.IsSucceed = &v
	return s
}

func (s *BaseAlbumFileOperationResult) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}
