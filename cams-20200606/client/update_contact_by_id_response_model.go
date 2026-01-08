// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContactByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContactByIdResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContactByIdResponseBody) *UpdateContactByIdResponse
	GetBody() *UpdateContactByIdResponseBody
}

type UpdateContactByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContactByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContactByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactByIdResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContactByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContactByIdResponse) GetBody() *UpdateContactByIdResponseBody {
	return s.Body
}

func (s *UpdateContactByIdResponse) SetHeaders(v map[string]*string) *UpdateContactByIdResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactByIdResponse) SetStatusCode(v int32) *UpdateContactByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactByIdResponse) SetBody(v *UpdateContactByIdResponseBody) *UpdateContactByIdResponse {
	s.Body = v
	return s
}

func (s *UpdateContactByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
