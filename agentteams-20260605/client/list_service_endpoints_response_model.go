// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceEndpointsResponseBody) *ListServiceEndpointsResponse
	GetBody() *ListServiceEndpointsResponseBody
}

type ListServiceEndpointsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceEndpointsResponse) GetBody() *ListServiceEndpointsResponseBody {
	return s.Body
}

func (s *ListServiceEndpointsResponse) SetHeaders(v map[string]*string) *ListServiceEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceEndpointsResponse) SetStatusCode(v int32) *ListServiceEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceEndpointsResponse) SetBody(v *ListServiceEndpointsResponseBody) *ListServiceEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListServiceEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
