// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEventRecordPlanDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddEventRecordPlanDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddEventRecordPlanDeviceResponse
	GetStatusCode() *int32
	SetBody(v *AddEventRecordPlanDeviceResponseBody) *AddEventRecordPlanDeviceResponse
	GetBody() *AddEventRecordPlanDeviceResponseBody
}

type AddEventRecordPlanDeviceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEventRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEventRecordPlanDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddEventRecordPlanDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddEventRecordPlanDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddEventRecordPlanDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddEventRecordPlanDeviceResponse) GetBody() *AddEventRecordPlanDeviceResponseBody {
	return s.Body
}

func (s *AddEventRecordPlanDeviceResponse) SetHeaders(v map[string]*string) *AddEventRecordPlanDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddEventRecordPlanDeviceResponse) SetStatusCode(v int32) *AddEventRecordPlanDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEventRecordPlanDeviceResponse) SetBody(v *AddEventRecordPlanDeviceResponseBody) *AddEventRecordPlanDeviceResponse {
	s.Body = v
	return s
}

func (s *AddEventRecordPlanDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
