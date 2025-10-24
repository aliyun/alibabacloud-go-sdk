// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBucketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *DescribeBucketsRequest
	GetFileType() *string
}

type DescribeBucketsRequest struct {
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s DescribeBucketsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBucketsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBucketsRequest) GetFileType() *string {
	return s.FileType
}

func (s *DescribeBucketsRequest) SetFileType(v string) *DescribeBucketsRequest {
	s.FileType = &v
	return s
}

func (s *DescribeBucketsRequest) Validate() error {
	return dara.Validate(s)
}
