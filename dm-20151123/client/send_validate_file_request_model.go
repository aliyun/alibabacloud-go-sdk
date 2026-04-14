// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendValidateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressColumn(v int32) *SendValidateFileRequest
	GetAddressColumn() *int32
	SetFileName(v string) *SendValidateFileRequest
	GetFileName() *string
	SetFileUrl(v string) *SendValidateFileRequest
	GetFileUrl() *string
	SetHasHeaderRow(v bool) *SendValidateFileRequest
	GetHasHeaderRow() *bool
	SetRemoveDuplicate(v bool) *SendValidateFileRequest
	GetRemoveDuplicate() *bool
}

type SendValidateFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AddressColumn *int32 `json:"AddressColumn,omitempty" xml:"AddressColumn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxx.oss-xxx.com/file.txt
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	HasHeaderRow *bool `json:"HasHeaderRow,omitempty" xml:"HasHeaderRow,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	RemoveDuplicate *bool `json:"RemoveDuplicate,omitempty" xml:"RemoveDuplicate,omitempty"`
}

func (s SendValidateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SendValidateFileRequest) GoString() string {
	return s.String()
}

func (s *SendValidateFileRequest) GetAddressColumn() *int32 {
	return s.AddressColumn
}

func (s *SendValidateFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *SendValidateFileRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SendValidateFileRequest) GetHasHeaderRow() *bool {
	return s.HasHeaderRow
}

func (s *SendValidateFileRequest) GetRemoveDuplicate() *bool {
	return s.RemoveDuplicate
}

func (s *SendValidateFileRequest) SetAddressColumn(v int32) *SendValidateFileRequest {
	s.AddressColumn = &v
	return s
}

func (s *SendValidateFileRequest) SetFileName(v string) *SendValidateFileRequest {
	s.FileName = &v
	return s
}

func (s *SendValidateFileRequest) SetFileUrl(v string) *SendValidateFileRequest {
	s.FileUrl = &v
	return s
}

func (s *SendValidateFileRequest) SetHasHeaderRow(v bool) *SendValidateFileRequest {
	s.HasHeaderRow = &v
	return s
}

func (s *SendValidateFileRequest) SetRemoveDuplicate(v bool) *SendValidateFileRequest {
	s.RemoveDuplicate = &v
	return s
}

func (s *SendValidateFileRequest) Validate() error {
	return dara.Validate(s)
}
