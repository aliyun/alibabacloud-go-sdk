// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyTaskWaitingQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *DescribeComfyTaskWaitingQueueRequest
	GetHiveId() *string
}

type DescribeComfyTaskWaitingQueueRequest struct {
	// The waiting queue information of a specified Hive.
	//
	// example:
	//
	// hive-26cd567b35c04a0a90f0xxxxx
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
}

func (s DescribeComfyTaskWaitingQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyTaskWaitingQueueRequest) GoString() string {
	return s.String()
}

func (s *DescribeComfyTaskWaitingQueueRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *DescribeComfyTaskWaitingQueueRequest) SetHiveId(v string) *DescribeComfyTaskWaitingQueueRequest {
	s.HiveId = &v
	return s
}

func (s *DescribeComfyTaskWaitingQueueRequest) Validate() error {
	return dara.Validate(s)
}
