// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationScoreHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvaluationScoreHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvaluationScoreHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListEvaluationScoreHistoryResponseBody) *ListEvaluationScoreHistoryResponse
	GetBody() *ListEvaluationScoreHistoryResponseBody
}

type ListEvaluationScoreHistoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationScoreHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationScoreHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationScoreHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvaluationScoreHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvaluationScoreHistoryResponse) GetBody() *ListEvaluationScoreHistoryResponseBody {
	return s.Body
}

func (s *ListEvaluationScoreHistoryResponse) SetHeaders(v map[string]*string) *ListEvaluationScoreHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationScoreHistoryResponse) SetStatusCode(v int32) *ListEvaluationScoreHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationScoreHistoryResponse) SetBody(v *ListEvaluationScoreHistoryResponseBody) *ListEvaluationScoreHistoryResponse {
	s.Body = v
	return s
}

func (s *ListEvaluationScoreHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
