// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContactsResponse
	GetStatusCode() *int32
	SetBody(v *ListContactsResponseBody) *ListContactsResponse
	GetBody() *ListContactsResponseBody
}

type ListContactsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContactsResponse) GoString() string {
	return s.String()
}

func (s *ListContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContactsResponse) GetBody() *ListContactsResponseBody {
	return s.Body
}

func (s *ListContactsResponse) SetHeaders(v map[string]*string) *ListContactsResponse {
	s.Headers = v
	return s
}

func (s *ListContactsResponse) SetStatusCode(v int32) *ListContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactsResponse) SetBody(v *ListContactsResponseBody) *ListContactsResponse {
	s.Body = v
	return s
}

func (s *ListContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
