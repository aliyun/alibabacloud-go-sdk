// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceNameResponseBody) *UpdateInstanceNameResponse
	GetBody() *UpdateInstanceNameResponseBody
}

type UpdateInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceNameResponse) GetBody() *UpdateInstanceNameResponseBody {
	return s.Body
}

func (s *UpdateInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNameResponse) SetStatusCode(v int32) *UpdateInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponse) SetBody(v *UpdateInstanceNameResponseBody) *UpdateInstanceNameResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
