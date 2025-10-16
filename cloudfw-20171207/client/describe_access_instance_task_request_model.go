// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *DescribeAccessInstanceTaskRequest
	GetAccessInstanceId() *string
	SetAccessInstanceType(v string) *DescribeAccessInstanceTaskRequest
	GetAccessInstanceType() *string
	SetLang(v string) *DescribeAccessInstanceTaskRequest
	GetLang() *string
	SetRegionNo(v string) *DescribeAccessInstanceTaskRequest
	GetRegionNo() *string
	SetTaskId(v string) *DescribeAccessInstanceTaskRequest
	GetTaskId() *string
}

type DescribeAccessInstanceTaskRequest struct {
	// example:
	//
	// pdi-3bc2f91695ee48bd9377
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// example:
	//
	// AckClusterConnector
	AccessInstanceType *string `json:"AccessInstanceType,omitempty" xml:"AccessInstanceType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// task-c92d4544ef7b6a42
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeAccessInstanceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceTaskRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DescribeAccessInstanceTaskRequest) GetAccessInstanceType() *string {
	return s.AccessInstanceType
}

func (s *DescribeAccessInstanceTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAccessInstanceTaskRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAccessInstanceTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeAccessInstanceTaskRequest) SetAccessInstanceId(v string) *DescribeAccessInstanceTaskRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *DescribeAccessInstanceTaskRequest) SetAccessInstanceType(v string) *DescribeAccessInstanceTaskRequest {
	s.AccessInstanceType = &v
	return s
}

func (s *DescribeAccessInstanceTaskRequest) SetLang(v string) *DescribeAccessInstanceTaskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAccessInstanceTaskRequest) SetRegionNo(v string) *DescribeAccessInstanceTaskRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAccessInstanceTaskRequest) SetTaskId(v string) *DescribeAccessInstanceTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeAccessInstanceTaskRequest) Validate() error {
	return dara.Validate(s)
}
