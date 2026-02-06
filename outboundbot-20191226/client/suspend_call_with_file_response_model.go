// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendCallWithFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendCallWithFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendCallWithFileResponse
	GetStatusCode() *int32
	SetBody(v *SuspendCallWithFileResponseBody) *SuspendCallWithFileResponse
	GetBody() *SuspendCallWithFileResponseBody
}

type SuspendCallWithFileResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendCallWithFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendCallWithFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendCallWithFileResponse) GoString() string {
	return s.String()
}

func (s *SuspendCallWithFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendCallWithFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendCallWithFileResponse) GetBody() *SuspendCallWithFileResponseBody {
	return s.Body
}

func (s *SuspendCallWithFileResponse) SetHeaders(v map[string]*string) *SuspendCallWithFileResponse {
	s.Headers = v
	return s
}

func (s *SuspendCallWithFileResponse) SetStatusCode(v int32) *SuspendCallWithFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendCallWithFileResponse) SetBody(v *SuspendCallWithFileResponseBody) *SuspendCallWithFileResponse {
	s.Body = v
	return s
}

func (s *SuspendCallWithFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
