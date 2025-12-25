// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateEnableWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateEnableWorkflowsRequest
	GetAppName() *string
	SetClusterId(v string) *OperateEnableWorkflowsRequest
	GetClusterId() *string
	SetWorkflowIds(v []*int64) *OperateEnableWorkflowsRequest
	GetWorkflowIds() []*int64
}

type OperateEnableWorkflowsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxl-job-executor-sample
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	WorkflowIds []*int64 `json:"WorkflowIds,omitempty" xml:"WorkflowIds,omitempty" type:"Repeated"`
}

func (s OperateEnableWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateEnableWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *OperateEnableWorkflowsRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateEnableWorkflowsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateEnableWorkflowsRequest) GetWorkflowIds() []*int64 {
	return s.WorkflowIds
}

func (s *OperateEnableWorkflowsRequest) SetAppName(v string) *OperateEnableWorkflowsRequest {
	s.AppName = &v
	return s
}

func (s *OperateEnableWorkflowsRequest) SetClusterId(v string) *OperateEnableWorkflowsRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateEnableWorkflowsRequest) SetWorkflowIds(v []*int64) *OperateEnableWorkflowsRequest {
	s.WorkflowIds = v
	return s
}

func (s *OperateEnableWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
