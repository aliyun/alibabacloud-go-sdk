// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyStorageRequirementsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationPolicyStorageRequirementsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationPolicyStorageRequirementsResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationPolicyStorageRequirementsResponseBody) *ListIntegrationPolicyStorageRequirementsResponse
	GetBody() *ListIntegrationPolicyStorageRequirementsResponseBody
}

type ListIntegrationPolicyStorageRequirementsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationPolicyStorageRequirementsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationPolicyStorageRequirementsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyStorageRequirementsResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyStorageRequirementsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationPolicyStorageRequirementsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationPolicyStorageRequirementsResponse) GetBody() *ListIntegrationPolicyStorageRequirementsResponseBody {
	return s.Body
}

func (s *ListIntegrationPolicyStorageRequirementsResponse) SetHeaders(v map[string]*string) *ListIntegrationPolicyStorageRequirementsResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponse) SetStatusCode(v int32) *ListIntegrationPolicyStorageRequirementsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponse) SetBody(v *ListIntegrationPolicyStorageRequirementsResponseBody) *ListIntegrationPolicyStorageRequirementsResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
