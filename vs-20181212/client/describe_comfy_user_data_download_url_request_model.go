// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDataDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *DescribeComfyUserDataDownloadUrlRequest
	GetFileName() *string
}

type DescribeComfyUserDataDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// myfile
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s DescribeComfyUserDataDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDataDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDataDownloadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *DescribeComfyUserDataDownloadUrlRequest) SetFileName(v string) *DescribeComfyUserDataDownloadUrlRequest {
	s.FileName = &v
	return s
}

func (s *DescribeComfyUserDataDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
