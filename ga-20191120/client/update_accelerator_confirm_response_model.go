// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorConfirmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAcceleratorConfirmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAcceleratorConfirmResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAcceleratorConfirmResponseBody) *UpdateAcceleratorConfirmResponse
	GetBody() *UpdateAcceleratorConfirmResponseBody
}

type UpdateAcceleratorConfirmResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAcceleratorConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAcceleratorConfirmResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorConfirmResponse) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorConfirmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAcceleratorConfirmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAcceleratorConfirmResponse) GetBody() *UpdateAcceleratorConfirmResponseBody {
	return s.Body
}

func (s *UpdateAcceleratorConfirmResponse) SetHeaders(v map[string]*string) *UpdateAcceleratorConfirmResponse {
	s.Headers = v
	return s
}

func (s *UpdateAcceleratorConfirmResponse) SetStatusCode(v int32) *UpdateAcceleratorConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAcceleratorConfirmResponse) SetBody(v *UpdateAcceleratorConfirmResponseBody) *UpdateAcceleratorConfirmResponse {
	s.Body = v
	return s
}

func (s *UpdateAcceleratorConfirmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
