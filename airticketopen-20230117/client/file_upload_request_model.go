// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileContent(v string) *FileUploadRequest
	GetFileContent() *string
	SetOrderNum(v int64) *FileUploadRequest
	GetOrderNum() *int64
}

type FileUploadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0r2LSuIsHlxEoGZcnGe34U1njBOR83Q4HNSvMDGrDPK5J71VjcGdRIWz2x3+tFxvQaduwHB46Z9K
	//
	// dbIoDN8xPQ5PWlky8rKOPmAqSZfIRyPmAwvPvTJFwr8bRgHPPaq2VO8kHJ6jFIpJJ5I7Zqd1BjGS
	//
	// SR/kULQZHsDDd2zgA9RRTsEQF2OSxFFFx2P/2Q==
	FileContent *string `json:"file_content,omitempty" xml:"file_content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s FileUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s FileUploadRequest) GoString() string {
	return s.String()
}

func (s *FileUploadRequest) GetFileContent() *string {
	return s.FileContent
}

func (s *FileUploadRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *FileUploadRequest) SetFileContent(v string) *FileUploadRequest {
	s.FileContent = &v
	return s
}

func (s *FileUploadRequest) SetOrderNum(v int64) *FileUploadRequest {
	s.OrderNum = &v
	return s
}

func (s *FileUploadRequest) Validate() error {
	return dara.Validate(s)
}
