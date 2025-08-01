// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMonitorResponse
	GetStatusCode() *int32
	SetBody(v *QueryMonitorResponseBody) *QueryMonitorResponse
	GetBody() *QueryMonitorResponseBody
}

type QueryMonitorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMonitorResponse) GoString() string {
	return s.String()
}

func (s *QueryMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMonitorResponse) GetBody() *QueryMonitorResponseBody {
	return s.Body
}

func (s *QueryMonitorResponse) SetHeaders(v map[string]*string) *QueryMonitorResponse {
	s.Headers = v
	return s
}

func (s *QueryMonitorResponse) SetStatusCode(v int32) *QueryMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMonitorResponse) SetBody(v *QueryMonitorResponseBody) *QueryMonitorResponse {
	s.Body = v
	return s
}

func (s *QueryMonitorResponse) Validate() error {
	return dara.Validate(s)
}
