// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReasonCode(v string) *CreateDiagnosisResponseBody
	GetReasonCode() *string
	SetReasonMessage(v string) *CreateDiagnosisResponseBody
	GetReasonMessage() *string
	SetSolutionMessage(v string) *CreateDiagnosisResponseBody
	GetSolutionMessage() *string
}

type CreateDiagnosisResponseBody struct {
	// example:
	//
	// Resource.InsufficientResource
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// Insufficient resource
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// Switch resource config
	SolutionMessage *string `json:"SolutionMessage,omitempty" xml:"SolutionMessage,omitempty"`
}

func (s CreateDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisResponseBody) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *CreateDiagnosisResponseBody) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *CreateDiagnosisResponseBody) GetSolutionMessage() *string {
	return s.SolutionMessage
}

func (s *CreateDiagnosisResponseBody) SetReasonCode(v string) *CreateDiagnosisResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *CreateDiagnosisResponseBody) SetReasonMessage(v string) *CreateDiagnosisResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *CreateDiagnosisResponseBody) SetSolutionMessage(v string) *CreateDiagnosisResponseBody {
	s.SolutionMessage = &v
	return s
}

func (s *CreateDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
