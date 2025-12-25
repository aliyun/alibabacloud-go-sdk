// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteWorkflowsRequest
	GetAppName() *string
	SetClusterId(v string) *DeleteWorkflowsRequest
	GetClusterId() *string
	SetDeleteJobs(v bool) *DeleteWorkflowsRequest
	GetDeleteJobs() *bool
	SetWorkflowIds(v []*int64) *DeleteWorkflowsRequest
	GetWorkflowIds() []*int64
}

type DeleteWorkflowsRequest struct {
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
	// xxljob-a1804a3226d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	DeleteJobs *bool `json:"DeleteJobs,omitempty" xml:"DeleteJobs,omitempty"`
	// This parameter is required.
	WorkflowIds []*int64 `json:"WorkflowIds,omitempty" xml:"WorkflowIds,omitempty" type:"Repeated"`
}

func (s DeleteWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowsRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteWorkflowsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteWorkflowsRequest) GetDeleteJobs() *bool {
	return s.DeleteJobs
}

func (s *DeleteWorkflowsRequest) GetWorkflowIds() []*int64 {
	return s.WorkflowIds
}

func (s *DeleteWorkflowsRequest) SetAppName(v string) *DeleteWorkflowsRequest {
	s.AppName = &v
	return s
}

func (s *DeleteWorkflowsRequest) SetClusterId(v string) *DeleteWorkflowsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteWorkflowsRequest) SetDeleteJobs(v bool) *DeleteWorkflowsRequest {
	s.DeleteJobs = &v
	return s
}

func (s *DeleteWorkflowsRequest) SetWorkflowIds(v []*int64) *DeleteWorkflowsRequest {
	s.WorkflowIds = v
	return s
}

func (s *DeleteWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
