// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContactResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContactResponseBody) *UpdateContactResponse
	GetBody() *UpdateContactResponseBody
}

type UpdateContactResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContactResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContactResponse) GetBody() *UpdateContactResponseBody {
	return s.Body
}

func (s *UpdateContactResponse) SetHeaders(v map[string]*string) *UpdateContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactResponse) SetStatusCode(v int32) *UpdateContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactResponse) SetBody(v *UpdateContactResponseBody) *UpdateContactResponse {
	s.Body = v
	return s
}

func (s *UpdateContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
