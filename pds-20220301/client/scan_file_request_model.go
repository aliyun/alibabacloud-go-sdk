// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ScanFileRequest
	GetDriveId() *string
	SetFields(v string) *ScanFileRequest
	GetFields() *string
	SetLimit(v int32) *ScanFileRequest
	GetLimit() *int32
	SetMarker(v string) *ScanFileRequest
	GetMarker() *string
}

type ScanFileRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file properties to return.
	//
	// 	- If you want to return all file properties, set this parameter to \\*.
	//
	// 	- By default, if you do not specify this parameter, the following properties of a file are returned: - file_id, - drive_id, - parent_file_id, - type, - created_at, - updated_at, - file_extention, - size, - starred, - status, - category, and - permissions.
	//
	// 	- You can also specify properties to return. Separate multiple properties with commas (,).
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ScanFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanFileRequest) GoString() string {
	return s.String()
}

func (s *ScanFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ScanFileRequest) GetFields() *string {
	return s.Fields
}

func (s *ScanFileRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ScanFileRequest) GetMarker() *string {
	return s.Marker
}

func (s *ScanFileRequest) SetDriveId(v string) *ScanFileRequest {
	s.DriveId = &v
	return s
}

func (s *ScanFileRequest) SetFields(v string) *ScanFileRequest {
	s.Fields = &v
	return s
}

func (s *ScanFileRequest) SetLimit(v int32) *ScanFileRequest {
	s.Limit = &v
	return s
}

func (s *ScanFileRequest) SetMarker(v string) *ScanFileRequest {
	s.Marker = &v
	return s
}

func (s *ScanFileRequest) Validate() error {
	return dara.Validate(s)
}
