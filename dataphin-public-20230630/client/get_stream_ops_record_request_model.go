// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamOpsRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetStreamOpsRecordRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetStreamOpsRecordRequest
	GetOpUserId() *string
	SetOpsRecordId(v string) *GetStreamOpsRecordRequest
	GetOpsRecordId() *string
	SetProjectId(v int64) *GetStreamOpsRecordRequest
	GetProjectId() *int64
}

type GetStreamOpsRecordRequest struct {
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
	// The ID of the operation record. Set this parameter to the opsRecordId value returned by the ListStreamOpsRecords operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 646
	OpsRecordId *string `json:"OpsRecordId,omitempty" xml:"OpsRecordId,omitempty"`
	// The ID of the Dataphin project to which the task belongs. This parameter is used for permission verification in the production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7299592831043008
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetStreamOpsRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStreamOpsRecordRequest) GoString() string {
	return s.String()
}

func (s *GetStreamOpsRecordRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStreamOpsRecordRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetStreamOpsRecordRequest) GetOpsRecordId() *string {
	return s.OpsRecordId
}

func (s *GetStreamOpsRecordRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetStreamOpsRecordRequest) SetOpTenantId(v int64) *GetStreamOpsRecordRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStreamOpsRecordRequest) SetOpUserId(v string) *GetStreamOpsRecordRequest {
	s.OpUserId = &v
	return s
}

func (s *GetStreamOpsRecordRequest) SetOpsRecordId(v string) *GetStreamOpsRecordRequest {
	s.OpsRecordId = &v
	return s
}

func (s *GetStreamOpsRecordRequest) SetProjectId(v int64) *GetStreamOpsRecordRequest {
	s.ProjectId = &v
	return s
}

func (s *GetStreamOpsRecordRequest) Validate() error {
	return dara.Validate(s)
}
