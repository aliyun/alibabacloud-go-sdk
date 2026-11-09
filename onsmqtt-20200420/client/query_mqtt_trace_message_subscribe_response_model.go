// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessageSubscribeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMqttTraceMessageSubscribeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMqttTraceMessageSubscribeResponse
	GetStatusCode() *int32
	SetBody(v *QueryMqttTraceMessageSubscribeResponseBody) *QueryMqttTraceMessageSubscribeResponse
	GetBody() *QueryMqttTraceMessageSubscribeResponseBody
}

type QueryMqttTraceMessageSubscribeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMqttTraceMessageSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMqttTraceMessageSubscribeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeResponse) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMqttTraceMessageSubscribeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMqttTraceMessageSubscribeResponse) GetBody() *QueryMqttTraceMessageSubscribeResponseBody {
	return s.Body
}

func (s *QueryMqttTraceMessageSubscribeResponse) SetHeaders(v map[string]*string) *QueryMqttTraceMessageSubscribeResponse {
	s.Headers = v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponse) SetStatusCode(v int32) *QueryMqttTraceMessageSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponse) SetBody(v *QueryMqttTraceMessageSubscribeResponseBody) *QueryMqttTraceMessageSubscribeResponse {
	s.Body = v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
