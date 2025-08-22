// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateExecuteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateExecuteJobRequest
	GetAppName() *string
	SetClusterId(v string) *OperateExecuteJobRequest
	GetClusterId() *string
	SetInstanceParameters(v string) *OperateExecuteJobRequest
	GetInstanceParameters() *string
	SetJobId(v int64) *OperateExecuteJobRequest
	GetJobId() *int64
	SetLabel(v string) *OperateExecuteJobRequest
	GetLabel() *string
	SetWorker(v string) *OperateExecuteJobRequest
	GetWorker() *string
}

type OperateExecuteJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// name=zhangsan
	InstanceParameters *string `json:"InstanceParameters,omitempty" xml:"InstanceParameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// http://192.168.1.5:9999/
	Worker *string `json:"Worker,omitempty" xml:"Worker,omitempty"`
}

func (s OperateExecuteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateExecuteJobRequest) GoString() string {
	return s.String()
}

func (s *OperateExecuteJobRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateExecuteJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateExecuteJobRequest) GetInstanceParameters() *string {
	return s.InstanceParameters
}

func (s *OperateExecuteJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *OperateExecuteJobRequest) GetLabel() *string {
	return s.Label
}

func (s *OperateExecuteJobRequest) GetWorker() *string {
	return s.Worker
}

func (s *OperateExecuteJobRequest) SetAppName(v string) *OperateExecuteJobRequest {
	s.AppName = &v
	return s
}

func (s *OperateExecuteJobRequest) SetClusterId(v string) *OperateExecuteJobRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateExecuteJobRequest) SetInstanceParameters(v string) *OperateExecuteJobRequest {
	s.InstanceParameters = &v
	return s
}

func (s *OperateExecuteJobRequest) SetJobId(v int64) *OperateExecuteJobRequest {
	s.JobId = &v
	return s
}

func (s *OperateExecuteJobRequest) SetLabel(v string) *OperateExecuteJobRequest {
	s.Label = &v
	return s
}

func (s *OperateExecuteJobRequest) SetWorker(v string) *OperateExecuteJobRequest {
	s.Worker = &v
	return s
}

func (s *OperateExecuteJobRequest) Validate() error {
	return dara.Validate(s)
}
