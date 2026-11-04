// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordPlanDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecordPlanDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecordPlanDeviceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecordPlanDeviceResponseBody) *DeleteRecordPlanDeviceResponse
	GetBody() *DeleteRecordPlanDeviceResponseBody
}

type DeleteRecordPlanDeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordPlanDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecordPlanDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordPlanDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecordPlanDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecordPlanDeviceResponse) GetBody() *DeleteRecordPlanDeviceResponseBody {
	return s.Body
}

func (s *DeleteRecordPlanDeviceResponse) SetHeaders(v map[string]*string) *DeleteRecordPlanDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordPlanDeviceResponse) SetStatusCode(v int32) *DeleteRecordPlanDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponse) SetBody(v *DeleteRecordPlanDeviceResponseBody) *DeleteRecordPlanDeviceResponse {
	s.Body = v
	return s
}

func (s *DeleteRecordPlanDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
