// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLgfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLgfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLgfResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLgfResponseBody) *UpdateLgfResponse
	GetBody() *UpdateLgfResponseBody
}

type UpdateLgfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLgfResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLgfResponse) GoString() string {
	return s.String()
}

func (s *UpdateLgfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLgfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLgfResponse) GetBody() *UpdateLgfResponseBody {
	return s.Body
}

func (s *UpdateLgfResponse) SetHeaders(v map[string]*string) *UpdateLgfResponse {
	s.Headers = v
	return s
}

func (s *UpdateLgfResponse) SetStatusCode(v int32) *UpdateLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLgfResponse) SetBody(v *UpdateLgfResponseBody) *UpdateLgfResponse {
	s.Body = v
	return s
}

func (s *UpdateLgfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
