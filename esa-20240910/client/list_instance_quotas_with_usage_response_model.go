// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceQuotasWithUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceQuotasWithUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceQuotasWithUsageResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceQuotasWithUsageResponseBody) *ListInstanceQuotasWithUsageResponse
	GetBody() *ListInstanceQuotasWithUsageResponseBody
}

type ListInstanceQuotasWithUsageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceQuotasWithUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceQuotasWithUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasWithUsageResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasWithUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceQuotasWithUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceQuotasWithUsageResponse) GetBody() *ListInstanceQuotasWithUsageResponseBody {
	return s.Body
}

func (s *ListInstanceQuotasWithUsageResponse) SetHeaders(v map[string]*string) *ListInstanceQuotasWithUsageResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceQuotasWithUsageResponse) SetStatusCode(v int32) *ListInstanceQuotasWithUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceQuotasWithUsageResponse) SetBody(v *ListInstanceQuotasWithUsageResponseBody) *ListInstanceQuotasWithUsageResponse {
	s.Body = v
	return s
}

func (s *ListInstanceQuotasWithUsageResponse) Validate() error {
	return dara.Validate(s)
}
