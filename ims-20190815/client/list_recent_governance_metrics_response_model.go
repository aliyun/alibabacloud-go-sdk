// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentGovernanceMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecentGovernanceMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecentGovernanceMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecentGovernanceMetricsResponseBody) *ListRecentGovernanceMetricsResponse
	GetBody() *ListRecentGovernanceMetricsResponseBody
}

type ListRecentGovernanceMetricsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentGovernanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecentGovernanceMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecentGovernanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentGovernanceMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecentGovernanceMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecentGovernanceMetricsResponse) GetBody() *ListRecentGovernanceMetricsResponseBody {
	return s.Body
}

func (s *ListRecentGovernanceMetricsResponse) SetHeaders(v map[string]*string) *ListRecentGovernanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentGovernanceMetricsResponse) SetStatusCode(v int32) *ListRecentGovernanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponse) SetBody(v *ListRecentGovernanceMetricsResponseBody) *ListRecentGovernanceMetricsResponse {
	s.Body = v
	return s
}

func (s *ListRecentGovernanceMetricsResponse) Validate() error {
	return dara.Validate(s)
}
