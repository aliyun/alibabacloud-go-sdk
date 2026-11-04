// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRecordPlanDeviceByDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEventRecordPlanDeviceByDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEventRecordPlanDeviceByDeviceResponse
	GetStatusCode() *int32
	SetBody(v *QueryEventRecordPlanDeviceByDeviceResponseBody) *QueryEventRecordPlanDeviceByDeviceResponse
	GetBody() *QueryEventRecordPlanDeviceByDeviceResponseBody
}

type QueryEventRecordPlanDeviceByDeviceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventRecordPlanDeviceByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) GetBody() *QueryEventRecordPlanDeviceByDeviceResponseBody {
	return s.Body
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) SetHeaders(v map[string]*string) *QueryEventRecordPlanDeviceByDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) SetStatusCode(v int32) *QueryEventRecordPlanDeviceByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) SetBody(v *QueryEventRecordPlanDeviceByDeviceResponseBody) *QueryEventRecordPlanDeviceByDeviceResponse {
	s.Body = v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
