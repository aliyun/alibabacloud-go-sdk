// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorCrossBorderStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAcceleratorCrossBorderStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAcceleratorCrossBorderStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAcceleratorCrossBorderStatusResponseBody) *UpdateAcceleratorCrossBorderStatusResponse
	GetBody() *UpdateAcceleratorCrossBorderStatusResponseBody
}

type UpdateAcceleratorCrossBorderStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAcceleratorCrossBorderStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAcceleratorCrossBorderStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorCrossBorderStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorCrossBorderStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAcceleratorCrossBorderStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAcceleratorCrossBorderStatusResponse) GetBody() *UpdateAcceleratorCrossBorderStatusResponseBody {
	return s.Body
}

func (s *UpdateAcceleratorCrossBorderStatusResponse) SetHeaders(v map[string]*string) *UpdateAcceleratorCrossBorderStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAcceleratorCrossBorderStatusResponse) SetStatusCode(v int32) *UpdateAcceleratorCrossBorderStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderStatusResponse) SetBody(v *UpdateAcceleratorCrossBorderStatusResponseBody) *UpdateAcceleratorCrossBorderStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAcceleratorCrossBorderStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
