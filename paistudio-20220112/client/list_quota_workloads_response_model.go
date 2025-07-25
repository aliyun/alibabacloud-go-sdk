// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaWorkloadsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaWorkloadsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaWorkloadsResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaWorkloadsResponseBody) *ListQuotaWorkloadsResponse
	GetBody() *ListQuotaWorkloadsResponseBody
}

type ListQuotaWorkloadsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaWorkloadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaWorkloadsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaWorkloadsResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaWorkloadsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaWorkloadsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaWorkloadsResponse) GetBody() *ListQuotaWorkloadsResponseBody {
	return s.Body
}

func (s *ListQuotaWorkloadsResponse) SetHeaders(v map[string]*string) *ListQuotaWorkloadsResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaWorkloadsResponse) SetStatusCode(v int32) *ListQuotaWorkloadsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaWorkloadsResponse) SetBody(v *ListQuotaWorkloadsResponseBody) *ListQuotaWorkloadsResponse {
	s.Body = v
	return s
}

func (s *ListQuotaWorkloadsResponse) Validate() error {
	return dara.Validate(s)
}
