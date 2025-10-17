// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSessionClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopSessionClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopSessionClusterResponse
	GetStatusCode() *int32
	SetBody(v *StopSessionClusterResponseBody) *StopSessionClusterResponse
	GetBody() *StopSessionClusterResponseBody
}

type StopSessionClusterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSessionClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSessionClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StopSessionClusterResponse) GoString() string {
	return s.String()
}

func (s *StopSessionClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopSessionClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopSessionClusterResponse) GetBody() *StopSessionClusterResponseBody {
	return s.Body
}

func (s *StopSessionClusterResponse) SetHeaders(v map[string]*string) *StopSessionClusterResponse {
	s.Headers = v
	return s
}

func (s *StopSessionClusterResponse) SetStatusCode(v int32) *StopSessionClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSessionClusterResponse) SetBody(v *StopSessionClusterResponseBody) *StopSessionClusterResponse {
	s.Body = v
	return s
}

func (s *StopSessionClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
