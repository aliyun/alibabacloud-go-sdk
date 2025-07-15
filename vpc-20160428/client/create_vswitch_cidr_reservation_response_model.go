// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVSwitchCidrReservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVSwitchCidrReservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVSwitchCidrReservationResponse
	GetStatusCode() *int32
	SetBody(v *CreateVSwitchCidrReservationResponseBody) *CreateVSwitchCidrReservationResponse
	GetBody() *CreateVSwitchCidrReservationResponseBody
}

type CreateVSwitchCidrReservationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVSwitchCidrReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVSwitchCidrReservationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchCidrReservationResponse) GoString() string {
	return s.String()
}

func (s *CreateVSwitchCidrReservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVSwitchCidrReservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVSwitchCidrReservationResponse) GetBody() *CreateVSwitchCidrReservationResponseBody {
	return s.Body
}

func (s *CreateVSwitchCidrReservationResponse) SetHeaders(v map[string]*string) *CreateVSwitchCidrReservationResponse {
	s.Headers = v
	return s
}

func (s *CreateVSwitchCidrReservationResponse) SetStatusCode(v int32) *CreateVSwitchCidrReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVSwitchCidrReservationResponse) SetBody(v *CreateVSwitchCidrReservationResponseBody) *CreateVSwitchCidrReservationResponse {
	s.Body = v
	return s
}

func (s *CreateVSwitchCidrReservationResponse) Validate() error {
	return dara.Validate(s)
}
