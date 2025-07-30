// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyFileResponse
	GetStatusCode() *int32
	SetBody(v *VerifyFileResponseBody) *VerifyFileResponse
	GetBody() *VerifyFileResponseBody
}

type VerifyFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyFileResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyFileResponse) GoString() string {
	return s.String()
}

func (s *VerifyFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyFileResponse) GetBody() *VerifyFileResponseBody {
	return s.Body
}

func (s *VerifyFileResponse) SetHeaders(v map[string]*string) *VerifyFileResponse {
	s.Headers = v
	return s
}

func (s *VerifyFileResponse) SetStatusCode(v int32) *VerifyFileResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyFileResponse) SetBody(v *VerifyFileResponseBody) *VerifyFileResponse {
	s.Body = v
	return s
}

func (s *VerifyFileResponse) Validate() error {
	return dara.Validate(s)
}
