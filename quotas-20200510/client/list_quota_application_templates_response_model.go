// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaApplicationTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaApplicationTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaApplicationTemplatesResponseBody) *ListQuotaApplicationTemplatesResponse
	GetBody() *ListQuotaApplicationTemplatesResponseBody
}

type ListQuotaApplicationTemplatesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaApplicationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaApplicationTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaApplicationTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaApplicationTemplatesResponse) GetBody() *ListQuotaApplicationTemplatesResponseBody {
	return s.Body
}

func (s *ListQuotaApplicationTemplatesResponse) SetHeaders(v map[string]*string) *ListQuotaApplicationTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponse) SetStatusCode(v int32) *ListQuotaApplicationTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaApplicationTemplatesResponse) SetBody(v *ListQuotaApplicationTemplatesResponseBody) *ListQuotaApplicationTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListQuotaApplicationTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
