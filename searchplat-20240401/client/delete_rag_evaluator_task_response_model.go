// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRagEvaluatorTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRagEvaluatorTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRagEvaluatorTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRagEvaluatorTaskResponseBody) *DeleteRagEvaluatorTaskResponse
	GetBody() *DeleteRagEvaluatorTaskResponseBody
}

type DeleteRagEvaluatorTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRagEvaluatorTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRagEvaluatorTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRagEvaluatorTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteRagEvaluatorTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRagEvaluatorTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRagEvaluatorTaskResponse) GetBody() *DeleteRagEvaluatorTaskResponseBody {
	return s.Body
}

func (s *DeleteRagEvaluatorTaskResponse) SetHeaders(v map[string]*string) *DeleteRagEvaluatorTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteRagEvaluatorTaskResponse) SetStatusCode(v int32) *DeleteRagEvaluatorTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRagEvaluatorTaskResponse) SetBody(v *DeleteRagEvaluatorTaskResponseBody) *DeleteRagEvaluatorTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteRagEvaluatorTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
