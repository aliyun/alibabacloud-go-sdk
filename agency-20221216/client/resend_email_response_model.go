// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResendEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResendEmailResponse
	GetStatusCode() *int32
	SetBody(v *ResendEmailResponseBody) *ResendEmailResponse
	GetBody() *ResendEmailResponseBody
}

type ResendEmailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s ResendEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResendEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResendEmailResponse) GetBody() *ResendEmailResponseBody {
	return s.Body
}

func (s *ResendEmailResponse) SetHeaders(v map[string]*string) *ResendEmailResponse {
	s.Headers = v
	return s
}

func (s *ResendEmailResponse) SetStatusCode(v int32) *ResendEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendEmailResponse) SetBody(v *ResendEmailResponseBody) *ResendEmailResponse {
	s.Body = v
	return s
}

func (s *ResendEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
