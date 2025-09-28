// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMobileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyMobileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyMobileResponse
	GetStatusCode() *int32
	SetBody(v *VerifyMobileResponseBody) *VerifyMobileResponse
	GetBody() *VerifyMobileResponseBody
}

type VerifyMobileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyMobileResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyMobileResponse) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyMobileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyMobileResponse) GetBody() *VerifyMobileResponseBody {
	return s.Body
}

func (s *VerifyMobileResponse) SetHeaders(v map[string]*string) *VerifyMobileResponse {
	s.Headers = v
	return s
}

func (s *VerifyMobileResponse) SetStatusCode(v int32) *VerifyMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyMobileResponse) SetBody(v *VerifyMobileResponseBody) *VerifyMobileResponse {
	s.Body = v
	return s
}

func (s *VerifyMobileResponse) Validate() error {
	return dara.Validate(s)
}
