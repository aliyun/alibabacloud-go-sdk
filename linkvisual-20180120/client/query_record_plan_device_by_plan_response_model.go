// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlanDeviceByPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordPlanDeviceByPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordPlanDeviceByPlanResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordPlanDeviceByPlanResponseBody) *QueryRecordPlanDeviceByPlanResponse
	GetBody() *QueryRecordPlanDeviceByPlanResponseBody
}

type QueryRecordPlanDeviceByPlanResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordPlanDeviceByPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordPlanDeviceByPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlanDeviceByPlanResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordPlanDeviceByPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordPlanDeviceByPlanResponse) GetBody() *QueryRecordPlanDeviceByPlanResponseBody {
	return s.Body
}

func (s *QueryRecordPlanDeviceByPlanResponse) SetHeaders(v map[string]*string) *QueryRecordPlanDeviceByPlanResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponse) SetStatusCode(v int32) *QueryRecordPlanDeviceByPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponse) SetBody(v *QueryRecordPlanDeviceByPlanResponseBody) *QueryRecordPlanDeviceByPlanResponse {
	s.Body = v
	return s
}

func (s *QueryRecordPlanDeviceByPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
