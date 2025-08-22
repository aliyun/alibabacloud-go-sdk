// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnSLSRealtimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDcdnSLSRealtimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDcdnSLSRealtimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) *UpdateDcdnSLSRealtimeLogDeliveryResponse
	GetBody() *UpdateDcdnSLSRealtimeLogDeliveryResponseBody
}

type UpdateDcdnSLSRealtimeLogDeliveryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDcdnSLSRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) GetBody() *UpdateDcdnSLSRealtimeLogDeliveryResponseBody {
	return s.Body
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *UpdateDcdnSLSRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) SetStatusCode(v int32) *UpdateDcdnSLSRealtimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) SetBody(v *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) *UpdateDcdnSLSRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
