// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartCollectorResponse
	GetStatusCode() *int32
	SetBody(v *RestartCollectorResponseBody) *RestartCollectorResponse
	GetBody() *RestartCollectorResponseBody
}

type RestartCollectorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartCollectorResponse) GoString() string {
	return s.String()
}

func (s *RestartCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartCollectorResponse) GetBody() *RestartCollectorResponseBody {
	return s.Body
}

func (s *RestartCollectorResponse) SetHeaders(v map[string]*string) *RestartCollectorResponse {
	s.Headers = v
	return s
}

func (s *RestartCollectorResponse) SetStatusCode(v int32) *RestartCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartCollectorResponse) SetBody(v *RestartCollectorResponseBody) *RestartCollectorResponse {
	s.Body = v
	return s
}

func (s *RestartCollectorResponse) Validate() error {
	return dara.Validate(s)
}
