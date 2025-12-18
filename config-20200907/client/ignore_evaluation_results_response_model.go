// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IgnoreEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IgnoreEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *IgnoreEvaluationResultsResponseBody) *IgnoreEvaluationResultsResponse
	GetBody() *IgnoreEvaluationResultsResponseBody
}

type IgnoreEvaluationResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IgnoreEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IgnoreEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s IgnoreEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IgnoreEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IgnoreEvaluationResultsResponse) GetBody() *IgnoreEvaluationResultsResponseBody {
	return s.Body
}

func (s *IgnoreEvaluationResultsResponse) SetHeaders(v map[string]*string) *IgnoreEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *IgnoreEvaluationResultsResponse) SetStatusCode(v int32) *IgnoreEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *IgnoreEvaluationResultsResponse) SetBody(v *IgnoreEvaluationResultsResponseBody) *IgnoreEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *IgnoreEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
