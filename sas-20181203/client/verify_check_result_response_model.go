// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *VerifyCheckResultResponseBody) *VerifyCheckResultResponse
	GetBody() *VerifyCheckResultResponseBody
}

type VerifyCheckResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckResultResponse) GoString() string {
	return s.String()
}

func (s *VerifyCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyCheckResultResponse) GetBody() *VerifyCheckResultResponseBody {
	return s.Body
}

func (s *VerifyCheckResultResponse) SetHeaders(v map[string]*string) *VerifyCheckResultResponse {
	s.Headers = v
	return s
}

func (s *VerifyCheckResultResponse) SetStatusCode(v int32) *VerifyCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCheckResultResponse) SetBody(v *VerifyCheckResultResponseBody) *VerifyCheckResultResponse {
	s.Body = v
	return s
}

func (s *VerifyCheckResultResponse) Validate() error {
	return dara.Validate(s)
}
