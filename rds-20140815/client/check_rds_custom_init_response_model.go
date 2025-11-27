// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRdsCustomInitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckRdsCustomInitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckRdsCustomInitResponse
	GetStatusCode() *int32
	SetBody(v *CheckRdsCustomInitResponseBody) *CheckRdsCustomInitResponse
	GetBody() *CheckRdsCustomInitResponseBody
}

type CheckRdsCustomInitResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckRdsCustomInitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckRdsCustomInitResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckRdsCustomInitResponse) GoString() string {
	return s.String()
}

func (s *CheckRdsCustomInitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckRdsCustomInitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckRdsCustomInitResponse) GetBody() *CheckRdsCustomInitResponseBody {
	return s.Body
}

func (s *CheckRdsCustomInitResponse) SetHeaders(v map[string]*string) *CheckRdsCustomInitResponse {
	s.Headers = v
	return s
}

func (s *CheckRdsCustomInitResponse) SetStatusCode(v int32) *CheckRdsCustomInitResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRdsCustomInitResponse) SetBody(v *CheckRdsCustomInitResponseBody) *CheckRdsCustomInitResponse {
	s.Body = v
	return s
}

func (s *CheckRdsCustomInitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
