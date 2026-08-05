// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCapabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCapabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCapabilityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCapabilityResponseBody) *UpdateCapabilityResponse
	GetBody() *UpdateCapabilityResponseBody
}

type UpdateCapabilityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCapabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCapabilityResponse) GoString() string {
	return s.String()
}

func (s *UpdateCapabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCapabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCapabilityResponse) GetBody() *UpdateCapabilityResponseBody {
	return s.Body
}

func (s *UpdateCapabilityResponse) SetHeaders(v map[string]*string) *UpdateCapabilityResponse {
	s.Headers = v
	return s
}

func (s *UpdateCapabilityResponse) SetStatusCode(v int32) *UpdateCapabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCapabilityResponse) SetBody(v *UpdateCapabilityResponseBody) *UpdateCapabilityResponse {
	s.Body = v
	return s
}

func (s *UpdateCapabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
