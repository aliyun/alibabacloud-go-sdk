// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyDashboardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationPolicyDashboardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationPolicyDashboardsResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationPolicyDashboardsResponseBody) *ListIntegrationPolicyDashboardsResponse
	GetBody() *ListIntegrationPolicyDashboardsResponseBody
}

type ListIntegrationPolicyDashboardsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationPolicyDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationPolicyDashboardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyDashboardsResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyDashboardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationPolicyDashboardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationPolicyDashboardsResponse) GetBody() *ListIntegrationPolicyDashboardsResponseBody {
	return s.Body
}

func (s *ListIntegrationPolicyDashboardsResponse) SetHeaders(v map[string]*string) *ListIntegrationPolicyDashboardsResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponse) SetStatusCode(v int32) *ListIntegrationPolicyDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponse) SetBody(v *ListIntegrationPolicyDashboardsResponseBody) *ListIntegrationPolicyDashboardsResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponse) Validate() error {
	return dara.Validate(s)
}
