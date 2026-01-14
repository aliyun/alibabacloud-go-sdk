// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorCrossBorderModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAcceleratorCrossBorderModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAcceleratorCrossBorderModeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAcceleratorCrossBorderModeResponseBody) *UpdateAcceleratorCrossBorderModeResponse
	GetBody() *UpdateAcceleratorCrossBorderModeResponseBody
}

type UpdateAcceleratorCrossBorderModeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAcceleratorCrossBorderModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAcceleratorCrossBorderModeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorCrossBorderModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorCrossBorderModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAcceleratorCrossBorderModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAcceleratorCrossBorderModeResponse) GetBody() *UpdateAcceleratorCrossBorderModeResponseBody {
	return s.Body
}

func (s *UpdateAcceleratorCrossBorderModeResponse) SetHeaders(v map[string]*string) *UpdateAcceleratorCrossBorderModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAcceleratorCrossBorderModeResponse) SetStatusCode(v int32) *UpdateAcceleratorCrossBorderModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderModeResponse) SetBody(v *UpdateAcceleratorCrossBorderModeResponseBody) *UpdateAcceleratorCrossBorderModeResponse {
	s.Body = v
	return s
}

func (s *UpdateAcceleratorCrossBorderModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
