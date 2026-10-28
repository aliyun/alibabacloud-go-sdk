// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContactResponse
	GetStatusCode() *int32
	SetBody(v *CreateContactResponseBody) *CreateContactResponse
	GetBody() *CreateContactResponseBody
}

type CreateContactResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContactResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContactResponse) GoString() string {
	return s.String()
}

func (s *CreateContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContactResponse) GetBody() *CreateContactResponseBody {
	return s.Body
}

func (s *CreateContactResponse) SetHeaders(v map[string]*string) *CreateContactResponse {
	s.Headers = v
	return s
}

func (s *CreateContactResponse) SetStatusCode(v int32) *CreateContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContactResponse) SetBody(v *CreateContactResponseBody) *CreateContactResponse {
	s.Body = v
	return s
}

func (s *CreateContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
