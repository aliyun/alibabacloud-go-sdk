// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBaselineSecurityCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartBaselineSecurityCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartBaselineSecurityCheckResponse
	GetStatusCode() *int32
	SetBody(v *StartBaselineSecurityCheckResponseBody) *StartBaselineSecurityCheckResponse
	GetBody() *StartBaselineSecurityCheckResponseBody
}

type StartBaselineSecurityCheckResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartBaselineSecurityCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartBaselineSecurityCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s StartBaselineSecurityCheckResponse) GoString() string {
	return s.String()
}

func (s *StartBaselineSecurityCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartBaselineSecurityCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartBaselineSecurityCheckResponse) GetBody() *StartBaselineSecurityCheckResponseBody {
	return s.Body
}

func (s *StartBaselineSecurityCheckResponse) SetHeaders(v map[string]*string) *StartBaselineSecurityCheckResponse {
	s.Headers = v
	return s
}

func (s *StartBaselineSecurityCheckResponse) SetStatusCode(v int32) *StartBaselineSecurityCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBaselineSecurityCheckResponse) SetBody(v *StartBaselineSecurityCheckResponseBody) *StartBaselineSecurityCheckResponse {
	s.Body = v
	return s
}

func (s *StartBaselineSecurityCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
