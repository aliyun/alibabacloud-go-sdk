// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecyclebinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ListRecyclebinRequest
	GetDriveId() *string
	SetFields(v string) *ListRecyclebinRequest
	GetFields() *string
	SetLimit(v int32) *ListRecyclebinRequest
	GetLimit() *int32
	SetMarker(v string) *ListRecyclebinRequest
	GetMarker() *string
	SetThumbnailProcesses(v map[string]*ImageProcess) *ListRecyclebinRequest
	GetThumbnailProcesses() map[string]*ImageProcess
}

type ListRecyclebinRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The fields of an entry (file or folder) to return.
	//
	// 1\\. If you set this parameter to \\*, all fields are returned.
	//
	// 2\\. If you set this parameter to a null value or leave this parameter empty, the fields, such as file creator, file modifier, and custom tags, are not returned.
	//
	// The default value is a null value, which indicates that only some fields are returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 200. Default value: 50.
	//
	// The number of returned entries must be less than or equal to the value of this parameter.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The name of the entry after which the list begins. Entries whose names are alphabetically after the value of this parameter are returned. If you do not specify this parameter, all entries are returned. This parameter is left empty by default.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The thumbnail configurations. Up to five thumbnails can be returned at a time. The value contains key-value pairs. You can customize the keys. The URL of a thumbnail is returned based on the key.
	ThumbnailProcesses map[string]*ImageProcess `json:"thumbnail_processes,omitempty" xml:"thumbnail_processes,omitempty"`
}

func (s ListRecyclebinRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecyclebinRequest) GoString() string {
	return s.String()
}

func (s *ListRecyclebinRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ListRecyclebinRequest) GetFields() *string {
	return s.Fields
}

func (s *ListRecyclebinRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListRecyclebinRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListRecyclebinRequest) GetThumbnailProcesses() map[string]*ImageProcess {
	return s.ThumbnailProcesses
}

func (s *ListRecyclebinRequest) SetDriveId(v string) *ListRecyclebinRequest {
	s.DriveId = &v
	return s
}

func (s *ListRecyclebinRequest) SetFields(v string) *ListRecyclebinRequest {
	s.Fields = &v
	return s
}

func (s *ListRecyclebinRequest) SetLimit(v int32) *ListRecyclebinRequest {
	s.Limit = &v
	return s
}

func (s *ListRecyclebinRequest) SetMarker(v string) *ListRecyclebinRequest {
	s.Marker = &v
	return s
}

func (s *ListRecyclebinRequest) SetThumbnailProcesses(v map[string]*ImageProcess) *ListRecyclebinRequest {
	s.ThumbnailProcesses = v
	return s
}

func (s *ListRecyclebinRequest) Validate() error {
	return dara.Validate(s)
}
