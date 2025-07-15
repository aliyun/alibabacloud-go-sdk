// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVerifyLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckVerifyLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckVerifyLogResponse
	GetStatusCode() *int32
	SetBody(v *CheckVerifyLogResponseBody) *CheckVerifyLogResponse
	GetBody() *CheckVerifyLogResponseBody
}

type CheckVerifyLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckVerifyLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckVerifyLogResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckVerifyLogResponse) GoString() string {
	return s.String()
}

func (s *CheckVerifyLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckVerifyLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckVerifyLogResponse) GetBody() *CheckVerifyLogResponseBody {
	return s.Body
}

func (s *CheckVerifyLogResponse) SetHeaders(v map[string]*string) *CheckVerifyLogResponse {
	s.Headers = v
	return s
}

func (s *CheckVerifyLogResponse) SetStatusCode(v int32) *CheckVerifyLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckVerifyLogResponse) SetBody(v *CheckVerifyLogResponseBody) *CheckVerifyLogResponse {
	s.Body = v
	return s
}

func (s *CheckVerifyLogResponse) Validate() error {
	return dara.Validate(s)
}
