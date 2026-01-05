// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallFrom(v string) *DeleteFileUploadRequest
	GetCallFrom() *string
	SetDmsUnit(v string) *DeleteFileUploadRequest
	GetDmsUnit() *string
	SetFileId(v string) *DeleteFileUploadRequest
	GetFileId() *string
}

type DeleteFileUploadRequest struct {
	// example:
	//
	// TrailCenter
	CallFrom *string `json:"CallFrom,omitempty" xml:"CallFrom,omitempty"`
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-8*******01m
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s DeleteFileUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileUploadRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileUploadRequest) GetCallFrom() *string {
	return s.CallFrom
}

func (s *DeleteFileUploadRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *DeleteFileUploadRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteFileUploadRequest) SetCallFrom(v string) *DeleteFileUploadRequest {
	s.CallFrom = &v
	return s
}

func (s *DeleteFileUploadRequest) SetDmsUnit(v string) *DeleteFileUploadRequest {
	s.DmsUnit = &v
	return s
}

func (s *DeleteFileUploadRequest) SetFileId(v string) *DeleteFileUploadRequest {
	s.FileId = &v
	return s
}

func (s *DeleteFileUploadRequest) Validate() error {
	return dara.Validate(s)
}
