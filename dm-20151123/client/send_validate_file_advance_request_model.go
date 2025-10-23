// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSendValidateFileAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressColumn(v int32) *SendValidateFileAdvanceRequest
	GetAddressColumn() *int32
	SetFileName(v string) *SendValidateFileAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *SendValidateFileAdvanceRequest
	GetFileUrlObject() io.Reader
	SetHasHeaderRow(v bool) *SendValidateFileAdvanceRequest
	GetHasHeaderRow() *bool
	SetRemoveDuplicate(v bool) *SendValidateFileAdvanceRequest
	GetRemoveDuplicate() *bool
}

type SendValidateFileAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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

func (s SendValidateFileAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SendValidateFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SendValidateFileAdvanceRequest) GetAddressColumn() *int32 {
	return s.AddressColumn
}

func (s *SendValidateFileAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SendValidateFileAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SendValidateFileAdvanceRequest) GetHasHeaderRow() *bool {
	return s.HasHeaderRow
}

func (s *SendValidateFileAdvanceRequest) GetRemoveDuplicate() *bool {
	return s.RemoveDuplicate
}

func (s *SendValidateFileAdvanceRequest) SetAddressColumn(v int32) *SendValidateFileAdvanceRequest {
	s.AddressColumn = &v
	return s
}

func (s *SendValidateFileAdvanceRequest) SetFileName(v string) *SendValidateFileAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SendValidateFileAdvanceRequest) SetFileUrlObject(v io.Reader) *SendValidateFileAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SendValidateFileAdvanceRequest) SetHasHeaderRow(v bool) *SendValidateFileAdvanceRequest {
	s.HasHeaderRow = &v
	return s
}

func (s *SendValidateFileAdvanceRequest) SetRemoveDuplicate(v bool) *SendValidateFileAdvanceRequest {
	s.RemoveDuplicate = &v
	return s
}

func (s *SendValidateFileAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
