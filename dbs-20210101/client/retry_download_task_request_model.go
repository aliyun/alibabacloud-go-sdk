// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDownloadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *RetryDownloadTaskRequest
	GetClusterName() *string
	SetInstanceName(v string) *RetryDownloadTaskRequest
	GetInstanceName() *string
	SetRegionCode(v string) *RetryDownloadTaskRequest
	GetRegionCode() *string
	SetTaskId(v string) *RetryDownloadTaskRequest
	GetTaskId() *string
}

type RetryDownloadTaskRequest struct {
	// example:
	//
	// dds-example
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// rm-example
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// example:
	//
	// dt-example
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RetryDownloadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *RetryDownloadTaskRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *RetryDownloadTaskRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RetryDownloadTaskRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *RetryDownloadTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RetryDownloadTaskRequest) SetClusterName(v string) *RetryDownloadTaskRequest {
	s.ClusterName = &v
	return s
}

func (s *RetryDownloadTaskRequest) SetInstanceName(v string) *RetryDownloadTaskRequest {
	s.InstanceName = &v
	return s
}

func (s *RetryDownloadTaskRequest) SetRegionCode(v string) *RetryDownloadTaskRequest {
	s.RegionCode = &v
	return s
}

func (s *RetryDownloadTaskRequest) SetTaskId(v string) *RetryDownloadTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RetryDownloadTaskRequest) Validate() error {
	return dara.Validate(s)
}
