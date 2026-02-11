// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *DescribeFileDownloadUrlRequest
	GetFileId() *int64
	SetProjectId(v int64) *DescribeFileDownloadUrlRequest
	GetProjectId() *int64
}

type DescribeFileDownloadUrlRequest struct {
	// File ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeFileDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileDownloadUrlRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *DescribeFileDownloadUrlRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DescribeFileDownloadUrlRequest) SetFileId(v int64) *DescribeFileDownloadUrlRequest {
	s.FileId = &v
	return s
}

func (s *DescribeFileDownloadUrlRequest) SetProjectId(v int64) *DescribeFileDownloadUrlRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFileDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
