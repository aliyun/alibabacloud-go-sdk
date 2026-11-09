// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessageOfClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMqttTraceMessageOfClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMqttTraceMessageOfClientResponse
	GetStatusCode() *int32
	SetBody(v *QueryMqttTraceMessageOfClientResponseBody) *QueryMqttTraceMessageOfClientResponse
	GetBody() *QueryMqttTraceMessageOfClientResponseBody
}

type QueryMqttTraceMessageOfClientResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMqttTraceMessageOfClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMqttTraceMessageOfClientResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientResponse) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMqttTraceMessageOfClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMqttTraceMessageOfClientResponse) GetBody() *QueryMqttTraceMessageOfClientResponseBody {
	return s.Body
}

func (s *QueryMqttTraceMessageOfClientResponse) SetHeaders(v map[string]*string) *QueryMqttTraceMessageOfClientResponse {
	s.Headers = v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponse) SetStatusCode(v int32) *QueryMqttTraceMessageOfClientResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponse) SetBody(v *QueryMqttTraceMessageOfClientResponseBody) *QueryMqttTraceMessageOfClientResponse {
	s.Body = v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
