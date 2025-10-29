// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComponentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComponentResponseBody) *UpdateComponentResponse
	GetBody() *UpdateComponentResponseBody
}

type UpdateComponentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentResponse) GoString() string {
	return s.String()
}

func (s *UpdateComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComponentResponse) GetBody() *UpdateComponentResponseBody {
	return s.Body
}

func (s *UpdateComponentResponse) SetHeaders(v map[string]*string) *UpdateComponentResponse {
	s.Headers = v
	return s
}

func (s *UpdateComponentResponse) SetStatusCode(v int32) *UpdateComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComponentResponse) SetBody(v *UpdateComponentResponseBody) *UpdateComponentResponse {
	s.Body = v
	return s
}

func (s *UpdateComponentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
