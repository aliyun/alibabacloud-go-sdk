// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccessApplicationResponseBody) *UpdateAccessApplicationResponse
	GetBody() *UpdateAccessApplicationResponseBody
}

type UpdateAccessApplicationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccessApplicationResponse) GetBody() *UpdateAccessApplicationResponseBody {
	return s.Body
}

func (s *UpdateAccessApplicationResponse) SetHeaders(v map[string]*string) *UpdateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessApplicationResponse) SetStatusCode(v int32) *UpdateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessApplicationResponse) SetBody(v *UpdateAccessApplicationResponseBody) *UpdateAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *UpdateAccessApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
