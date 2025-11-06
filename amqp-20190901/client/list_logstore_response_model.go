// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogstoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogstoreResponse
	GetStatusCode() *int32
	SetBody(v *ListLogstoreResponseBody) *ListLogstoreResponse
	GetBody() *ListLogstoreResponseBody
}

type ListLogstoreResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogstoreResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogstoreResponse) GoString() string {
	return s.String()
}

func (s *ListLogstoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogstoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogstoreResponse) GetBody() *ListLogstoreResponseBody {
	return s.Body
}

func (s *ListLogstoreResponse) SetHeaders(v map[string]*string) *ListLogstoreResponse {
	s.Headers = v
	return s
}

func (s *ListLogstoreResponse) SetStatusCode(v int32) *ListLogstoreResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogstoreResponse) SetBody(v *ListLogstoreResponseBody) *ListLogstoreResponse {
	s.Body = v
	return s
}

func (s *ListLogstoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
