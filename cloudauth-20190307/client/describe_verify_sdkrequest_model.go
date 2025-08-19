// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySDKRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeVerifySDKRequest
	GetTaskId() *string
}

type DescribeVerifySDKRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1KQMcnLd4m37LN2D0F0WCD
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeVerifySDKRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySDKRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeVerifySDKRequest) SetTaskId(v string) *DescribeVerifySDKRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeVerifySDKRequest) Validate() error {
	return dara.Validate(s)
}
