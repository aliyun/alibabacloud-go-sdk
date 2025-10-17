// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDedicatedClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDedicatedClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDedicatedClusterResponse
	GetStatusCode() *int32
	SetBody(v *StopDedicatedClusterResponseBody) *StopDedicatedClusterResponse
	GetBody() *StopDedicatedClusterResponseBody
}

type StopDedicatedClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDedicatedClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDedicatedClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDedicatedClusterResponse) GoString() string {
	return s.String()
}

func (s *StopDedicatedClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDedicatedClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDedicatedClusterResponse) GetBody() *StopDedicatedClusterResponseBody {
	return s.Body
}

func (s *StopDedicatedClusterResponse) SetHeaders(v map[string]*string) *StopDedicatedClusterResponse {
	s.Headers = v
	return s
}

func (s *StopDedicatedClusterResponse) SetStatusCode(v int32) *StopDedicatedClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDedicatedClusterResponse) SetBody(v *StopDedicatedClusterResponseBody) *StopDedicatedClusterResponse {
	s.Body = v
	return s
}

func (s *StopDedicatedClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
