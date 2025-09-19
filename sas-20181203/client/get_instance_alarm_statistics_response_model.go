// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAlarmStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceAlarmStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceAlarmStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceAlarmStatisticsResponseBody) *GetInstanceAlarmStatisticsResponse
	GetBody() *GetInstanceAlarmStatisticsResponseBody
}

type GetInstanceAlarmStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceAlarmStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceAlarmStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAlarmStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceAlarmStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceAlarmStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceAlarmStatisticsResponse) GetBody() *GetInstanceAlarmStatisticsResponseBody {
	return s.Body
}

func (s *GetInstanceAlarmStatisticsResponse) SetHeaders(v map[string]*string) *GetInstanceAlarmStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceAlarmStatisticsResponse) SetStatusCode(v int32) *GetInstanceAlarmStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceAlarmStatisticsResponse) SetBody(v *GetInstanceAlarmStatisticsResponseBody) *GetInstanceAlarmStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetInstanceAlarmStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
