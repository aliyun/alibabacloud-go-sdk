// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestStatByCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncErrorRequestStatByCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncErrorRequestStatByCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetAsyncErrorRequestStatByCodeResponseBody) *GetAsyncErrorRequestStatByCodeResponse
	GetBody() *GetAsyncErrorRequestStatByCodeResponseBody
}

type GetAsyncErrorRequestStatByCodeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncErrorRequestStatByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncErrorRequestStatByCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatByCodeResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatByCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncErrorRequestStatByCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncErrorRequestStatByCodeResponse) GetBody() *GetAsyncErrorRequestStatByCodeResponseBody {
	return s.Body
}

func (s *GetAsyncErrorRequestStatByCodeResponse) SetHeaders(v map[string]*string) *GetAsyncErrorRequestStatByCodeResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponse) SetStatusCode(v int32) *GetAsyncErrorRequestStatByCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponse) SetBody(v *GetAsyncErrorRequestStatByCodeResponseBody) *GetAsyncErrorRequestStatByCodeResponse {
	s.Body = v
	return s
}

func (s *GetAsyncErrorRequestStatByCodeResponse) Validate() error {
	return dara.Validate(s)
}
