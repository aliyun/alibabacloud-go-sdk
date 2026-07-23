// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateExperimentRunResponseBody
	GetMessage() *string
	SetRecordId(v string) *CreateExperimentRunResponseBody
	GetRecordId() *string
	SetRequestId(v string) *CreateExperimentRunResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateExperimentRunResponseBody
	GetStatus() *string
}

type CreateExperimentRunResponseBody struct {
	// The message.
	//
	// example:
	//
	// Experiment created, execution started.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The experiment record ID. For online experiments, the format is typically exp-run-{uuid32}. For offline experiments, the format may also be a standard UUID.
	//
	// example:
	//
	// exp-run-f6d419b0ed3d43a7b585948a55efc07b
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019F89B5-1B07-3BB3-A32E-F5B007029E9C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The experiment record status. After creation, the status is typically pending.
	//
	// example:
	//
	// pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateExperimentRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentRunResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExperimentRunResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *CreateExperimentRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExperimentRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateExperimentRunResponseBody) SetMessage(v string) *CreateExperimentRunResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExperimentRunResponseBody) SetRecordId(v string) *CreateExperimentRunResponseBody {
	s.RecordId = &v
	return s
}

func (s *CreateExperimentRunResponseBody) SetRequestId(v string) *CreateExperimentRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperimentRunResponseBody) SetStatus(v string) *CreateExperimentRunResponseBody {
	s.Status = &v
	return s
}

func (s *CreateExperimentRunResponseBody) Validate() error {
	return dara.Validate(s)
}
