// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRecommendTaskDetailRequest
	GetLang() *string
	SetRegId(v string) *DescribeRecommendTaskDetailRequest
	GetRegId() *string
	SetTaskId(v int64) *DescribeRecommendTaskDetailRequest
	GetTaskId() *int64
}

type DescribeRecommendTaskDetailRequest struct {
	// Sets the language type for requests and responses, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 887
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DescribeRecommendTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecommendTaskDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRecommendTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeRecommendTaskDetailRequest) SetLang(v string) *DescribeRecommendTaskDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecommendTaskDetailRequest) SetRegId(v string) *DescribeRecommendTaskDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRecommendTaskDetailRequest) SetTaskId(v int64) *DescribeRecommendTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeRecommendTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
