// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBlackIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBlackIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBlackIpsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBlackIpsResponseBody) *UpdateBlackIpsResponse
	GetBody() *UpdateBlackIpsResponseBody
}

type UpdateBlackIpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBlackIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBlackIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBlackIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdateBlackIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBlackIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBlackIpsResponse) GetBody() *UpdateBlackIpsResponseBody {
	return s.Body
}

func (s *UpdateBlackIpsResponse) SetHeaders(v map[string]*string) *UpdateBlackIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdateBlackIpsResponse) SetStatusCode(v int32) *UpdateBlackIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBlackIpsResponse) SetBody(v *UpdateBlackIpsResponseBody) *UpdateBlackIpsResponse {
	s.Body = v
	return s
}

func (s *UpdateBlackIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
