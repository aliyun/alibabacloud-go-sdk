// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAcceleratorAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAcceleratorAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAcceleratorAutoRenewAttributeResponseBody) *UpdateAcceleratorAutoRenewAttributeResponse
	GetBody() *UpdateAcceleratorAutoRenewAttributeResponseBody
}

type UpdateAcceleratorAutoRenewAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAcceleratorAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAcceleratorAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) GetBody() *UpdateAcceleratorAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *UpdateAcceleratorAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) SetStatusCode(v int32) *UpdateAcceleratorAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) SetBody(v *UpdateAcceleratorAutoRenewAttributeResponseBody) *UpdateAcceleratorAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
