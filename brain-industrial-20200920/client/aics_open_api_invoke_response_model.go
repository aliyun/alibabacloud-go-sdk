// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAicsOpenApiInvokeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AicsOpenApiInvokeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AicsOpenApiInvokeResponse
	GetStatusCode() *int32
	SetBody(v *AicsOpenApiInvokeResponseBody) *AicsOpenApiInvokeResponse
	GetBody() *AicsOpenApiInvokeResponseBody
}

type AicsOpenApiInvokeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AicsOpenApiInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AicsOpenApiInvokeResponse) String() string {
	return dara.Prettify(s)
}

func (s AicsOpenApiInvokeResponse) GoString() string {
	return s.String()
}

func (s *AicsOpenApiInvokeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AicsOpenApiInvokeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AicsOpenApiInvokeResponse) GetBody() *AicsOpenApiInvokeResponseBody {
	return s.Body
}

func (s *AicsOpenApiInvokeResponse) SetHeaders(v map[string]*string) *AicsOpenApiInvokeResponse {
	s.Headers = v
	return s
}

func (s *AicsOpenApiInvokeResponse) SetStatusCode(v int32) *AicsOpenApiInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *AicsOpenApiInvokeResponse) SetBody(v *AicsOpenApiInvokeResponseBody) *AicsOpenApiInvokeResponse {
	s.Body = v
	return s
}

func (s *AicsOpenApiInvokeResponse) Validate() error {
	return dara.Validate(s)
}
