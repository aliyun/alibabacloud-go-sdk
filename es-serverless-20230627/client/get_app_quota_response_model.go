// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetAppQuotaResponseBody) *GetAppQuotaResponse
	GetBody() *GetAppQuotaResponseBody
}

type GetAppQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppQuotaResponse) GetBody() *GetAppQuotaResponseBody {
	return s.Body
}

func (s *GetAppQuotaResponse) SetHeaders(v map[string]*string) *GetAppQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetAppQuotaResponse) SetStatusCode(v int32) *GetAppQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppQuotaResponse) SetBody(v *GetAppQuotaResponseBody) *GetAppQuotaResponse {
	s.Body = v
	return s
}

func (s *GetAppQuotaResponse) Validate() error {
	return dara.Validate(s)
}
