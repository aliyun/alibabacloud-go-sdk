// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVerifyCodeResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVerifyCodeResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVerifyCodeResultResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVerifyCodeResultResponseBody) *SubmitVerifyCodeResultResponse
	GetBody() *SubmitVerifyCodeResultResponseBody
}

type SubmitVerifyCodeResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVerifyCodeResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVerifyCodeResultResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVerifyCodeResultResponse) GoString() string {
	return s.String()
}

func (s *SubmitVerifyCodeResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVerifyCodeResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVerifyCodeResultResponse) GetBody() *SubmitVerifyCodeResultResponseBody {
	return s.Body
}

func (s *SubmitVerifyCodeResultResponse) SetHeaders(v map[string]*string) *SubmitVerifyCodeResultResponse {
	s.Headers = v
	return s
}

func (s *SubmitVerifyCodeResultResponse) SetStatusCode(v int32) *SubmitVerifyCodeResultResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVerifyCodeResultResponse) SetBody(v *SubmitVerifyCodeResultResponseBody) *SubmitVerifyCodeResultResponse {
	s.Body = v
	return s
}

func (s *SubmitVerifyCodeResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
