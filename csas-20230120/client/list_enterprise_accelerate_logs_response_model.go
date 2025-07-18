// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAccelerateLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterpriseAccelerateLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterpriseAccelerateLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnterpriseAccelerateLogsResponseBody) *ListEnterpriseAccelerateLogsResponse
	GetBody() *ListEnterpriseAccelerateLogsResponseBody
}

type ListEnterpriseAccelerateLogsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseAccelerateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseAccelerateLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAccelerateLogsResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAccelerateLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterpriseAccelerateLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterpriseAccelerateLogsResponse) GetBody() *ListEnterpriseAccelerateLogsResponseBody {
	return s.Body
}

func (s *ListEnterpriseAccelerateLogsResponse) SetHeaders(v map[string]*string) *ListEnterpriseAccelerateLogsResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponse) SetStatusCode(v int32) *ListEnterpriseAccelerateLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponse) SetBody(v *ListEnterpriseAccelerateLogsResponseBody) *ListEnterpriseAccelerateLogsResponse {
	s.Body = v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponse) Validate() error {
	return dara.Validate(s)
}
