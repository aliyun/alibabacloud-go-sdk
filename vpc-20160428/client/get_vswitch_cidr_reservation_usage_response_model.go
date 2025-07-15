// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVSwitchCidrReservationUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVSwitchCidrReservationUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVSwitchCidrReservationUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetVSwitchCidrReservationUsageResponseBody) *GetVSwitchCidrReservationUsageResponse
	GetBody() *GetVSwitchCidrReservationUsageResponseBody
}

type GetVSwitchCidrReservationUsageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVSwitchCidrReservationUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVSwitchCidrReservationUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVSwitchCidrReservationUsageResponse) GoString() string {
	return s.String()
}

func (s *GetVSwitchCidrReservationUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVSwitchCidrReservationUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVSwitchCidrReservationUsageResponse) GetBody() *GetVSwitchCidrReservationUsageResponseBody {
	return s.Body
}

func (s *GetVSwitchCidrReservationUsageResponse) SetHeaders(v map[string]*string) *GetVSwitchCidrReservationUsageResponse {
	s.Headers = v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponse) SetStatusCode(v int32) *GetVSwitchCidrReservationUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponse) SetBody(v *GetVSwitchCidrReservationUsageResponseBody) *GetVSwitchCidrReservationUsageResponse {
	s.Body = v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponse) Validate() error {
	return dara.Validate(s)
}
