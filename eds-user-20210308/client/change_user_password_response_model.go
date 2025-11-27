// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ChangeUserPasswordResponseBody) *ChangeUserPasswordResponse
	GetBody() *ChangeUserPasswordResponseBody
}

type ChangeUserPasswordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangeUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeUserPasswordResponse) GetBody() *ChangeUserPasswordResponseBody {
	return s.Body
}

func (s *ChangeUserPasswordResponse) SetHeaders(v map[string]*string) *ChangeUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangeUserPasswordResponse) SetStatusCode(v int32) *ChangeUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeUserPasswordResponse) SetBody(v *ChangeUserPasswordResponseBody) *ChangeUserPasswordResponse {
	s.Body = v
	return s
}

func (s *ChangeUserPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
