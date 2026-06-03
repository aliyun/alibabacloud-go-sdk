// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobDataParsingTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateJobDataParsingTaskRequest
	GetInstanceId() *string
	SetJobFilePath(v string) *CreateJobDataParsingTaskRequest
	GetJobFilePath() *string
}

type CreateJobDataParsingTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b3865dc3-40fa-4afd-9fe4-dc7cda305a24
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobFilePath *string `json:"JobFilePath,omitempty" xml:"JobFilePath,omitempty"`
}

func (s CreateJobDataParsingTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobDataParsingTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateJobDataParsingTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateJobDataParsingTaskRequest) GetJobFilePath() *string {
	return s.JobFilePath
}

func (s *CreateJobDataParsingTaskRequest) SetInstanceId(v string) *CreateJobDataParsingTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateJobDataParsingTaskRequest) SetJobFilePath(v string) *CreateJobDataParsingTaskRequest {
	s.JobFilePath = &v
	return s
}

func (s *CreateJobDataParsingTaskRequest) Validate() error {
	return dara.Validate(s)
}
