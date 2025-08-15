// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyConsumeMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyConsumeMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyConsumeMessageResponse
	GetStatusCode() *int32
	SetBody(v *VerifyConsumeMessageResponseBody) *VerifyConsumeMessageResponse
	GetBody() *VerifyConsumeMessageResponseBody
}

type VerifyConsumeMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyConsumeMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyConsumeMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyConsumeMessageResponse) GoString() string {
	return s.String()
}

func (s *VerifyConsumeMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyConsumeMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyConsumeMessageResponse) GetBody() *VerifyConsumeMessageResponseBody {
	return s.Body
}

func (s *VerifyConsumeMessageResponse) SetHeaders(v map[string]*string) *VerifyConsumeMessageResponse {
	s.Headers = v
	return s
}

func (s *VerifyConsumeMessageResponse) SetStatusCode(v int32) *VerifyConsumeMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyConsumeMessageResponse) SetBody(v *VerifyConsumeMessageResponseBody) *VerifyConsumeMessageResponse {
	s.Body = v
	return s
}

func (s *VerifyConsumeMessageResponse) Validate() error {
	return dara.Validate(s)
}
