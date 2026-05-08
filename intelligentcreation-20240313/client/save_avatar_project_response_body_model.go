// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAvatarProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *SaveAvatarProjectResponseBody
	GetAgentId() *string
	SetErrorCode(v string) *SaveAvatarProjectResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SaveAvatarProjectResponseBody
	GetErrorMessage() *string
	SetErrorMsg(v string) *SaveAvatarProjectResponseBody
	GetErrorMsg() *string
	SetProjectId(v string) *SaveAvatarProjectResponseBody
	GetProjectId() *string
	SetProjectName(v string) *SaveAvatarProjectResponseBody
	GetProjectName() *string
	SetRequestId(v string) *SaveAvatarProjectResponseBody
	GetRequestId() *string
	SetStatus(v string) *SaveAvatarProjectResponseBody
	GetStatus() *string
}

type SaveAvatarProjectResponseBody struct {
	// example:
	//
	// 1000222
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 040002
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 812907463682949120
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// doc_test_3
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveAvatarProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *SaveAvatarProjectResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SaveAvatarProjectResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveAvatarProjectResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SaveAvatarProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *SaveAvatarProjectResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *SaveAvatarProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveAvatarProjectResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SaveAvatarProjectResponseBody) SetAgentId(v string) *SaveAvatarProjectResponseBody {
	s.AgentId = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetErrorCode(v string) *SaveAvatarProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetErrorMessage(v string) *SaveAvatarProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetErrorMsg(v string) *SaveAvatarProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetProjectId(v string) *SaveAvatarProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetProjectName(v string) *SaveAvatarProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetRequestId(v string) *SaveAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) SetStatus(v string) *SaveAvatarProjectResponseBody {
	s.Status = &v
	return s
}

func (s *SaveAvatarProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
