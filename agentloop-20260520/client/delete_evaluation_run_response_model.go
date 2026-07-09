// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluationRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEvaluationRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEvaluationRunResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEvaluationRunResponseBody) *DeleteEvaluationRunResponse
	GetBody() *DeleteEvaluationRunResponseBody
}

type DeleteEvaluationRunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEvaluationRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEvaluationRunResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluationRunResponse) GoString() string {
	return s.String()
}

func (s *DeleteEvaluationRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEvaluationRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEvaluationRunResponse) GetBody() *DeleteEvaluationRunResponseBody {
	return s.Body
}

func (s *DeleteEvaluationRunResponse) SetHeaders(v map[string]*string) *DeleteEvaluationRunResponse {
	s.Headers = v
	return s
}

func (s *DeleteEvaluationRunResponse) SetStatusCode(v int32) *DeleteEvaluationRunResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEvaluationRunResponse) SetBody(v *DeleteEvaluationRunResponseBody) *DeleteEvaluationRunResponse {
	s.Body = v
	return s
}

func (s *DeleteEvaluationRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
