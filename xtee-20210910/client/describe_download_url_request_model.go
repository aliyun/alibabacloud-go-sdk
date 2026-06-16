// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *DescribeDownloadUrlRequest
	GetFileType() *string
	SetLang(v string) *DescribeDownloadUrlRequest
	GetLang() *string
	SetRegId(v string) *DescribeDownloadUrlRequest
	GetRegId() *string
	SetTaskId(v int64) *DescribeDownloadUrlRequest
	GetTaskId() *int64
	SetType(v string) *DescribeDownloadUrlRequest
	GetType() *string
}

type DescribeDownloadUrlRequest struct {
	// The file type.
	//
	// example:
	//
	// CSV
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 18191
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The download type.
	//
	// example:
	//
	// FILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlRequest) GetFileType() *string {
	return s.FileType
}

func (s *DescribeDownloadUrlRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDownloadUrlRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeDownloadUrlRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeDownloadUrlRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDownloadUrlRequest) SetFileType(v string) *DescribeDownloadUrlRequest {
	s.FileType = &v
	return s
}

func (s *DescribeDownloadUrlRequest) SetLang(v string) *DescribeDownloadUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadUrlRequest) SetRegId(v string) *DescribeDownloadUrlRequest {
	s.RegId = &v
	return s
}

func (s *DescribeDownloadUrlRequest) SetTaskId(v int64) *DescribeDownloadUrlRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDownloadUrlRequest) SetType(v string) *DescribeDownloadUrlRequest {
	s.Type = &v
	return s
}

func (s *DescribeDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
