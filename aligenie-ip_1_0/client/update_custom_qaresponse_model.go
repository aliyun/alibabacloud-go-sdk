// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomQAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomQAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomQAResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomQAResponseBody) *UpdateCustomQAResponse
	GetBody() *UpdateCustomQAResponseBody
}

type UpdateCustomQAResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomQAResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomQAResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomQAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomQAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomQAResponse) GetBody() *UpdateCustomQAResponseBody {
	return s.Body
}

func (s *UpdateCustomQAResponse) SetHeaders(v map[string]*string) *UpdateCustomQAResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomQAResponse) SetStatusCode(v int32) *UpdateCustomQAResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomQAResponse) SetBody(v *UpdateCustomQAResponseBody) *UpdateCustomQAResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomQAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
