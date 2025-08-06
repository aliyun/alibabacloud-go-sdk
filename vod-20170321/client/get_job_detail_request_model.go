// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetJobDetailRequest
	GetJobId() *string
	SetJobType(v string) *GetJobDetailRequest
	GetJobType() *string
}

type GetJobDetailRequest struct {
	// The ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c9dff***************59d50a967f5
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task type. Valid values:
	//
	// 	- transcode
	//
	// 	- snapshot
	//
	// 	- ai
	//
	// This parameter is required.
	//
	// example:
	//
	// transcode
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
}

func (s GetJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobDetailRequest) GoString() string {
	return s.String()
}

func (s *GetJobDetailRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetJobDetailRequest) GetJobType() *string {
	return s.JobType
}

func (s *GetJobDetailRequest) SetJobId(v string) *GetJobDetailRequest {
	s.JobId = &v
	return s
}

func (s *GetJobDetailRequest) SetJobType(v string) *GetJobDetailRequest {
	s.JobType = &v
	return s
}

func (s *GetJobDetailRequest) Validate() error {
	return dara.Validate(s)
}
