// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamOpsRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *ListStreamOpsRecordsRequest
	GetFileId() *string
	SetOpTenantId(v int64) *ListStreamOpsRecordsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListStreamOpsRecordsRequest
	GetOpUserId() *string
	SetProjectId(v int64) *ListStreamOpsRecordsRequest
	GetProjectId() *int64
}

type ListStreamOpsRecordsRequest struct {
	// The ID of the task file. Set this parameter to the fileId value returned by the GetStreamTask operation. The value must be a positive integer. Do not set this parameter to nodeId or taskId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7333610201479808
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID of the Dataphin project to which the task belongs. This parameter is used for permission verification in the production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7299592831043008
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListStreamOpsRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStreamOpsRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListStreamOpsRecordsRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListStreamOpsRecordsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListStreamOpsRecordsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListStreamOpsRecordsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListStreamOpsRecordsRequest) SetFileId(v string) *ListStreamOpsRecordsRequest {
	s.FileId = &v
	return s
}

func (s *ListStreamOpsRecordsRequest) SetOpTenantId(v int64) *ListStreamOpsRecordsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListStreamOpsRecordsRequest) SetOpUserId(v string) *ListStreamOpsRecordsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListStreamOpsRecordsRequest) SetProjectId(v int64) *ListStreamOpsRecordsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListStreamOpsRecordsRequest) Validate() error {
	return dara.Validate(s)
}
