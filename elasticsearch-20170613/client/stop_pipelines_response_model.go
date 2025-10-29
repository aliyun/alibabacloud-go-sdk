// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopPipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopPipelinesResponse
	GetStatusCode() *int32
	SetBody(v *StopPipelinesResponseBody) *StopPipelinesResponse
	GetBody() *StopPipelinesResponseBody
}

type StopPipelinesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopPipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopPipelinesResponse) GoString() string {
	return s.String()
}

func (s *StopPipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopPipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopPipelinesResponse) GetBody() *StopPipelinesResponseBody {
	return s.Body
}

func (s *StopPipelinesResponse) SetHeaders(v map[string]*string) *StopPipelinesResponse {
	s.Headers = v
	return s
}

func (s *StopPipelinesResponse) SetStatusCode(v int32) *StopPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPipelinesResponse) SetBody(v *StopPipelinesResponseBody) *StopPipelinesResponse {
	s.Body = v
	return s
}

func (s *StopPipelinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
