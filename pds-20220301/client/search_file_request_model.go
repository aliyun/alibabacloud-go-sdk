// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *SearchFileRequest
	GetDriveId() *string
	SetFields(v string) *SearchFileRequest
	GetFields() *string
	SetLimit(v int32) *SearchFileRequest
	GetLimit() *int32
	SetMarker(v string) *SearchFileRequest
	GetMarker() *string
	SetOrderBy(v string) *SearchFileRequest
	GetOrderBy() *string
	SetQuery(v string) *SearchFileRequest
	GetQuery() *string
	SetRecursive(v bool) *SearchFileRequest
	GetRecursive() *bool
	SetReturnTotalCount(v bool) *SearchFileRequest
	GetReturnTotalCount() *bool
	SetThumbnailProcesses(v map[string]*ImageProcess) *SearchFileRequest
	GetThumbnailProcesses() map[string]*ImageProcess
}

type SearchFileRequest struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The field that is used to return additional information about files. Valid values:
	//
	// 	- dir_size: returns the statistics on each subfolder in the response.
	//
	// 	- id_path: returns the id_path value of each child subject in the response.
	//
	// 	- name_path: returns the name_path value of each child subject in the response.
	//
	// You can specify multiple fields by separating them with commas (,). Example: "id_path,name_path,dir_size".
	//
	// example:
	//
	// url,thumbnail
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 100.
	//
	// The number of returned entries must be less than or equal to the value of this parameter.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The name of the entry after which the list begins. Entries whose names are alphabetically after the value of this parameter are returned. If you do not specify this parameter, all entries are returned.\\
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The field by which to sort the returned entries. Default value: created_at. Valid values:
	//
	// 	- created_at: sorts the entries by creation time.
	//
	// 	- updated_at: sorts the entries by update time.
	//
	// 	- size: sorts the entries by file size.
	//
	// 	- name: sorts the entries by file name.
	//
	// The order in which you want to sort the returned entries. Valid values:
	//
	// 	- ASC: ascending order
	//
	// 	- DESC: descending order
	//
	// You must specify this parameter in the \\<field> \\<ASC or DESC> format. Separate multiple fields with commas (,). A preceding field has a higher priority than a following field. Examples:
	//
	// 	- If you want to sort the entries by file name in ascending order, set this parameter to "name ASC".
	//
	// 	- If you want to sort the entries by creation time in descending order, set this parameter to "created_at DESC".
	//
	// 	- If you want to sort the entries by creation time in descending order and sort the entries by file name in ascending order in case of the same creation time, set this parameter to "created_at DESC,name ASC".
	//
	// example:
	//
	// name
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// The search conditions. Fuzzy searches based on the file name or directory name are supported. The value of this parameter can be up to 4,096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// not name=123
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Specifies whether to perform recursive search on a folder that is specified by setting parent_file_id in the query parameter.
	//
	// example:
	//
	// true
	Recursive *bool `json:"recursive,omitempty" xml:"recursive,omitempty"`
	// Specifies whether to return the total number of retrieved files.
	//
	// example:
	//
	// true
	ReturnTotalCount *bool `json:"return_total_count,omitempty" xml:"return_total_count,omitempty"`
	// The thumbnail configurations. Up to five thumbnails can be returned at a time. The value contains key-value pairs. You can customize the keys. The URL of a thumbnail is returned based on the key.
	ThumbnailProcesses map[string]*ImageProcess `json:"thumbnail_processes,omitempty" xml:"thumbnail_processes,omitempty"`
}

func (s SearchFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFileRequest) GoString() string {
	return s.String()
}

func (s *SearchFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *SearchFileRequest) GetFields() *string {
	return s.Fields
}

func (s *SearchFileRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchFileRequest) GetMarker() *string {
	return s.Marker
}

func (s *SearchFileRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchFileRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchFileRequest) GetRecursive() *bool {
	return s.Recursive
}

func (s *SearchFileRequest) GetReturnTotalCount() *bool {
	return s.ReturnTotalCount
}

func (s *SearchFileRequest) GetThumbnailProcesses() map[string]*ImageProcess {
	return s.ThumbnailProcesses
}

func (s *SearchFileRequest) SetDriveId(v string) *SearchFileRequest {
	s.DriveId = &v
	return s
}

func (s *SearchFileRequest) SetFields(v string) *SearchFileRequest {
	s.Fields = &v
	return s
}

func (s *SearchFileRequest) SetLimit(v int32) *SearchFileRequest {
	s.Limit = &v
	return s
}

func (s *SearchFileRequest) SetMarker(v string) *SearchFileRequest {
	s.Marker = &v
	return s
}

func (s *SearchFileRequest) SetOrderBy(v string) *SearchFileRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchFileRequest) SetQuery(v string) *SearchFileRequest {
	s.Query = &v
	return s
}

func (s *SearchFileRequest) SetRecursive(v bool) *SearchFileRequest {
	s.Recursive = &v
	return s
}

func (s *SearchFileRequest) SetReturnTotalCount(v bool) *SearchFileRequest {
	s.ReturnTotalCount = &v
	return s
}

func (s *SearchFileRequest) SetThumbnailProcesses(v map[string]*ImageProcess) *SearchFileRequest {
	s.ThumbnailProcesses = v
	return s
}

func (s *SearchFileRequest) Validate() error {
	return dara.Validate(s)
}
