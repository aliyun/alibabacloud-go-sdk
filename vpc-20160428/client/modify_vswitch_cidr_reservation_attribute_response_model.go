// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchCidrReservationAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVSwitchCidrReservationAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVSwitchCidrReservationAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVSwitchCidrReservationAttributeResponseBody) *ModifyVSwitchCidrReservationAttributeResponse
	GetBody() *ModifyVSwitchCidrReservationAttributeResponseBody
}

type ModifyVSwitchCidrReservationAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVSwitchCidrReservationAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVSwitchCidrReservationAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchCidrReservationAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchCidrReservationAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVSwitchCidrReservationAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVSwitchCidrReservationAttributeResponse) GetBody() *ModifyVSwitchCidrReservationAttributeResponseBody {
	return s.Body
}

func (s *ModifyVSwitchCidrReservationAttributeResponse) SetHeaders(v map[string]*string) *ModifyVSwitchCidrReservationAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeResponse) SetStatusCode(v int32) *ModifyVSwitchCidrReservationAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeResponse) SetBody(v *ModifyVSwitchCidrReservationAttributeResponseBody) *ModifyVSwitchCidrReservationAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
