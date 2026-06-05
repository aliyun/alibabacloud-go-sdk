// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSeoStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppSeoStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppSeoStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppSeoStatusResponseBody) *UpdateAppSeoStatusResponse
	GetBody() *UpdateAppSeoStatusResponseBody
}

type UpdateAppSeoStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppSeoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppSeoStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSeoStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppSeoStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppSeoStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppSeoStatusResponse) GetBody() *UpdateAppSeoStatusResponseBody {
	return s.Body
}

func (s *UpdateAppSeoStatusResponse) SetHeaders(v map[string]*string) *UpdateAppSeoStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppSeoStatusResponse) SetStatusCode(v int32) *UpdateAppSeoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppSeoStatusResponse) SetBody(v *UpdateAppSeoStatusResponseBody) *UpdateAppSeoStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAppSeoStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
