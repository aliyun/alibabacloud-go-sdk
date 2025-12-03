// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemWorkFlowStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListWorkItemWorkFlowStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListWorkItemWorkFlowStatusResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListWorkItemWorkFlowStatusResponseBody
	GetRequestId() *string
	SetStatuses(v []*ListWorkItemWorkFlowStatusResponseBodyStatuses) *ListWorkItemWorkFlowStatusResponseBody
	GetStatuses() []*ListWorkItemWorkFlowStatusResponseBodyStatuses
	SetSuccess(v bool) *ListWorkItemWorkFlowStatusResponseBody
	GetSuccess() *bool
}

type ListWorkItemWorkFlowStatusResponseBody struct {
	// example:
	//
	// 例：Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// errormessage
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Statuses  []*ListWorkItemWorkFlowStatusResponseBodyStatuses `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	// example:
	//
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListWorkItemWorkFlowStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemWorkFlowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkItemWorkFlowStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkItemWorkFlowStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListWorkItemWorkFlowStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkItemWorkFlowStatusResponseBody) GetStatuses() []*ListWorkItemWorkFlowStatusResponseBodyStatuses {
	return s.Statuses
}

func (s *ListWorkItemWorkFlowStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetErrorCode(v string) *ListWorkItemWorkFlowStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetErrorMessage(v string) *ListWorkItemWorkFlowStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetRequestId(v string) *ListWorkItemWorkFlowStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetStatuses(v []*ListWorkItemWorkFlowStatusResponseBodyStatuses) *ListWorkItemWorkFlowStatusResponseBody {
	s.Statuses = v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetSuccess(v bool) *ListWorkItemWorkFlowStatusResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) Validate() error {
	if s.Statuses != nil {
		for _, item := range s.Statuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkItemWorkFlowStatusResponseBodyStatuses struct {
	// example:
	//
	// 用户阿里云pk，例如19xxxx31947xxxx
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 该状态的具体信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1620455467000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1641870287000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// dfexxxxxf4fee18xxxxx36
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 用户阿里云pk，例如19xxxx31947xxxx
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 例：待处理
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Workitem
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// system
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// 例如：1
	WorkflowStageIdentifier *string `json:"workflowStageIdentifier,omitempty" xml:"workflowStageIdentifier,omitempty"`
	// example:
	//
	// 例：确认阶段
	WorkflowStageName *string `json:"workflowStageName,omitempty" xml:"workflowStageName,omitempty"`
}

func (s ListWorkItemWorkFlowStatusResponseBodyStatuses) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemWorkFlowStatusResponseBodyStatuses) GoString() string {
	return s.String()
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetDescription() *string {
	return s.Description
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetModifier() *string {
	return s.Modifier
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetName() *string {
	return s.Name
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetSource() *string {
	return s.Source
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetWorkflowStageIdentifier() *string {
	return s.WorkflowStageIdentifier
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) GetWorkflowStageName() *string {
	return s.WorkflowStageName
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetCreator(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Creator = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetDescription(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Description = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetGmtCreate(v int64) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetGmtModified(v int64) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.GmtModified = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetIdentifier(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Identifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetModifier(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Modifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetName(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Name = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetResourceType(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.ResourceType = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetSource(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Source = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetWorkflowStageIdentifier(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.WorkflowStageIdentifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetWorkflowStageName(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.WorkflowStageName = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) Validate() error {
	return dara.Validate(s)
}
