// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHostAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableHostAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableHostAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *DisableHostAvailabilityResponseBody) *DisableHostAvailabilityResponse
	GetBody() *DisableHostAvailabilityResponseBody
}

type DisableHostAvailabilityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableHostAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *DisableHostAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableHostAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableHostAvailabilityResponse) GetBody() *DisableHostAvailabilityResponseBody {
	return s.Body
}

func (s *DisableHostAvailabilityResponse) SetHeaders(v map[string]*string) *DisableHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *DisableHostAvailabilityResponse) SetStatusCode(v int32) *DisableHostAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableHostAvailabilityResponse) SetBody(v *DisableHostAvailabilityResponseBody) *DisableHostAvailabilityResponse {
	s.Body = v
	return s
}

func (s *DisableHostAvailabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
