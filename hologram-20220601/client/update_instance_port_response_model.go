// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstancePortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstancePortResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstancePortResponseBody) *UpdateInstancePortResponse
	GetBody() *UpdateInstancePortResponseBody
}

type UpdateInstancePortResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstancePortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstancePortResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePortResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstancePortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstancePortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstancePortResponse) GetBody() *UpdateInstancePortResponseBody {
	return s.Body
}

func (s *UpdateInstancePortResponse) SetHeaders(v map[string]*string) *UpdateInstancePortResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstancePortResponse) SetStatusCode(v int32) *UpdateInstancePortResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstancePortResponse) SetBody(v *UpdateInstancePortResponseBody) *UpdateInstancePortResponse {
	s.Body = v
	return s
}

func (s *UpdateInstancePortResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
