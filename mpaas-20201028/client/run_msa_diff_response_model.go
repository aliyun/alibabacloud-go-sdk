// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMsaDiffResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunMsaDiffResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunMsaDiffResponse
	GetStatusCode() *int32
	SetBody(v *RunMsaDiffResponseBody) *RunMsaDiffResponse
	GetBody() *RunMsaDiffResponseBody
}

type RunMsaDiffResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunMsaDiffResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunMsaDiffResponse) String() string {
	return dara.Prettify(s)
}

func (s RunMsaDiffResponse) GoString() string {
	return s.String()
}

func (s *RunMsaDiffResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunMsaDiffResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunMsaDiffResponse) GetBody() *RunMsaDiffResponseBody {
	return s.Body
}

func (s *RunMsaDiffResponse) SetHeaders(v map[string]*string) *RunMsaDiffResponse {
	s.Headers = v
	return s
}

func (s *RunMsaDiffResponse) SetStatusCode(v int32) *RunMsaDiffResponse {
	s.StatusCode = &v
	return s
}

func (s *RunMsaDiffResponse) SetBody(v *RunMsaDiffResponseBody) *RunMsaDiffResponse {
	s.Body = v
	return s
}

func (s *RunMsaDiffResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
