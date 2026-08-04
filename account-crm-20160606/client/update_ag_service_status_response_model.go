// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgServiceStatusResponseBody) *UpdateAgServiceStatusResponse
	GetBody() *UpdateAgServiceStatusResponseBody
}

type UpdateAgServiceStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgServiceStatusResponse) GetBody() *UpdateAgServiceStatusResponseBody {
	return s.Body
}

func (s *UpdateAgServiceStatusResponse) SetHeaders(v map[string]*string) *UpdateAgServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgServiceStatusResponse) SetStatusCode(v int32) *UpdateAgServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgServiceStatusResponse) SetBody(v *UpdateAgServiceStatusResponseBody) *UpdateAgServiceStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAgServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
