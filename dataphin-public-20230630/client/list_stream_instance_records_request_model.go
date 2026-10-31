// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamInstanceRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *ListStreamInstanceRecordsRequest
	GetFileId() *string
	SetOpTenantId(v int64) *ListStreamInstanceRecordsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListStreamInstanceRecordsRequest
	GetOpUserId() *string
	SetProjectId(v int64) *ListStreamInstanceRecordsRequest
	GetProjectId() *int64
}

type ListStreamInstanceRecordsRequest struct {
	// The ID of the task file. Set this parameter to the fileId value returned by the GetStreamTask operation. The value must be a positive integer.
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
	// The ID of the operator.
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

func (s ListStreamInstanceRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStreamInstanceRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListStreamInstanceRecordsRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListStreamInstanceRecordsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListStreamInstanceRecordsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListStreamInstanceRecordsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListStreamInstanceRecordsRequest) SetFileId(v string) *ListStreamInstanceRecordsRequest {
	s.FileId = &v
	return s
}

func (s *ListStreamInstanceRecordsRequest) SetOpTenantId(v int64) *ListStreamInstanceRecordsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListStreamInstanceRecordsRequest) SetOpUserId(v string) *ListStreamInstanceRecordsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListStreamInstanceRecordsRequest) SetProjectId(v int64) *ListStreamInstanceRecordsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListStreamInstanceRecordsRequest) Validate() error {
	return dara.Validate(s)
}
