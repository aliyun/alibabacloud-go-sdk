// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVSwitchCidrReservationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVSwitchCidrReservationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVSwitchCidrReservationsResponse
	GetStatusCode() *int32
	SetBody(v *ListVSwitchCidrReservationsResponseBody) *ListVSwitchCidrReservationsResponse
	GetBody() *ListVSwitchCidrReservationsResponseBody
}

type ListVSwitchCidrReservationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVSwitchCidrReservationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVSwitchCidrReservationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVSwitchCidrReservationsResponse) GoString() string {
	return s.String()
}

func (s *ListVSwitchCidrReservationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVSwitchCidrReservationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVSwitchCidrReservationsResponse) GetBody() *ListVSwitchCidrReservationsResponseBody {
	return s.Body
}

func (s *ListVSwitchCidrReservationsResponse) SetHeaders(v map[string]*string) *ListVSwitchCidrReservationsResponse {
	s.Headers = v
	return s
}

func (s *ListVSwitchCidrReservationsResponse) SetStatusCode(v int32) *ListVSwitchCidrReservationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponse) SetBody(v *ListVSwitchCidrReservationsResponseBody) *ListVSwitchCidrReservationsResponse {
	s.Body = v
	return s
}

func (s *ListVSwitchCidrReservationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
