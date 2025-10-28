// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCodeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCodeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCodeSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCodeSourceResponseBody) *UpdateCodeSourceResponse
	GetBody() *UpdateCodeSourceResponseBody
}

type UpdateCodeSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCodeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateCodeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCodeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCodeSourceResponse) GetBody() *UpdateCodeSourceResponseBody {
	return s.Body
}

func (s *UpdateCodeSourceResponse) SetHeaders(v map[string]*string) *UpdateCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateCodeSourceResponse) SetStatusCode(v int32) *UpdateCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCodeSourceResponse) SetBody(v *UpdateCodeSourceResponseBody) *UpdateCodeSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateCodeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
