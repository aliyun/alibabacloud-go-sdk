// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeIspFlushCacheTaskRequest
	GetLang() *string
	SetTaskId(v string) *DescribeIspFlushCacheTaskRequest
	GetTaskId() *string
}

type DescribeIspFlushCacheTaskRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeIspFlushCacheTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeIspFlushCacheTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeIspFlushCacheTaskRequest) SetLang(v string) *DescribeIspFlushCacheTaskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIspFlushCacheTaskRequest) SetTaskId(v string) *DescribeIspFlushCacheTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeIspFlushCacheTaskRequest) Validate() error {
	return dara.Validate(s)
}
