// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskFromResourceImportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTaskFromResourceImportRequest
	GetClientToken() *string
	SetExportTaskId(v string) *CreateTaskFromResourceImportRequest
	GetExportTaskId() *string
	SetExportVersion(v string) *CreateTaskFromResourceImportRequest
	GetExportVersion() *string
	SetTaskName(v string) *CreateTaskFromResourceImportRequest
	GetTaskName() *string
}

type CreateTaskFromResourceImportRequest struct {
	// The client token that is used to ensure the idempotence of the request. The token must be 1 to 64 characters in length and can contain only letters, digits, and hyphens (-). A UUID is recommended.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The ID of the resource export task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ex-a1b2c3d4e5f60718293a4b5cd
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// The resource export version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-vpc-hangzhou
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateTaskFromResourceImportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskFromResourceImportRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskFromResourceImportRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTaskFromResourceImportRequest) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *CreateTaskFromResourceImportRequest) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *CreateTaskFromResourceImportRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateTaskFromResourceImportRequest) SetClientToken(v string) *CreateTaskFromResourceImportRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTaskFromResourceImportRequest) SetExportTaskId(v string) *CreateTaskFromResourceImportRequest {
	s.ExportTaskId = &v
	return s
}

func (s *CreateTaskFromResourceImportRequest) SetExportVersion(v string) *CreateTaskFromResourceImportRequest {
	s.ExportVersion = &v
	return s
}

func (s *CreateTaskFromResourceImportRequest) SetTaskName(v string) *CreateTaskFromResourceImportRequest {
	s.TaskName = &v
	return s
}

func (s *CreateTaskFromResourceImportRequest) Validate() error {
	return dara.Validate(s)
}
