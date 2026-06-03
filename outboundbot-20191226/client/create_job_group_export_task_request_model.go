// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobGroupExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateJobGroupExportTaskRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *CreateJobGroupExportTaskRequest
	GetJobGroupId() *string
	SetOption(v []*string) *CreateJobGroupExportTaskRequest
	GetOption() []*string
}

type CreateJobGroupExportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// baf6dfdc-eb79-4c63-ab19-c56388b1fbdd
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// []
	Option []*string `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
}

func (s CreateJobGroupExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateJobGroupExportTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateJobGroupExportTaskRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *CreateJobGroupExportTaskRequest) GetOption() []*string {
	return s.Option
}

func (s *CreateJobGroupExportTaskRequest) SetInstanceId(v string) *CreateJobGroupExportTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateJobGroupExportTaskRequest) SetJobGroupId(v string) *CreateJobGroupExportTaskRequest {
	s.JobGroupId = &v
	return s
}

func (s *CreateJobGroupExportTaskRequest) SetOption(v []*string) *CreateJobGroupExportTaskRequest {
	s.Option = v
	return s
}

func (s *CreateJobGroupExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
