// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDemoInstanceExistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDemoInstanceExistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDemoInstanceExistsResponse
	GetStatusCode() *int32
	SetBody(v *CheckDemoInstanceExistsResponseBody) *CheckDemoInstanceExistsResponse
	GetBody() *CheckDemoInstanceExistsResponseBody
}

type CheckDemoInstanceExistsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDemoInstanceExistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDemoInstanceExistsResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDemoInstanceExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckDemoInstanceExistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDemoInstanceExistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDemoInstanceExistsResponse) GetBody() *CheckDemoInstanceExistsResponseBody {
	return s.Body
}

func (s *CheckDemoInstanceExistsResponse) SetHeaders(v map[string]*string) *CheckDemoInstanceExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckDemoInstanceExistsResponse) SetStatusCode(v int32) *CheckDemoInstanceExistsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDemoInstanceExistsResponse) SetBody(v *CheckDemoInstanceExistsResponseBody) *CheckDemoInstanceExistsResponse {
	s.Body = v
	return s
}

func (s *CheckDemoInstanceExistsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
