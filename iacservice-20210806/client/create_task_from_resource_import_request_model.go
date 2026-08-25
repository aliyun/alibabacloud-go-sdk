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
	// This parameter is required.
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// This parameter is required.
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// This parameter is required.
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
