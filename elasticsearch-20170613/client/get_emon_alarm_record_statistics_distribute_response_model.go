// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonAlarmRecordStatisticsDistributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmonAlarmRecordStatisticsDistributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmonAlarmRecordStatisticsDistributeResponse
	GetStatusCode() *int32
	SetBody(v *GetEmonAlarmRecordStatisticsDistributeResponseBody) *GetEmonAlarmRecordStatisticsDistributeResponse
	GetBody() *GetEmonAlarmRecordStatisticsDistributeResponseBody
}

type GetEmonAlarmRecordStatisticsDistributeResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmonAlarmRecordStatisticsDistributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmonAlarmRecordStatisticsDistributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeResponse) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponse) GetBody() *GetEmonAlarmRecordStatisticsDistributeResponseBody {
	return s.Body
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponse) SetHeaders(v map[string]*string) *GetEmonAlarmRecordStatisticsDistributeResponse {
	s.Headers = v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponse) SetStatusCode(v int32) *GetEmonAlarmRecordStatisticsDistributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponse) SetBody(v *GetEmonAlarmRecordStatisticsDistributeResponseBody) *GetEmonAlarmRecordStatisticsDistributeResponse {
	s.Body = v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponse) Validate() error {
	return dara.Validate(s)
}
