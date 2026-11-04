// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRecordPlanDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventRecordPlanDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventRecordPlanDeviceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventRecordPlanDeviceResponseBody) *DeleteEventRecordPlanDeviceResponse
	GetBody() *DeleteEventRecordPlanDeviceResponseBody
}

type DeleteEventRecordPlanDeviceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventRecordPlanDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRecordPlanDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventRecordPlanDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventRecordPlanDeviceResponse) GetBody() *DeleteEventRecordPlanDeviceResponseBody {
	return s.Body
}

func (s *DeleteEventRecordPlanDeviceResponse) SetHeaders(v map[string]*string) *DeleteEventRecordPlanDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponse) SetStatusCode(v int32) *DeleteEventRecordPlanDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponse) SetBody(v *DeleteEventRecordPlanDeviceResponseBody) *DeleteEventRecordPlanDeviceResponse {
	s.Body = v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
