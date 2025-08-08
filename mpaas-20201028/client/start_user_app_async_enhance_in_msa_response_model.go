// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartUserAppAsyncEnhanceInMsaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartUserAppAsyncEnhanceInMsaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartUserAppAsyncEnhanceInMsaResponse
	GetStatusCode() *int32
	SetBody(v *StartUserAppAsyncEnhanceInMsaResponseBody) *StartUserAppAsyncEnhanceInMsaResponse
	GetBody() *StartUserAppAsyncEnhanceInMsaResponseBody
}

type StartUserAppAsyncEnhanceInMsaResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartUserAppAsyncEnhanceInMsaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartUserAppAsyncEnhanceInMsaResponse) String() string {
	return dara.Prettify(s)
}

func (s StartUserAppAsyncEnhanceInMsaResponse) GoString() string {
	return s.String()
}

func (s *StartUserAppAsyncEnhanceInMsaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartUserAppAsyncEnhanceInMsaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartUserAppAsyncEnhanceInMsaResponse) GetBody() *StartUserAppAsyncEnhanceInMsaResponseBody {
	return s.Body
}

func (s *StartUserAppAsyncEnhanceInMsaResponse) SetHeaders(v map[string]*string) *StartUserAppAsyncEnhanceInMsaResponse {
	s.Headers = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponse) SetStatusCode(v int32) *StartUserAppAsyncEnhanceInMsaResponse {
	s.StatusCode = &v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponse) SetBody(v *StartUserAppAsyncEnhanceInMsaResponseBody) *StartUserAppAsyncEnhanceInMsaResponse {
	s.Body = v
	return s
}

func (s *StartUserAppAsyncEnhanceInMsaResponse) Validate() error {
	return dara.Validate(s)
}
