// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogStoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogStoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogStoresResponse
	GetStatusCode() *int32
	SetBody(v *ListLogStoresResponseBody) *ListLogStoresResponse
	GetBody() *ListLogStoresResponseBody
}

type ListLogStoresResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogStoresResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogStoresResponse) GoString() string {
	return s.String()
}

func (s *ListLogStoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogStoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogStoresResponse) GetBody() *ListLogStoresResponseBody {
	return s.Body
}

func (s *ListLogStoresResponse) SetHeaders(v map[string]*string) *ListLogStoresResponse {
	s.Headers = v
	return s
}

func (s *ListLogStoresResponse) SetStatusCode(v int32) *ListLogStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogStoresResponse) SetBody(v *ListLogStoresResponseBody) *ListLogStoresResponse {
	s.Body = v
	return s
}

func (s *ListLogStoresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
