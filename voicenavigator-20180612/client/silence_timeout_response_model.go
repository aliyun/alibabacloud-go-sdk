// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSilenceTimeoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SilenceTimeoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SilenceTimeoutResponse
	GetStatusCode() *int32
	SetBody(v *SilenceTimeoutResponseBody) *SilenceTimeoutResponse
	GetBody() *SilenceTimeoutResponseBody
}

type SilenceTimeoutResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SilenceTimeoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SilenceTimeoutResponse) String() string {
	return dara.Prettify(s)
}

func (s SilenceTimeoutResponse) GoString() string {
	return s.String()
}

func (s *SilenceTimeoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SilenceTimeoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SilenceTimeoutResponse) GetBody() *SilenceTimeoutResponseBody {
	return s.Body
}

func (s *SilenceTimeoutResponse) SetHeaders(v map[string]*string) *SilenceTimeoutResponse {
	s.Headers = v
	return s
}

func (s *SilenceTimeoutResponse) SetStatusCode(v int32) *SilenceTimeoutResponse {
	s.StatusCode = &v
	return s
}

func (s *SilenceTimeoutResponse) SetBody(v *SilenceTimeoutResponseBody) *SilenceTimeoutResponse {
	s.Body = v
	return s
}

func (s *SilenceTimeoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
