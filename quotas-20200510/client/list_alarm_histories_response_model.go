// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlarmHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlarmHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAlarmHistoriesResponseBody) *ListAlarmHistoriesResponse
	GetBody() *ListAlarmHistoriesResponseBody
}

type ListAlarmHistoriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlarmHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlarmHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlarmHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlarmHistoriesResponse) GetBody() *ListAlarmHistoriesResponseBody {
	return s.Body
}

func (s *ListAlarmHistoriesResponse) SetHeaders(v map[string]*string) *ListAlarmHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmHistoriesResponse) SetStatusCode(v int32) *ListAlarmHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlarmHistoriesResponse) SetBody(v *ListAlarmHistoriesResponseBody) *ListAlarmHistoriesResponse {
	s.Body = v
	return s
}

func (s *ListAlarmHistoriesResponse) Validate() error {
	return dara.Validate(s)
}
