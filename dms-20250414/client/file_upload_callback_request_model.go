// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileUploadCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallFrom(v string) *FileUploadCallbackRequest
	GetCallFrom() *string
	SetDmsUnit(v string) *FileUploadCallbackRequest
	GetDmsUnit() *string
	SetFileSize(v int64) *FileUploadCallbackRequest
	GetFileSize() *int64
	SetFilename(v string) *FileUploadCallbackRequest
	GetFilename() *string
	SetUploadLocation(v string) *FileUploadCallbackRequest
	GetUploadLocation() *string
}

type FileUploadCallbackRequest struct {
	// example:
	//
	// TrailCenter
	CallFrom *string `json:"CallFrom,omitempty" xml:"CallFrom,omitempty"`
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// example:
	//
	// 8110
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// conversion_metrics.csv
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// This parameter is required.
	UploadLocation *string `json:"UploadLocation,omitempty" xml:"UploadLocation,omitempty"`
}

func (s FileUploadCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s FileUploadCallbackRequest) GoString() string {
	return s.String()
}

func (s *FileUploadCallbackRequest) GetCallFrom() *string {
	return s.CallFrom
}

func (s *FileUploadCallbackRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *FileUploadCallbackRequest) GetFileSize() *int64 {
	return s.FileSize
}

func (s *FileUploadCallbackRequest) GetFilename() *string {
	return s.Filename
}

func (s *FileUploadCallbackRequest) GetUploadLocation() *string {
	return s.UploadLocation
}

func (s *FileUploadCallbackRequest) SetCallFrom(v string) *FileUploadCallbackRequest {
	s.CallFrom = &v
	return s
}

func (s *FileUploadCallbackRequest) SetDmsUnit(v string) *FileUploadCallbackRequest {
	s.DmsUnit = &v
	return s
}

func (s *FileUploadCallbackRequest) SetFileSize(v int64) *FileUploadCallbackRequest {
	s.FileSize = &v
	return s
}

func (s *FileUploadCallbackRequest) SetFilename(v string) *FileUploadCallbackRequest {
	s.Filename = &v
	return s
}

func (s *FileUploadCallbackRequest) SetUploadLocation(v string) *FileUploadCallbackRequest {
	s.UploadLocation = &v
	return s
}

func (s *FileUploadCallbackRequest) Validate() error {
	return dara.Validate(s)
}
