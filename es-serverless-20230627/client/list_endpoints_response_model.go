// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListEndpointsResponseBody) *ListEndpointsResponse
	GetBody() *ListEndpointsResponseBody
}

type ListEndpointsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEndpointsResponse) GetBody() *ListEndpointsResponseBody {
	return s.Body
}

func (s *ListEndpointsResponse) SetHeaders(v map[string]*string) *ListEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListEndpointsResponse) SetStatusCode(v int32) *ListEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEndpointsResponse) SetBody(v *ListEndpointsResponseBody) *ListEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
