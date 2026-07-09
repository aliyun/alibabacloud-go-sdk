// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEvaluatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEvaluatorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEvaluatorResponseBody) *DeleteEvaluatorResponse
	GetBody() *DeleteEvaluatorResponseBody
}

type DeleteEvaluatorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEvaluatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEvaluatorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluatorResponse) GoString() string {
	return s.String()
}

func (s *DeleteEvaluatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEvaluatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEvaluatorResponse) GetBody() *DeleteEvaluatorResponseBody {
	return s.Body
}

func (s *DeleteEvaluatorResponse) SetHeaders(v map[string]*string) *DeleteEvaluatorResponse {
	s.Headers = v
	return s
}

func (s *DeleteEvaluatorResponse) SetStatusCode(v int32) *DeleteEvaluatorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEvaluatorResponse) SetBody(v *DeleteEvaluatorResponseBody) *DeleteEvaluatorResponse {
	s.Body = v
	return s
}

func (s *DeleteEvaluatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
