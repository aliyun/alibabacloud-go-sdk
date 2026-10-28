// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendVerifyActivateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResendVerifyActivateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResendVerifyActivateResponse
	GetStatusCode() *int32
	SetBody(v *ResendVerifyActivateResponseBody) *ResendVerifyActivateResponse
	GetBody() *ResendVerifyActivateResponseBody
}

type ResendVerifyActivateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendVerifyActivateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendVerifyActivateResponse) String() string {
	return dara.Prettify(s)
}

func (s ResendVerifyActivateResponse) GoString() string {
	return s.String()
}

func (s *ResendVerifyActivateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResendVerifyActivateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResendVerifyActivateResponse) GetBody() *ResendVerifyActivateResponseBody {
	return s.Body
}

func (s *ResendVerifyActivateResponse) SetHeaders(v map[string]*string) *ResendVerifyActivateResponse {
	s.Headers = v
	return s
}

func (s *ResendVerifyActivateResponse) SetStatusCode(v int32) *ResendVerifyActivateResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendVerifyActivateResponse) SetBody(v *ResendVerifyActivateResponseBody) *ResendVerifyActivateResponse {
	s.Body = v
	return s
}

func (s *ResendVerifyActivateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
