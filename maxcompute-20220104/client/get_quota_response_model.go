// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaResponseBody) *GetQuotaResponse
	GetBody() *GetQuotaResponseBody
}

type GetQuotaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaResponse) GetBody() *GetQuotaResponseBody {
	return s.Body
}

func (s *GetQuotaResponse) SetHeaders(v map[string]*string) *GetQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaResponse) SetStatusCode(v int32) *GetQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaResponse) SetBody(v *GetQuotaResponseBody) *GetQuotaResponse {
	s.Body = v
	return s
}

func (s *GetQuotaResponse) Validate() error {
	return dara.Validate(s)
}
