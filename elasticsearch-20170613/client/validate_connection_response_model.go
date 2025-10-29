// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ValidateConnectionResponseBody) *ValidateConnectionResponse
	GetBody() *ValidateConnectionResponseBody
}

type ValidateConnectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateConnectionResponse) GoString() string {
	return s.String()
}

func (s *ValidateConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateConnectionResponse) GetBody() *ValidateConnectionResponseBody {
	return s.Body
}

func (s *ValidateConnectionResponse) SetHeaders(v map[string]*string) *ValidateConnectionResponse {
	s.Headers = v
	return s
}

func (s *ValidateConnectionResponse) SetStatusCode(v int32) *ValidateConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateConnectionResponse) SetBody(v *ValidateConnectionResponseBody) *ValidateConnectionResponse {
	s.Body = v
	return s
}

func (s *ValidateConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
