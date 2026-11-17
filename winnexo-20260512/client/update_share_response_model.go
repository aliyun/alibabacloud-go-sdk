// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateShareResponse
	GetStatusCode() *int32
	SetBody(v *UpdateShareResponseBody) *UpdateShareResponse
	GetBody() *UpdateShareResponseBody
}

type UpdateShareResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateShareResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateShareResponse) GoString() string {
	return s.String()
}

func (s *UpdateShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateShareResponse) GetBody() *UpdateShareResponseBody {
	return s.Body
}

func (s *UpdateShareResponse) SetHeaders(v map[string]*string) *UpdateShareResponse {
	s.Headers = v
	return s
}

func (s *UpdateShareResponse) SetStatusCode(v int32) *UpdateShareResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateShareResponse) SetBody(v *UpdateShareResponseBody) *UpdateShareResponse {
	s.Body = v
	return s
}

func (s *UpdateShareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
