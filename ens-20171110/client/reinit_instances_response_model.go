// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinitInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReinitInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReinitInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ReinitInstancesResponseBody) *ReinitInstancesResponse
	GetBody() *ReinitInstancesResponseBody
}

type ReinitInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReinitInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReinitInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ReinitInstancesResponse) GoString() string {
	return s.String()
}

func (s *ReinitInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReinitInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReinitInstancesResponse) GetBody() *ReinitInstancesResponseBody {
	return s.Body
}

func (s *ReinitInstancesResponse) SetHeaders(v map[string]*string) *ReinitInstancesResponse {
	s.Headers = v
	return s
}

func (s *ReinitInstancesResponse) SetStatusCode(v int32) *ReinitInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReinitInstancesResponse) SetBody(v *ReinitInstancesResponseBody) *ReinitInstancesResponse {
	s.Body = v
	return s
}

func (s *ReinitInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
