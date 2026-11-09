// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessagePublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMqttTraceMessagePublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMqttTraceMessagePublishResponse
	GetStatusCode() *int32
	SetBody(v *QueryMqttTraceMessagePublishResponseBody) *QueryMqttTraceMessagePublishResponse
	GetBody() *QueryMqttTraceMessagePublishResponseBody
}

type QueryMqttTraceMessagePublishResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMqttTraceMessagePublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMqttTraceMessagePublishResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessagePublishResponse) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMqttTraceMessagePublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMqttTraceMessagePublishResponse) GetBody() *QueryMqttTraceMessagePublishResponseBody {
	return s.Body
}

func (s *QueryMqttTraceMessagePublishResponse) SetHeaders(v map[string]*string) *QueryMqttTraceMessagePublishResponse {
	s.Headers = v
	return s
}

func (s *QueryMqttTraceMessagePublishResponse) SetStatusCode(v int32) *QueryMqttTraceMessagePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponse) SetBody(v *QueryMqttTraceMessagePublishResponseBody) *QueryMqttTraceMessagePublishResponse {
	s.Body = v
	return s
}

func (s *QueryMqttTraceMessagePublishResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
