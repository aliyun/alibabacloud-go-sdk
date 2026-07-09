// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluationRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *UpdateEvaluationRunRequest
	GetStatus() *string
}

type UpdateEvaluationRunRequest struct {
	// The target status. Currently, only stop-related statuses are supported.
	//
	// example:
	//
	// Terminated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateEvaluationRunRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluationRunRequest) GoString() string {
	return s.String()
}

func (s *UpdateEvaluationRunRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateEvaluationRunRequest) SetStatus(v string) *UpdateEvaluationRunRequest {
	s.Status = &v
	return s
}

func (s *UpdateEvaluationRunRequest) Validate() error {
	return dara.Validate(s)
}
