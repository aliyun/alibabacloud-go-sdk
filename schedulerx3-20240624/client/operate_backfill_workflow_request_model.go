// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBackfillWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateBackfillWorkflowRequest
	GetAppName() *string
	SetClusterId(v string) *OperateBackfillWorkflowRequest
	GetClusterId() *string
	SetEndDate(v string) *OperateBackfillWorkflowRequest
	GetEndDate() *string
	SetStartDate(v string) *OperateBackfillWorkflowRequest
	GetStartDate() *string
	SetWorkflowId(v int64) *OperateBackfillWorkflowRequest
	GetWorkflowId() *int64
}

type OperateBackfillWorkflowRequest struct {
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
	//
	// example:
	//
	// 2025-11-03
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-11-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s OperateBackfillWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateBackfillWorkflowRequest) GoString() string {
	return s.String()
}

func (s *OperateBackfillWorkflowRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateBackfillWorkflowRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateBackfillWorkflowRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *OperateBackfillWorkflowRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *OperateBackfillWorkflowRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *OperateBackfillWorkflowRequest) SetAppName(v string) *OperateBackfillWorkflowRequest {
	s.AppName = &v
	return s
}

func (s *OperateBackfillWorkflowRequest) SetClusterId(v string) *OperateBackfillWorkflowRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateBackfillWorkflowRequest) SetEndDate(v string) *OperateBackfillWorkflowRequest {
	s.EndDate = &v
	return s
}

func (s *OperateBackfillWorkflowRequest) SetStartDate(v string) *OperateBackfillWorkflowRequest {
	s.StartDate = &v
	return s
}

func (s *OperateBackfillWorkflowRequest) SetWorkflowId(v int64) *OperateBackfillWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *OperateBackfillWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
