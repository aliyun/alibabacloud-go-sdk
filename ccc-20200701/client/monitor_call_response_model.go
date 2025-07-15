// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MonitorCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MonitorCallResponse
	GetStatusCode() *int32
	SetBody(v *MonitorCallResponseBody) *MonitorCallResponse
	GetBody() *MonitorCallResponseBody
}

type MonitorCallResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MonitorCallResponse) String() string {
	return dara.Prettify(s)
}

func (s MonitorCallResponse) GoString() string {
	return s.String()
}

func (s *MonitorCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MonitorCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MonitorCallResponse) GetBody() *MonitorCallResponseBody {
	return s.Body
}

func (s *MonitorCallResponse) SetHeaders(v map[string]*string) *MonitorCallResponse {
	s.Headers = v
	return s
}

func (s *MonitorCallResponse) SetStatusCode(v int32) *MonitorCallResponse {
	s.StatusCode = &v
	return s
}

func (s *MonitorCallResponse) SetBody(v *MonitorCallResponseBody) *MonitorCallResponse {
	s.Body = v
	return s
}

func (s *MonitorCallResponse) Validate() error {
	return dara.Validate(s)
}
