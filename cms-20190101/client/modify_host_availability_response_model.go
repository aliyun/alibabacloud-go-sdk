// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostAvailabilityResponseBody) *ModifyHostAvailabilityResponse
	GetBody() *ModifyHostAvailabilityResponseBody
}

type ModifyHostAvailabilityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostAvailabilityResponse) GetBody() *ModifyHostAvailabilityResponseBody {
	return s.Body
}

func (s *ModifyHostAvailabilityResponse) SetHeaders(v map[string]*string) *ModifyHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostAvailabilityResponse) SetStatusCode(v int32) *ModifyHostAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostAvailabilityResponse) SetBody(v *ModifyHostAvailabilityResponseBody) *ModifyHostAvailabilityResponse {
	s.Body = v
	return s
}

func (s *ModifyHostAvailabilityResponse) Validate() error {
	return dara.Validate(s)
}
