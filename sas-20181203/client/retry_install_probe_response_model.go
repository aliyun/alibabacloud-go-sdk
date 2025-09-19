// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryInstallProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryInstallProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryInstallProbeResponse
	GetStatusCode() *int32
	SetBody(v *RetryInstallProbeResponseBody) *RetryInstallProbeResponse
	GetBody() *RetryInstallProbeResponseBody
}

type RetryInstallProbeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryInstallProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryInstallProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryInstallProbeResponse) GoString() string {
	return s.String()
}

func (s *RetryInstallProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryInstallProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryInstallProbeResponse) GetBody() *RetryInstallProbeResponseBody {
	return s.Body
}

func (s *RetryInstallProbeResponse) SetHeaders(v map[string]*string) *RetryInstallProbeResponse {
	s.Headers = v
	return s
}

func (s *RetryInstallProbeResponse) SetStatusCode(v int32) *RetryInstallProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryInstallProbeResponse) SetBody(v *RetryInstallProbeResponseBody) *RetryInstallProbeResponse {
	s.Body = v
	return s
}

func (s *RetryInstallProbeResponse) Validate() error {
	return dara.Validate(s)
}
