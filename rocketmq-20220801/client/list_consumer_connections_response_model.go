// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerConnectionsResponseBody) *ListConsumerConnectionsResponse
	GetBody() *ListConsumerConnectionsResponseBody
}

type ListConsumerConnectionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerConnectionsResponse) GetBody() *ListConsumerConnectionsResponseBody {
	return s.Body
}

func (s *ListConsumerConnectionsResponse) SetHeaders(v map[string]*string) *ListConsumerConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerConnectionsResponse) SetStatusCode(v int32) *ListConsumerConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerConnectionsResponse) SetBody(v *ListConsumerConnectionsResponseBody) *ListConsumerConnectionsResponse {
	s.Body = v
	return s
}

func (s *ListConsumerConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
