// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRecordDurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobRecordDuration(v int64) *UpdateJobRecordDurationRequest
	GetJobRecordDuration() *int64
}

type UpdateJobRecordDurationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30
	JobRecordDuration *int64 `json:"JobRecordDuration,omitempty" xml:"JobRecordDuration,omitempty"`
}

func (s UpdateJobRecordDurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRecordDurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRecordDurationRequest) GetJobRecordDuration() *int64 {
	return s.JobRecordDuration
}

func (s *UpdateJobRecordDurationRequest) SetJobRecordDuration(v int64) *UpdateJobRecordDurationRequest {
	s.JobRecordDuration = &v
	return s
}

func (s *UpdateJobRecordDurationRequest) Validate() error {
	return dara.Validate(s)
}
