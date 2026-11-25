// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScalingGroupZoneShiftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelScalingGroupZoneShiftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelScalingGroupZoneShiftResponse
	GetStatusCode() *int32
	SetBody(v *CancelScalingGroupZoneShiftResponseBody) *CancelScalingGroupZoneShiftResponse
	GetBody() *CancelScalingGroupZoneShiftResponseBody
}

type CancelScalingGroupZoneShiftResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelScalingGroupZoneShiftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelScalingGroupZoneShiftResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelScalingGroupZoneShiftResponse) GoString() string {
	return s.String()
}

func (s *CancelScalingGroupZoneShiftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelScalingGroupZoneShiftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelScalingGroupZoneShiftResponse) GetBody() *CancelScalingGroupZoneShiftResponseBody {
	return s.Body
}

func (s *CancelScalingGroupZoneShiftResponse) SetHeaders(v map[string]*string) *CancelScalingGroupZoneShiftResponse {
	s.Headers = v
	return s
}

func (s *CancelScalingGroupZoneShiftResponse) SetStatusCode(v int32) *CancelScalingGroupZoneShiftResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelScalingGroupZoneShiftResponse) SetBody(v *CancelScalingGroupZoneShiftResponseBody) *CancelScalingGroupZoneShiftResponse {
	s.Body = v
	return s
}

func (s *CancelScalingGroupZoneShiftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
