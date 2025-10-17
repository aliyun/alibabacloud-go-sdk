// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSolutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSolutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSolutionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSolutionResponseBody) *DeleteSolutionResponse
	GetBody() *DeleteSolutionResponseBody
}

type DeleteSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSolutionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSolutionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSolutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSolutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSolutionResponse) GetBody() *DeleteSolutionResponseBody {
	return s.Body
}

func (s *DeleteSolutionResponse) SetHeaders(v map[string]*string) *DeleteSolutionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSolutionResponse) SetStatusCode(v int32) *DeleteSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSolutionResponse) SetBody(v *DeleteSolutionResponseBody) *DeleteSolutionResponse {
	s.Body = v
	return s
}

func (s *DeleteSolutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
