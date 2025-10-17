// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBindRamUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckBindRamUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckBindRamUserResponse
	GetStatusCode() *int32
	SetBody(v *CheckBindRamUserResponseBody) *CheckBindRamUserResponse
	GetBody() *CheckBindRamUserResponseBody
}

type CheckBindRamUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckBindRamUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckBindRamUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckBindRamUserResponse) GoString() string {
	return s.String()
}

func (s *CheckBindRamUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckBindRamUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckBindRamUserResponse) GetBody() *CheckBindRamUserResponseBody {
	return s.Body
}

func (s *CheckBindRamUserResponse) SetHeaders(v map[string]*string) *CheckBindRamUserResponse {
	s.Headers = v
	return s
}

func (s *CheckBindRamUserResponse) SetStatusCode(v int32) *CheckBindRamUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckBindRamUserResponse) SetBody(v *CheckBindRamUserResponseBody) *CheckBindRamUserResponse {
	s.Body = v
	return s
}

func (s *CheckBindRamUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
