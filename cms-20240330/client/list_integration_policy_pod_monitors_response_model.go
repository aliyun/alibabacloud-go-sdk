// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyPodMonitorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationPolicyPodMonitorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationPolicyPodMonitorsResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationPolicyPodMonitorsResponseBody) *ListIntegrationPolicyPodMonitorsResponse
	GetBody() *ListIntegrationPolicyPodMonitorsResponseBody
}

type ListIntegrationPolicyPodMonitorsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationPolicyPodMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationPolicyPodMonitorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyPodMonitorsResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyPodMonitorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationPolicyPodMonitorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationPolicyPodMonitorsResponse) GetBody() *ListIntegrationPolicyPodMonitorsResponseBody {
	return s.Body
}

func (s *ListIntegrationPolicyPodMonitorsResponse) SetHeaders(v map[string]*string) *ListIntegrationPolicyPodMonitorsResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponse) SetStatusCode(v int32) *ListIntegrationPolicyPodMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponse) SetBody(v *ListIntegrationPolicyPodMonitorsResponseBody) *ListIntegrationPolicyPodMonitorsResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationPolicyPodMonitorsResponse) Validate() error {
	return dara.Validate(s)
}
