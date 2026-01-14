// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAcceleratorResponseBody) *UpdateAcceleratorResponse
	GetBody() *UpdateAcceleratorResponseBody
}

type UpdateAcceleratorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAcceleratorResponse) GetBody() *UpdateAcceleratorResponseBody {
	return s.Body
}

func (s *UpdateAcceleratorResponse) SetHeaders(v map[string]*string) *UpdateAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *UpdateAcceleratorResponse) SetStatusCode(v int32) *UpdateAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAcceleratorResponse) SetBody(v *UpdateAcceleratorResponseBody) *UpdateAcceleratorResponse {
	s.Body = v
	return s
}

func (s *UpdateAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
