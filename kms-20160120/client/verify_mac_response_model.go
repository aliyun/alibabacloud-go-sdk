// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMacResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyMacResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyMacResponse
	GetStatusCode() *int32
	SetBody(v *VerifyMacResponseBody) *VerifyMacResponse
	GetBody() *VerifyMacResponseBody
}

type VerifyMacResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyMacResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyMacResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyMacResponse) GoString() string {
	return s.String()
}

func (s *VerifyMacResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyMacResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyMacResponse) GetBody() *VerifyMacResponseBody {
	return s.Body
}

func (s *VerifyMacResponse) SetHeaders(v map[string]*string) *VerifyMacResponse {
	s.Headers = v
	return s
}

func (s *VerifyMacResponse) SetStatusCode(v int32) *VerifyMacResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyMacResponse) SetBody(v *VerifyMacResponseBody) *VerifyMacResponse {
	s.Body = v
	return s
}

func (s *VerifyMacResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
