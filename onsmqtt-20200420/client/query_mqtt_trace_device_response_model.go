// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMqttTraceDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMqttTraceDeviceResponse
	GetStatusCode() *int32
	SetBody(v *QueryMqttTraceDeviceResponseBody) *QueryMqttTraceDeviceResponse
	GetBody() *QueryMqttTraceDeviceResponseBody
}

type QueryMqttTraceDeviceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMqttTraceDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMqttTraceDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMqttTraceDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMqttTraceDeviceResponse) GetBody() *QueryMqttTraceDeviceResponseBody {
	return s.Body
}

func (s *QueryMqttTraceDeviceResponse) SetHeaders(v map[string]*string) *QueryMqttTraceDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryMqttTraceDeviceResponse) SetStatusCode(v int32) *QueryMqttTraceDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqttTraceDeviceResponse) SetBody(v *QueryMqttTraceDeviceResponseBody) *QueryMqttTraceDeviceResponse {
	s.Body = v
	return s
}

func (s *QueryMqttTraceDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
