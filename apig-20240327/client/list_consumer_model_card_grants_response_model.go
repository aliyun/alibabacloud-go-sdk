// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerModelCardGrantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerModelCardGrantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerModelCardGrantsResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerModelCardGrantsResponseBody) *ListConsumerModelCardGrantsResponse
	GetBody() *ListConsumerModelCardGrantsResponseBody
}

type ListConsumerModelCardGrantsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerModelCardGrantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerModelCardGrantsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerModelCardGrantsResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerModelCardGrantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerModelCardGrantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerModelCardGrantsResponse) GetBody() *ListConsumerModelCardGrantsResponseBody {
	return s.Body
}

func (s *ListConsumerModelCardGrantsResponse) SetHeaders(v map[string]*string) *ListConsumerModelCardGrantsResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerModelCardGrantsResponse) SetStatusCode(v int32) *ListConsumerModelCardGrantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerModelCardGrantsResponse) SetBody(v *ListConsumerModelCardGrantsResponseBody) *ListConsumerModelCardGrantsResponse {
	s.Body = v
	return s
}

func (s *ListConsumerModelCardGrantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
