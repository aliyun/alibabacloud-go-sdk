// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceChargeTypeResponseBody) *UpdateInstanceChargeTypeResponse
	GetBody() *UpdateInstanceChargeTypeResponseBody
}

type UpdateInstanceChargeTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceChargeTypeResponse) GetBody() *UpdateInstanceChargeTypeResponseBody {
	return s.Body
}

func (s *UpdateInstanceChargeTypeResponse) SetHeaders(v map[string]*string) *UpdateInstanceChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceChargeTypeResponse) SetStatusCode(v int32) *UpdateInstanceChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceChargeTypeResponse) SetBody(v *UpdateInstanceChargeTypeResponseBody) *UpdateInstanceChargeTypeResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
