// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUploadedPartsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *ListUploadedPartsResponseBody
	GetFileId() *string
	SetNextPartNumberMarker(v string) *ListUploadedPartsResponseBody
	GetNextPartNumberMarker() *string
	SetParallelUpload(v bool) *ListUploadedPartsResponseBody
	GetParallelUpload() *bool
	SetUploadId(v string) *ListUploadedPartsResponseBody
	GetUploadId() *string
	SetUploadedParts(v []*UploadPartInfo) *ListUploadedPartsResponseBody
	GetUploadedParts() []*UploadPartInfo
}

type ListUploadedPartsResponseBody struct {
	// The file ID.
	//
	// example:
	//
	// 322fb07b975f4b0ae1b543fe8475eee4c19eb2b2
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextPartNumberMarker *string `json:"next_part_number_marker,omitempty" xml:"next_part_number_marker,omitempty"`
	// Indicates whether the parallel upload feature is enabled.
	//
	// example:
	//
	// false
	ParallelUpload *bool `json:"parallel_upload,omitempty" xml:"parallel_upload,omitempty"`
	// The ID of the upload task.
	//
	// example:
	//
	// 00166D06127B413BA1EC8ABB1144D101
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	// The information about the file parts.
	UploadedParts []*UploadPartInfo `json:"uploaded_parts,omitempty" xml:"uploaded_parts,omitempty" type:"Repeated"`
}

func (s ListUploadedPartsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUploadedPartsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUploadedPartsResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *ListUploadedPartsResponseBody) GetNextPartNumberMarker() *string {
	return s.NextPartNumberMarker
}

func (s *ListUploadedPartsResponseBody) GetParallelUpload() *bool {
	return s.ParallelUpload
}

func (s *ListUploadedPartsResponseBody) GetUploadId() *string {
	return s.UploadId
}

func (s *ListUploadedPartsResponseBody) GetUploadedParts() []*UploadPartInfo {
	return s.UploadedParts
}

func (s *ListUploadedPartsResponseBody) SetFileId(v string) *ListUploadedPartsResponseBody {
	s.FileId = &v
	return s
}

func (s *ListUploadedPartsResponseBody) SetNextPartNumberMarker(v string) *ListUploadedPartsResponseBody {
	s.NextPartNumberMarker = &v
	return s
}

func (s *ListUploadedPartsResponseBody) SetParallelUpload(v bool) *ListUploadedPartsResponseBody {
	s.ParallelUpload = &v
	return s
}

func (s *ListUploadedPartsResponseBody) SetUploadId(v string) *ListUploadedPartsResponseBody {
	s.UploadId = &v
	return s
}

func (s *ListUploadedPartsResponseBody) SetUploadedParts(v []*UploadPartInfo) *ListUploadedPartsResponseBody {
	s.UploadedParts = v
	return s
}

func (s *ListUploadedPartsResponseBody) Validate() error {
	if s.UploadedParts != nil {
		for _, item := range s.UploadedParts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
