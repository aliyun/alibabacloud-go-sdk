// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ValidateApplicationResponseBody) *ValidateApplicationResponse
	GetBody() *ValidateApplicationResponseBody
}

type ValidateApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateApplicationResponse) GoString() string {
	return s.String()
}

func (s *ValidateApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateApplicationResponse) GetBody() *ValidateApplicationResponseBody {
	return s.Body
}

func (s *ValidateApplicationResponse) SetHeaders(v map[string]*string) *ValidateApplicationResponse {
	s.Headers = v
	return s
}

func (s *ValidateApplicationResponse) SetStatusCode(v int32) *ValidateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateApplicationResponse) SetBody(v *ValidateApplicationResponseBody) *ValidateApplicationResponse {
	s.Body = v
	return s
}

func (s *ValidateApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
