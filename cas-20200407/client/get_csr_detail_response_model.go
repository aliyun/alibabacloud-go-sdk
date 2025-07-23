// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCsrDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCsrDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCsrDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCsrDetailResponseBody) *GetCsrDetailResponse
	GetBody() *GetCsrDetailResponseBody
}

type GetCsrDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCsrDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCsrDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCsrDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCsrDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCsrDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCsrDetailResponse) GetBody() *GetCsrDetailResponseBody {
	return s.Body
}

func (s *GetCsrDetailResponse) SetHeaders(v map[string]*string) *GetCsrDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCsrDetailResponse) SetStatusCode(v int32) *GetCsrDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCsrDetailResponse) SetBody(v *GetCsrDetailResponseBody) *GetCsrDetailResponse {
	s.Body = v
	return s
}

func (s *GetCsrDetailResponse) Validate() error {
	return dara.Validate(s)
}
