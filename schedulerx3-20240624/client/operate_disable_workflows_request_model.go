// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDisableWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateDisableWorkflowsRequest
	GetAppName() *string
	SetClusterId(v string) *OperateDisableWorkflowsRequest
	GetClusterId() *string
	SetWorkflowIds(v []*int64) *OperateDisableWorkflowsRequest
	GetWorkflowIds() []*int64
}

type OperateDisableWorkflowsRequest struct {
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
	// This parameter is required.
	WorkflowIds []*int64 `json:"WorkflowIds,omitempty" xml:"WorkflowIds,omitempty" type:"Repeated"`
}

func (s OperateDisableWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateDisableWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *OperateDisableWorkflowsRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateDisableWorkflowsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateDisableWorkflowsRequest) GetWorkflowIds() []*int64 {
	return s.WorkflowIds
}

func (s *OperateDisableWorkflowsRequest) SetAppName(v string) *OperateDisableWorkflowsRequest {
	s.AppName = &v
	return s
}

func (s *OperateDisableWorkflowsRequest) SetClusterId(v string) *OperateDisableWorkflowsRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDisableWorkflowsRequest) SetWorkflowIds(v []*int64) *OperateDisableWorkflowsRequest {
	s.WorkflowIds = v
	return s
}

func (s *OperateDisableWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
