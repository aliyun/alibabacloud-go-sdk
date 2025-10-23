// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateEmailResponse
	GetStatusCode() *int32
	SetBody(v *ValidateEmailResponseBody) *ValidateEmailResponse
	GetBody() *ValidateEmailResponseBody
}

type ValidateEmailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateEmailResponse) GoString() string {
	return s.String()
}

func (s *ValidateEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateEmailResponse) GetBody() *ValidateEmailResponseBody {
	return s.Body
}

func (s *ValidateEmailResponse) SetHeaders(v map[string]*string) *ValidateEmailResponse {
	s.Headers = v
	return s
}

func (s *ValidateEmailResponse) SetStatusCode(v int32) *ValidateEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateEmailResponse) SetBody(v *ValidateEmailResponseBody) *ValidateEmailResponse {
	s.Body = v
	return s
}

func (s *ValidateEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
