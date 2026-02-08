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
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// b3865dc3-40fa-4afd-9fe4-dc7cda305a24
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// File Path of the job
	//
	// > This parameter corresponds to the Folder value returned by the GetJobDataUploadParams API.
	//
	// This parameter is required.
	//
	// example:
	//
	// UPLOADED/JOB/024f8cf0-c842-4c01-b74b-c8667e4579c7/e5897675-91f0-4e53-8af8-7e1ce4f2c089_KDC HR- 外呼话术2- 外呼名单模板 - Copy.xlsx
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
