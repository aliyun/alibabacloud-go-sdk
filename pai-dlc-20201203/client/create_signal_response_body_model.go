// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSignalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateSignalResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateSignalResponseBody
	GetRequestId() *string
	SetSignal(v string) *CreateSignalResponseBody
	GetSignal() *string
	SetSignalId(v string) *CreateSignalResponseBody
	GetSignalId() *string
	SetStatus(v string) *CreateSignalResponseBody
	GetStatus() *string
}

type CreateSignalResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// dlc********
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019F6385-7481-57A7-BEC9-***********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signal code.
	//
	// example:
	//
	// SIGUSR1
	Signal *string `json:"Signal,omitempty" xml:"Signal,omitempty"`
	// The signal ID.
	//
	// example:
	//
	// oper************
	SignalId *string `json:"SignalId,omitempty" xml:"SignalId,omitempty"`
	// The signal status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateSignalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSignalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSignalResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateSignalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSignalResponseBody) GetSignal() *string {
	return s.Signal
}

func (s *CreateSignalResponseBody) GetSignalId() *string {
	return s.SignalId
}

func (s *CreateSignalResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateSignalResponseBody) SetJobId(v string) *CreateSignalResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateSignalResponseBody) SetRequestId(v string) *CreateSignalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSignalResponseBody) SetSignal(v string) *CreateSignalResponseBody {
	s.Signal = &v
	return s
}

func (s *CreateSignalResponseBody) SetSignalId(v string) *CreateSignalResponseBody {
	s.SignalId = &v
	return s
}

func (s *CreateSignalResponseBody) SetStatus(v string) *CreateSignalResponseBody {
	s.Status = &v
	return s
}

func (s *CreateSignalResponseBody) Validate() error {
	return dara.Validate(s)
}
