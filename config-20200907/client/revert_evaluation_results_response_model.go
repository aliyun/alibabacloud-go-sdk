// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevertEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevertEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *RevertEvaluationResultsResponseBody) *RevertEvaluationResultsResponse
	GetBody() *RevertEvaluationResultsResponseBody
}

type RevertEvaluationResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s RevertEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevertEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevertEvaluationResultsResponse) GetBody() *RevertEvaluationResultsResponseBody {
	return s.Body
}

func (s *RevertEvaluationResultsResponse) SetHeaders(v map[string]*string) *RevertEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *RevertEvaluationResultsResponse) SetStatusCode(v int32) *RevertEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertEvaluationResultsResponse) SetBody(v *RevertEvaluationResultsResponseBody) *RevertEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *RevertEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
