// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUploadedPartsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ListUploadedPartsRequest
	GetDriveId() *string
	SetFileId(v string) *ListUploadedPartsRequest
	GetFileId() *string
	SetLimit(v int32) *ListUploadedPartsRequest
	GetLimit() *int32
	SetPartNumberMarker(v int32) *ListUploadedPartsRequest
	GetPartNumberMarker() *int32
	SetShareId(v string) *ListUploadedPartsRequest
	GetShareId() *string
	SetUploadId(v string) *ListUploadedPartsRequest
	GetUploadId() *string
}

type ListUploadedPartsRequest struct {
	// The drive ID. This parameter is required if the file is not uploaded by using the share URL of the file.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 322fb07b975f4b0ae1b543fe8475eee4c19eb2b2
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	PartNumberMarker *int32 `json:"part_number_marker,omitempty" xml:"part_number_marker,omitempty"`
	// The share ID. This parameter is required if the file is uploaded by using the share URL of the file.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The ID of the upload task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00166D06127B413BA1EC8ABB1144D101
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s ListUploadedPartsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUploadedPartsRequest) GoString() string {
	return s.String()
}

func (s *ListUploadedPartsRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ListUploadedPartsRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListUploadedPartsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListUploadedPartsRequest) GetPartNumberMarker() *int32 {
	return s.PartNumberMarker
}

func (s *ListUploadedPartsRequest) GetShareId() *string {
	return s.ShareId
}

func (s *ListUploadedPartsRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *ListUploadedPartsRequest) SetDriveId(v string) *ListUploadedPartsRequest {
	s.DriveId = &v
	return s
}

func (s *ListUploadedPartsRequest) SetFileId(v string) *ListUploadedPartsRequest {
	s.FileId = &v
	return s
}

func (s *ListUploadedPartsRequest) SetLimit(v int32) *ListUploadedPartsRequest {
	s.Limit = &v
	return s
}

func (s *ListUploadedPartsRequest) SetPartNumberMarker(v int32) *ListUploadedPartsRequest {
	s.PartNumberMarker = &v
	return s
}

func (s *ListUploadedPartsRequest) SetShareId(v string) *ListUploadedPartsRequest {
	s.ShareId = &v
	return s
}

func (s *ListUploadedPartsRequest) SetUploadId(v string) *ListUploadedPartsRequest {
	s.UploadId = &v
	return s
}

func (s *ListUploadedPartsRequest) Validate() error {
	return dara.Validate(s)
}
