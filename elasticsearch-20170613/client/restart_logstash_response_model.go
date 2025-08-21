// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLogstashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartLogstashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartLogstashResponse
	GetStatusCode() *int32
	SetBody(v *RestartLogstashResponseBody) *RestartLogstashResponse
	GetBody() *RestartLogstashResponseBody
}

type RestartLogstashResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartLogstashResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartLogstashResponse) GoString() string {
	return s.String()
}

func (s *RestartLogstashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartLogstashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartLogstashResponse) GetBody() *RestartLogstashResponseBody {
	return s.Body
}

func (s *RestartLogstashResponse) SetHeaders(v map[string]*string) *RestartLogstashResponse {
	s.Headers = v
	return s
}

func (s *RestartLogstashResponse) SetStatusCode(v int32) *RestartLogstashResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartLogstashResponse) SetBody(v *RestartLogstashResponseBody) *RestartLogstashResponse {
	s.Body = v
	return s
}

func (s *RestartLogstashResponse) Validate() error {
	return dara.Validate(s)
}
