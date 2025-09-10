// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetProductQuotaResponseBody) *GetProductQuotaResponse
	GetBody() *GetProductQuotaResponseBody
}

type GetProductQuotaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductQuotaResponse) GetBody() *GetProductQuotaResponseBody {
	return s.Body
}

func (s *GetProductQuotaResponse) SetHeaders(v map[string]*string) *GetProductQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetProductQuotaResponse) SetStatusCode(v int32) *GetProductQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductQuotaResponse) SetBody(v *GetProductQuotaResponseBody) *GetProductQuotaResponse {
	s.Body = v
	return s
}

func (s *GetProductQuotaResponse) Validate() error {
	return dara.Validate(s)
}
