// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobFailure interface {
	dara.Model
	String() string
	GoString() string
	SetFailedAt(v int64) *JobFailure
	GetFailedAt() *int64
	SetMessage(v string) *JobFailure
	GetMessage() *string
	SetReason(v string) *JobFailure
	GetReason() *string
}

type JobFailure struct {
	// example:
	//
	// 1660120062
	FailedAt *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason   *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s JobFailure) String() string {
	return dara.Prettify(s)
}

func (s JobFailure) GoString() string {
	return s.String()
}

func (s *JobFailure) GetFailedAt() *int64 {
	return s.FailedAt
}

func (s *JobFailure) GetMessage() *string {
	return s.Message
}

func (s *JobFailure) GetReason() *string {
	return s.Reason
}

func (s *JobFailure) SetFailedAt(v int64) *JobFailure {
	s.FailedAt = &v
	return s
}

func (s *JobFailure) SetMessage(v string) *JobFailure {
	s.Message = &v
	return s
}

func (s *JobFailure) SetReason(v string) *JobFailure {
	s.Reason = &v
	return s
}

func (s *JobFailure) Validate() error {
	return dara.Validate(s)
}
