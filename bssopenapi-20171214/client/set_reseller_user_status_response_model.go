// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetResellerUserStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetResellerUserStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetResellerUserStatusResponseBody) *SetResellerUserStatusResponse
	GetBody() *SetResellerUserStatusResponseBody
}

type SetResellerUserStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetResellerUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetResellerUserStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserStatusResponse) GoString() string {
	return s.String()
}

func (s *SetResellerUserStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetResellerUserStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetResellerUserStatusResponse) GetBody() *SetResellerUserStatusResponseBody {
	return s.Body
}

func (s *SetResellerUserStatusResponse) SetHeaders(v map[string]*string) *SetResellerUserStatusResponse {
	s.Headers = v
	return s
}

func (s *SetResellerUserStatusResponse) SetStatusCode(v int32) *SetResellerUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetResellerUserStatusResponse) SetBody(v *SetResellerUserStatusResponseBody) *SetResellerUserStatusResponse {
	s.Body = v
	return s
}

func (s *SetResellerUserStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
