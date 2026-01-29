// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRayClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRayClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRayClusterResponse
	GetStatusCode() *int32
	SetBody(v *StopRayClusterResponseBody) *StopRayClusterResponse
	GetBody() *StopRayClusterResponseBody
}

type StopRayClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRayClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRayClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRayClusterResponse) GoString() string {
	return s.String()
}

func (s *StopRayClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRayClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRayClusterResponse) GetBody() *StopRayClusterResponseBody {
	return s.Body
}

func (s *StopRayClusterResponse) SetHeaders(v map[string]*string) *StopRayClusterResponse {
	s.Headers = v
	return s
}

func (s *StopRayClusterResponse) SetStatusCode(v int32) *StopRayClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRayClusterResponse) SetBody(v *StopRayClusterResponseBody) *StopRayClusterResponse {
	s.Body = v
	return s
}

func (s *StopRayClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
