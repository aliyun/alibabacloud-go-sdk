// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceConnectionsResponseBody) *ListServiceConnectionsResponse
	GetBody() *ListServiceConnectionsResponseBody
}

type ListServiceConnectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceConnectionsResponse) GetBody() *ListServiceConnectionsResponseBody {
	return s.Body
}

func (s *ListServiceConnectionsResponse) SetHeaders(v map[string]*string) *ListServiceConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceConnectionsResponse) SetStatusCode(v int32) *ListServiceConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceConnectionsResponse) SetBody(v *ListServiceConnectionsResponseBody) *ListServiceConnectionsResponse {
	s.Body = v
	return s
}

func (s *ListServiceConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
