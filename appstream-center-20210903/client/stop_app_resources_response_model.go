// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAppResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAppResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAppResourcesResponse
	GetStatusCode() *int32
	SetBody(v *StopAppResourcesResponseBody) *StopAppResourcesResponse
	GetBody() *StopAppResourcesResponseBody
}

type StopAppResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAppResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAppResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAppResourcesResponse) GoString() string {
	return s.String()
}

func (s *StopAppResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAppResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAppResourcesResponse) GetBody() *StopAppResourcesResponseBody {
	return s.Body
}

func (s *StopAppResourcesResponse) SetHeaders(v map[string]*string) *StopAppResourcesResponse {
	s.Headers = v
	return s
}

func (s *StopAppResourcesResponse) SetStatusCode(v int32) *StopAppResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAppResourcesResponse) SetBody(v *StopAppResourcesResponseBody) *StopAppResourcesResponse {
	s.Body = v
	return s
}

func (s *StopAppResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
