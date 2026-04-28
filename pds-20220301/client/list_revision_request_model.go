// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ListRevisionRequest
	GetDriveId() *string
	SetFields(v string) *ListRevisionRequest
	GetFields() *string
	SetFileId(v string) *ListRevisionRequest
	GetFileId() *string
	SetLimit(v int64) *ListRevisionRequest
	GetLimit() *int64
	SetMarker(v string) *ListRevisionRequest
	GetMarker() *string
}

type ListRevisionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Specifies the returned fields.
	//
	// By default, this parameter is left empty. If you set this parameter to \\*, all fields are returned. If you leave this parameter empty, the creator of the file is not returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 100
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 40CB7794C929
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRevisionRequest) GoString() string {
	return s.String()
}

func (s *ListRevisionRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ListRevisionRequest) GetFields() *string {
	return s.Fields
}

func (s *ListRevisionRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListRevisionRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ListRevisionRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListRevisionRequest) SetDriveId(v string) *ListRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *ListRevisionRequest) SetFields(v string) *ListRevisionRequest {
	s.Fields = &v
	return s
}

func (s *ListRevisionRequest) SetFileId(v string) *ListRevisionRequest {
	s.FileId = &v
	return s
}

func (s *ListRevisionRequest) SetLimit(v int64) *ListRevisionRequest {
	s.Limit = &v
	return s
}

func (s *ListRevisionRequest) SetMarker(v string) *ListRevisionRequest {
	s.Marker = &v
	return s
}

func (s *ListRevisionRequest) Validate() error {
	return dara.Validate(s)
}
