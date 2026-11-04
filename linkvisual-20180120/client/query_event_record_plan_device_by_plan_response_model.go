// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRecordPlanDeviceByPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEventRecordPlanDeviceByPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEventRecordPlanDeviceByPlanResponse
	GetStatusCode() *int32
	SetBody(v *QueryEventRecordPlanDeviceByPlanResponseBody) *QueryEventRecordPlanDeviceByPlanResponse
	GetBody() *QueryEventRecordPlanDeviceByPlanResponseBody
}

type QueryEventRecordPlanDeviceByPlanResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventRecordPlanDeviceByPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventRecordPlanDeviceByPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByPlanResponse) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) GetBody() *QueryEventRecordPlanDeviceByPlanResponseBody {
	return s.Body
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) SetHeaders(v map[string]*string) *QueryEventRecordPlanDeviceByPlanResponse {
	s.Headers = v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) SetStatusCode(v int32) *QueryEventRecordPlanDeviceByPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) SetBody(v *QueryEventRecordPlanDeviceByPlanResponseBody) *QueryEventRecordPlanDeviceByPlanResponse {
	s.Body = v
	return s
}

func (s *QueryEventRecordPlanDeviceByPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
