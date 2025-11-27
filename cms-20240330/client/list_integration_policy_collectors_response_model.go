// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyCollectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationPolicyCollectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationPolicyCollectorsResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationPolicyCollectorsResponseBody) *ListIntegrationPolicyCollectorsResponse
	GetBody() *ListIntegrationPolicyCollectorsResponseBody
}

type ListIntegrationPolicyCollectorsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationPolicyCollectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationPolicyCollectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCollectorsResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCollectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationPolicyCollectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationPolicyCollectorsResponse) GetBody() *ListIntegrationPolicyCollectorsResponseBody {
	return s.Body
}

func (s *ListIntegrationPolicyCollectorsResponse) SetHeaders(v map[string]*string) *ListIntegrationPolicyCollectorsResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponse) SetStatusCode(v int32) *ListIntegrationPolicyCollectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponse) SetBody(v *ListIntegrationPolicyCollectorsResponseBody) *ListIntegrationPolicyCollectorsResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationPolicyCollectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
