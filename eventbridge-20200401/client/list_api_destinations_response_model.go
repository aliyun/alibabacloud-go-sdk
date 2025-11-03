// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiDestinationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiDestinationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiDestinationsResponse
	GetStatusCode() *int32
	SetBody(v *ListApiDestinationsResponseBody) *ListApiDestinationsResponse
	GetBody() *ListApiDestinationsResponseBody
}

type ListApiDestinationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiDestinationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiDestinationsResponse) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiDestinationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiDestinationsResponse) GetBody() *ListApiDestinationsResponseBody {
	return s.Body
}

func (s *ListApiDestinationsResponse) SetHeaders(v map[string]*string) *ListApiDestinationsResponse {
	s.Headers = v
	return s
}

func (s *ListApiDestinationsResponse) SetStatusCode(v int32) *ListApiDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiDestinationsResponse) SetBody(v *ListApiDestinationsResponseBody) *ListApiDestinationsResponse {
	s.Body = v
	return s
}

func (s *ListApiDestinationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
