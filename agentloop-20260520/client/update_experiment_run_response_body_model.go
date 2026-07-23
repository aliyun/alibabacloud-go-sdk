// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateExperimentRunResponseBody
	GetMessage() *string
	SetRecordId(v string) *UpdateExperimentRunResponseBody
	GetRecordId() *string
	SetRequestId(v string) *UpdateExperimentRunResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateExperimentRunResponseBody
	GetStatus() *string
}

type UpdateExperimentRunResponseBody struct {
	// The prompt message.
	//
	// example:
	//
	// Experiment record updated successfully.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The experiment record ID.
	//
	// example:
	//
	// a5397261-6e6d-4e45-bf52-feb8686f7524
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019F89B5-1B07-3BB3-A32E-F5B007029E9C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The status after the update (the persisted value).
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateExperimentRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentRunResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRunResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateExperimentRunResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateExperimentRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExperimentRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateExperimentRunResponseBody) SetMessage(v string) *UpdateExperimentRunResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExperimentRunResponseBody) SetRecordId(v string) *UpdateExperimentRunResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateExperimentRunResponseBody) SetRequestId(v string) *UpdateExperimentRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentRunResponseBody) SetStatus(v string) *UpdateExperimentRunResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateExperimentRunResponseBody) Validate() error {
	return dara.Validate(s)
}
