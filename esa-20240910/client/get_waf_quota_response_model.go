// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWafQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWafQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetWafQuotaResponseBody) *GetWafQuotaResponse
	GetBody() *GetWafQuotaResponseBody
}

type GetWafQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWafQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetWafQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWafQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWafQuotaResponse) GetBody() *GetWafQuotaResponseBody {
	return s.Body
}

func (s *GetWafQuotaResponse) SetHeaders(v map[string]*string) *GetWafQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetWafQuotaResponse) SetStatusCode(v int32) *GetWafQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafQuotaResponse) SetBody(v *GetWafQuotaResponseBody) *GetWafQuotaResponse {
	s.Body = v
	return s
}

func (s *GetWafQuotaResponse) Validate() error {
	return dara.Validate(s)
}
