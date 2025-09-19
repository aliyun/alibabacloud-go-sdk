// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertStrategiesResponseBody) *ListAlertStrategiesResponse
	GetBody() *ListAlertStrategiesResponseBody
}

type ListAlertStrategiesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListAlertStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertStrategiesResponse) GetBody() *ListAlertStrategiesResponseBody {
	return s.Body
}

func (s *ListAlertStrategiesResponse) SetHeaders(v map[string]*string) *ListAlertStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListAlertStrategiesResponse) SetStatusCode(v int32) *ListAlertStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertStrategiesResponse) SetBody(v *ListAlertStrategiesResponseBody) *ListAlertStrategiesResponse {
	s.Body = v
	return s
}

func (s *ListAlertStrategiesResponse) Validate() error {
	return dara.Validate(s)
}
