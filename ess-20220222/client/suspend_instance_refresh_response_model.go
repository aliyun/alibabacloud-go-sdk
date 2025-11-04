// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendInstanceRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendInstanceRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendInstanceRefreshResponse
	GetStatusCode() *int32
	SetBody(v *SuspendInstanceRefreshResponseBody) *SuspendInstanceRefreshResponse
	GetBody() *SuspendInstanceRefreshResponseBody
}

type SuspendInstanceRefreshResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendInstanceRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendInstanceRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendInstanceRefreshResponse) GoString() string {
	return s.String()
}

func (s *SuspendInstanceRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendInstanceRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendInstanceRefreshResponse) GetBody() *SuspendInstanceRefreshResponseBody {
	return s.Body
}

func (s *SuspendInstanceRefreshResponse) SetHeaders(v map[string]*string) *SuspendInstanceRefreshResponse {
	s.Headers = v
	return s
}

func (s *SuspendInstanceRefreshResponse) SetStatusCode(v int32) *SuspendInstanceRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendInstanceRefreshResponse) SetBody(v *SuspendInstanceRefreshResponseBody) *SuspendInstanceRefreshResponse {
	s.Body = v
	return s
}

func (s *SuspendInstanceRefreshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
