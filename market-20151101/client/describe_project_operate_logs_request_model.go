// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectOperateLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeProjectOperateLogsRequest
	GetInstanceId() *string
}

type DescribeProjectOperateLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProjectOperateLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectOperateLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectOperateLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeProjectOperateLogsRequest) SetInstanceId(v string) *DescribeProjectOperateLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProjectOperateLogsRequest) Validate() error {
	return dara.Validate(s)
}
