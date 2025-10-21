// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSavepointFailure interface {
	dara.Model
	String() string
	GoString() string
	SetFailedAt(v int64) *SavepointFailure
	GetFailedAt() *int64
	SetMessage(v string) *SavepointFailure
	GetMessage() *string
	SetReason(v string) *SavepointFailure
	GetReason() *string
}

type SavepointFailure struct {
	// example:
	//
	// 1655006835
	FailedAt *int64  `json:"failedAt,omitempty" xml:"failedAt,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason   *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s SavepointFailure) String() string {
	return dara.Prettify(s)
}

func (s SavepointFailure) GoString() string {
	return s.String()
}

func (s *SavepointFailure) GetFailedAt() *int64 {
	return s.FailedAt
}

func (s *SavepointFailure) GetMessage() *string {
	return s.Message
}

func (s *SavepointFailure) GetReason() *string {
	return s.Reason
}

func (s *SavepointFailure) SetFailedAt(v int64) *SavepointFailure {
	s.FailedAt = &v
	return s
}

func (s *SavepointFailure) SetMessage(v string) *SavepointFailure {
	s.Message = &v
	return s
}

func (s *SavepointFailure) SetReason(v string) *SavepointFailure {
	s.Reason = &v
	return s
}

func (s *SavepointFailure) Validate() error {
	return dara.Validate(s)
}
