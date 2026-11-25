// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartScalingGroupZoneShiftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartScalingGroupZoneShiftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartScalingGroupZoneShiftResponse
	GetStatusCode() *int32
	SetBody(v *StartScalingGroupZoneShiftResponseBody) *StartScalingGroupZoneShiftResponse
	GetBody() *StartScalingGroupZoneShiftResponseBody
}

type StartScalingGroupZoneShiftResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartScalingGroupZoneShiftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartScalingGroupZoneShiftResponse) String() string {
	return dara.Prettify(s)
}

func (s StartScalingGroupZoneShiftResponse) GoString() string {
	return s.String()
}

func (s *StartScalingGroupZoneShiftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartScalingGroupZoneShiftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartScalingGroupZoneShiftResponse) GetBody() *StartScalingGroupZoneShiftResponseBody {
	return s.Body
}

func (s *StartScalingGroupZoneShiftResponse) SetHeaders(v map[string]*string) *StartScalingGroupZoneShiftResponse {
	s.Headers = v
	return s
}

func (s *StartScalingGroupZoneShiftResponse) SetStatusCode(v int32) *StartScalingGroupZoneShiftResponse {
	s.StatusCode = &v
	return s
}

func (s *StartScalingGroupZoneShiftResponse) SetBody(v *StartScalingGroupZoneShiftResponseBody) *StartScalingGroupZoneShiftResponse {
	s.Body = v
	return s
}

func (s *StartScalingGroupZoneShiftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
