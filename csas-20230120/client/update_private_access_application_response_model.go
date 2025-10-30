// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrivateAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrivateAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrivateAccessApplicationResponseBody) *UpdatePrivateAccessApplicationResponse
	GetBody() *UpdatePrivateAccessApplicationResponseBody
}

type UpdatePrivateAccessApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrivateAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrivateAccessApplicationResponse) GetBody() *UpdatePrivateAccessApplicationResponseBody {
	return s.Body
}

func (s *UpdatePrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *UpdatePrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateAccessApplicationResponse) SetStatusCode(v int32) *UpdatePrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateAccessApplicationResponse) SetBody(v *UpdatePrivateAccessApplicationResponseBody) *UpdatePrivateAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *UpdatePrivateAccessApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
