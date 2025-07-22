// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestStatResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncErrorRequestStatResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncErrorRequestStatResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAsyncErrorRequestStatResultResponseBody) *GetAsyncErrorRequestStatResultResponse
	GetBody() *GetAsyncErrorRequestStatResultResponseBody
}

type GetAsyncErrorRequestStatResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncErrorRequestStatResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncErrorRequestStatResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestStatResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestStatResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncErrorRequestStatResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncErrorRequestStatResultResponse) GetBody() *GetAsyncErrorRequestStatResultResponseBody {
	return s.Body
}

func (s *GetAsyncErrorRequestStatResultResponse) SetHeaders(v map[string]*string) *GetAsyncErrorRequestStatResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponse) SetStatusCode(v int32) *GetAsyncErrorRequestStatResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponse) SetBody(v *GetAsyncErrorRequestStatResultResponseBody) *GetAsyncErrorRequestStatResultResponse {
	s.Body = v
	return s
}

func (s *GetAsyncErrorRequestStatResultResponse) Validate() error {
	return dara.Validate(s)
}
