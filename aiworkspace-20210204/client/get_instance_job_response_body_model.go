// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *GetInstanceJobResponseBody
	GetGmtCreateTime() *string
	SetInstanceId(v string) *GetInstanceJobResponseBody
	GetInstanceId() *string
	SetInstanceJobId(v string) *GetInstanceJobResponseBody
	GetInstanceJobId() *string
	SetReasonMessage(v string) *GetInstanceJobResponseBody
	GetReasonMessage() *string
	SetRequestId(v string) *GetInstanceJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetInstanceJobResponseBody
	GetStatus() *string
	SetType(v string) *GetInstanceJobResponseBody
	GetType() *string
}

type GetInstanceJobResponseBody struct {
	// The creation time in UTC, in ISO 8601 format.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The instance ID. For example, if a job creates a custom role, this parameter returns the corresponding role ID.
	//
	// example:
	//
	// *****12qb3*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// CreateWorkspaceCustomRole-role-***abc*******
	InstanceJobId *string `json:"InstanceJobId,omitempty" xml:"InstanceJobId,omitempty"`
	// A message providing details about the job.
	//
	// example:
	//
	// workspace-example
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the job. Valid values:
	//
	// - Creating: The job is being created.
	//
	// - Updating: The job is being updated.
	//
	// - Deleting: The job is being deleted.
	//
	// - Successful: The job completed successfully (a final state).
	//
	// - Failed: The job failed (a final state).
	//
	// example:
	//
	// Successful
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job type. Valid values:
	//
	// - CreateWorkspaceCustomRole
	//
	// - UpdateWorkspaceCustomRole
	//
	// - DeleteWorkspaceCustomRole
	//
	// example:
	//
	// CreateWorkspaceCustomRole
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceJobResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceJobResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceJobResponseBody) GetInstanceJobId() *string {
	return s.InstanceJobId
}

func (s *GetInstanceJobResponseBody) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetInstanceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceJobResponseBody) GetType() *string {
	return s.Type
}

func (s *GetInstanceJobResponseBody) SetGmtCreateTime(v string) *GetInstanceJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceJobResponseBody) SetInstanceId(v string) *GetInstanceJobResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceJobResponseBody) SetInstanceJobId(v string) *GetInstanceJobResponseBody {
	s.InstanceJobId = &v
	return s
}

func (s *GetInstanceJobResponseBody) SetReasonMessage(v string) *GetInstanceJobResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceJobResponseBody) SetRequestId(v string) *GetInstanceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceJobResponseBody) SetStatus(v string) *GetInstanceJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceJobResponseBody) SetType(v string) *GetInstanceJobResponseBody {
	s.Type = &v
	return s
}

func (s *GetInstanceJobResponseBody) Validate() error {
	return dara.Validate(s)
}
