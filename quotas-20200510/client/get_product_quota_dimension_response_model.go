// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductQuotaDimensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductQuotaDimensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductQuotaDimensionResponse
	GetStatusCode() *int32
	SetBody(v *GetProductQuotaDimensionResponseBody) *GetProductQuotaDimensionResponse
	GetBody() *GetProductQuotaDimensionResponseBody
}

type GetProductQuotaDimensionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductQuotaDimensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductQuotaDimensionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaDimensionResponse) GoString() string {
	return s.String()
}

func (s *GetProductQuotaDimensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductQuotaDimensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductQuotaDimensionResponse) GetBody() *GetProductQuotaDimensionResponseBody {
	return s.Body
}

func (s *GetProductQuotaDimensionResponse) SetHeaders(v map[string]*string) *GetProductQuotaDimensionResponse {
	s.Headers = v
	return s
}

func (s *GetProductQuotaDimensionResponse) SetStatusCode(v int32) *GetProductQuotaDimensionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductQuotaDimensionResponse) SetBody(v *GetProductQuotaDimensionResponseBody) *GetProductQuotaDimensionResponse {
	s.Body = v
	return s
}

func (s *GetProductQuotaDimensionResponse) Validate() error {
	return dara.Validate(s)
}
