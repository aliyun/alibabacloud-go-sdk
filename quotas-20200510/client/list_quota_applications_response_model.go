// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaApplicationsResponseBody) *ListQuotaApplicationsResponse
	GetBody() *ListQuotaApplicationsResponseBody
}

type ListQuotaApplicationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaApplicationsResponse) GetBody() *ListQuotaApplicationsResponseBody {
	return s.Body
}

func (s *ListQuotaApplicationsResponse) SetHeaders(v map[string]*string) *ListQuotaApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaApplicationsResponse) SetStatusCode(v int32) *ListQuotaApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaApplicationsResponse) SetBody(v *ListQuotaApplicationsResponseBody) *ListQuotaApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListQuotaApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
