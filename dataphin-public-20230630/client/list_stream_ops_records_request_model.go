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
	// The ID of the real-time task file. The value is a string representing a positive integer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7333610201479808
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the operation tenant. The gateway verifies that the current identity has access to the tenant and then injects the tenant ID into the backend.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The optional ID of the operator user. This parameter follows the gateway identity verification rules and is not directly bound to the backend user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The project ID.
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
