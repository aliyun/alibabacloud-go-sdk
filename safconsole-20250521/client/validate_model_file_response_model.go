// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateModelFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateModelFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateModelFileResponse
	GetStatusCode() *int32
	SetBody(v *ValidateModelFileResponseBody) *ValidateModelFileResponse
	GetBody() *ValidateModelFileResponseBody
}

type ValidateModelFileResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateModelFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateModelFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateModelFileResponse) GoString() string {
	return s.String()
}

func (s *ValidateModelFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateModelFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateModelFileResponse) GetBody() *ValidateModelFileResponseBody {
	return s.Body
}

func (s *ValidateModelFileResponse) SetHeaders(v map[string]*string) *ValidateModelFileResponse {
	s.Headers = v
	return s
}

func (s *ValidateModelFileResponse) SetStatusCode(v int32) *ValidateModelFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateModelFileResponse) SetBody(v *ValidateModelFileResponseBody) *ValidateModelFileResponse {
	s.Body = v
	return s
}

func (s *ValidateModelFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
