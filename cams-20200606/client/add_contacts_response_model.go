// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddContactsResponse
	GetStatusCode() *int32
	SetBody(v *AddContactsResponseBody) *AddContactsResponse
	GetBody() *AddContactsResponseBody
}

type AddContactsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddContactsResponse) GoString() string {
	return s.String()
}

func (s *AddContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddContactsResponse) GetBody() *AddContactsResponseBody {
	return s.Body
}

func (s *AddContactsResponse) SetHeaders(v map[string]*string) *AddContactsResponse {
	s.Headers = v
	return s
}

func (s *AddContactsResponse) SetStatusCode(v int32) *AddContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContactsResponse) SetBody(v *AddContactsResponseBody) *AddContactsResponse {
	s.Body = v
	return s
}

func (s *AddContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
