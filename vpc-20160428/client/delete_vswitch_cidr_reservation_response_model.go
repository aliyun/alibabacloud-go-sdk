// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVSwitchCidrReservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVSwitchCidrReservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVSwitchCidrReservationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVSwitchCidrReservationResponseBody) *DeleteVSwitchCidrReservationResponse
	GetBody() *DeleteVSwitchCidrReservationResponseBody
}

type DeleteVSwitchCidrReservationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVSwitchCidrReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVSwitchCidrReservationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVSwitchCidrReservationResponse) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchCidrReservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVSwitchCidrReservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVSwitchCidrReservationResponse) GetBody() *DeleteVSwitchCidrReservationResponseBody {
	return s.Body
}

func (s *DeleteVSwitchCidrReservationResponse) SetHeaders(v map[string]*string) *DeleteVSwitchCidrReservationResponse {
	s.Headers = v
	return s
}

func (s *DeleteVSwitchCidrReservationResponse) SetStatusCode(v int32) *DeleteVSwitchCidrReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVSwitchCidrReservationResponse) SetBody(v *DeleteVSwitchCidrReservationResponseBody) *DeleteVSwitchCidrReservationResponse {
	s.Body = v
	return s
}

func (s *DeleteVSwitchCidrReservationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
