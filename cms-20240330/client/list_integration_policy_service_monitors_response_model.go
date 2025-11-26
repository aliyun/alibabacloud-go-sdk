// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyServiceMonitorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationPolicyServiceMonitorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationPolicyServiceMonitorsResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationPolicyServiceMonitorsResponseBody) *ListIntegrationPolicyServiceMonitorsResponse
	GetBody() *ListIntegrationPolicyServiceMonitorsResponseBody
}

type ListIntegrationPolicyServiceMonitorsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationPolicyServiceMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationPolicyServiceMonitorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyServiceMonitorsResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyServiceMonitorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationPolicyServiceMonitorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationPolicyServiceMonitorsResponse) GetBody() *ListIntegrationPolicyServiceMonitorsResponseBody {
	return s.Body
}

func (s *ListIntegrationPolicyServiceMonitorsResponse) SetHeaders(v map[string]*string) *ListIntegrationPolicyServiceMonitorsResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponse) SetStatusCode(v int32) *ListIntegrationPolicyServiceMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponse) SetBody(v *ListIntegrationPolicyServiceMonitorsResponseBody) *ListIntegrationPolicyServiceMonitorsResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationPolicyServiceMonitorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
