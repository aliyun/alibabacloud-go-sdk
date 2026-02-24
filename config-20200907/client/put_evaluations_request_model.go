// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEvaluationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteMode(v bool) *PutEvaluationsRequest
	GetDeleteMode() *bool
	SetEvaluations(v string) *PutEvaluationsRequest
	GetEvaluations() *string
	SetResultToken(v string) *PutEvaluationsRequest
	GetResultToken() *string
}

type PutEvaluationsRequest struct {
	DeleteMode  *bool   `json:"DeleteMode,omitempty" xml:"DeleteMode,omitempty"`
	Evaluations *string `json:"Evaluations,omitempty" xml:"Evaluations,omitempty"`
	// This parameter is required.
	ResultToken *string `json:"ResultToken,omitempty" xml:"ResultToken,omitempty"`
}

func (s PutEvaluationsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutEvaluationsRequest) GoString() string {
	return s.String()
}

func (s *PutEvaluationsRequest) GetDeleteMode() *bool {
	return s.DeleteMode
}

func (s *PutEvaluationsRequest) GetEvaluations() *string {
	return s.Evaluations
}

func (s *PutEvaluationsRequest) GetResultToken() *string {
	return s.ResultToken
}

func (s *PutEvaluationsRequest) SetDeleteMode(v bool) *PutEvaluationsRequest {
	s.DeleteMode = &v
	return s
}

func (s *PutEvaluationsRequest) SetEvaluations(v string) *PutEvaluationsRequest {
	s.Evaluations = &v
	return s
}

func (s *PutEvaluationsRequest) SetResultToken(v string) *PutEvaluationsRequest {
	s.ResultToken = &v
	return s
}

func (s *PutEvaluationsRequest) Validate() error {
	return dara.Validate(s)
}
