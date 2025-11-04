// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSourceResponseBody) *UpdateSourceResponse
	GetBody() *UpdateSourceResponseBody
}

type UpdateSourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSourceResponse) GetBody() *UpdateSourceResponseBody {
	return s.Body
}

func (s *UpdateSourceResponse) SetHeaders(v map[string]*string) *UpdateSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateSourceResponse) SetStatusCode(v int32) *UpdateSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSourceResponse) SetBody(v *UpdateSourceResponseBody) *UpdateSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
