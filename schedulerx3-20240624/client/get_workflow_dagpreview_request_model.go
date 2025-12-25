// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDAGPreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetWorkflowDAGPreviewRequest
	GetAppName() *string
	SetClusterId(v string) *GetWorkflowDAGPreviewRequest
	GetClusterId() *string
	SetDagVersion(v string) *GetWorkflowDAGPreviewRequest
	GetDagVersion() *string
	SetWorkflowId(v int64) *GetWorkflowDAGPreviewRequest
	GetWorkflowId() *int64
}

type GetWorkflowDAGPreviewRequest struct {
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
	// v1
	DagVersion *string `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
	// example:
	//
	// 20
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowDAGPreviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGPreviewRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGPreviewRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetWorkflowDAGPreviewRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetWorkflowDAGPreviewRequest) GetDagVersion() *string {
	return s.DagVersion
}

func (s *GetWorkflowDAGPreviewRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkflowDAGPreviewRequest) SetAppName(v string) *GetWorkflowDAGPreviewRequest {
	s.AppName = &v
	return s
}

func (s *GetWorkflowDAGPreviewRequest) SetClusterId(v string) *GetWorkflowDAGPreviewRequest {
	s.ClusterId = &v
	return s
}

func (s *GetWorkflowDAGPreviewRequest) SetDagVersion(v string) *GetWorkflowDAGPreviewRequest {
	s.DagVersion = &v
	return s
}

func (s *GetWorkflowDAGPreviewRequest) SetWorkflowId(v int64) *GetWorkflowDAGPreviewRequest {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowDAGPreviewRequest) Validate() error {
	return dara.Validate(s)
}
