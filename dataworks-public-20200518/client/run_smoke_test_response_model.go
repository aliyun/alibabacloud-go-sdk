// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSmokeTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSmokeTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSmokeTestResponse
	GetStatusCode() *int32
	SetBody(v *RunSmokeTestResponseBody) *RunSmokeTestResponse
	GetBody() *RunSmokeTestResponseBody
}

type RunSmokeTestResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSmokeTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSmokeTestResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSmokeTestResponse) GoString() string {
	return s.String()
}

func (s *RunSmokeTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSmokeTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSmokeTestResponse) GetBody() *RunSmokeTestResponseBody {
	return s.Body
}

func (s *RunSmokeTestResponse) SetHeaders(v map[string]*string) *RunSmokeTestResponse {
	s.Headers = v
	return s
}

func (s *RunSmokeTestResponse) SetStatusCode(v int32) *RunSmokeTestResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSmokeTestResponse) SetBody(v *RunSmokeTestResponseBody) *RunSmokeTestResponse {
	s.Body = v
	return s
}

func (s *RunSmokeTestResponse) Validate() error {
	return dara.Validate(s)
}
