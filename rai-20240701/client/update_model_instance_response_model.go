// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelInstanceResponseBody) *UpdateModelInstanceResponse
	GetBody() *UpdateModelInstanceResponseBody
}

type UpdateModelInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelInstanceResponse) GetBody() *UpdateModelInstanceResponseBody {
	return s.Body
}

func (s *UpdateModelInstanceResponse) SetHeaders(v map[string]*string) *UpdateModelInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelInstanceResponse) SetStatusCode(v int32) *UpdateModelInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelInstanceResponse) SetBody(v *UpdateModelInstanceResponseBody) *UpdateModelInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateModelInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
