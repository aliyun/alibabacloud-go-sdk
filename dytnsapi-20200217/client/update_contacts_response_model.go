// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContactsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContactsResponseBody) *UpdateContactsResponse
	GetBody() *UpdateContactsResponseBody
}

type UpdateContactsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactsResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContactsResponse) GetBody() *UpdateContactsResponseBody {
	return s.Body
}

func (s *UpdateContactsResponse) SetHeaders(v map[string]*string) *UpdateContactsResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactsResponse) SetStatusCode(v int32) *UpdateContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactsResponse) SetBody(v *UpdateContactsResponseBody) *UpdateContactsResponse {
	s.Body = v
	return s
}

func (s *UpdateContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
