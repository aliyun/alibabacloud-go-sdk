// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDatasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *DescribeComfyUserDatasRequest
	GetFileName() *string
	SetPageNumber(v int32) *DescribeComfyUserDatasRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeComfyUserDatasRequest
	GetPageSize() *int32
}

type DescribeComfyUserDatasRequest struct {
	// example:
	//
	// myfile
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeComfyUserDatasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDatasRequest) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDatasRequest) GetFileName() *string {
	return s.FileName
}

func (s *DescribeComfyUserDatasRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeComfyUserDatasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeComfyUserDatasRequest) SetFileName(v string) *DescribeComfyUserDatasRequest {
	s.FileName = &v
	return s
}

func (s *DescribeComfyUserDatasRequest) SetPageNumber(v int32) *DescribeComfyUserDatasRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeComfyUserDatasRequest) SetPageSize(v int32) *DescribeComfyUserDatasRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeComfyUserDatasRequest) Validate() error {
	return dara.Validate(s)
}
