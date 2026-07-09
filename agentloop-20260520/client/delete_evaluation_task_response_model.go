// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEvaluationTaskResponseBody) *DeleteEvaluationTaskResponse
	GetBody() *DeleteEvaluationTaskResponseBody
}

type DeleteEvaluationTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEvaluationTaskResponse) GetBody() *DeleteEvaluationTaskResponseBody {
	return s.Body
}

func (s *DeleteEvaluationTaskResponse) SetHeaders(v map[string]*string) *DeleteEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteEvaluationTaskResponse) SetStatusCode(v int32) *DeleteEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEvaluationTaskResponse) SetBody(v *DeleteEvaluationTaskResponseBody) *DeleteEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteEvaluationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
