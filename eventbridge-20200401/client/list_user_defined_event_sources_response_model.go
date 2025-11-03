// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDefinedEventSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserDefinedEventSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserDefinedEventSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserDefinedEventSourcesResponseBody) *ListUserDefinedEventSourcesResponse
	GetBody() *ListUserDefinedEventSourcesResponseBody
}

type ListUserDefinedEventSourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDefinedEventSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDefinedEventSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserDefinedEventSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserDefinedEventSourcesResponse) GetBody() *ListUserDefinedEventSourcesResponseBody {
	return s.Body
}

func (s *ListUserDefinedEventSourcesResponse) SetHeaders(v map[string]*string) *ListUserDefinedEventSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListUserDefinedEventSourcesResponse) SetStatusCode(v int32) *ListUserDefinedEventSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponse) SetBody(v *ListUserDefinedEventSourcesResponseBody) *ListUserDefinedEventSourcesResponse {
	s.Body = v
	return s
}

func (s *ListUserDefinedEventSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
