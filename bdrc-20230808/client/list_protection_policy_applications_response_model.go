// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectionPolicyApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProtectionPolicyApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProtectionPolicyApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListProtectionPolicyApplicationsResponseBody) *ListProtectionPolicyApplicationsResponse
	GetBody() *ListProtectionPolicyApplicationsResponseBody
}

type ListProtectionPolicyApplicationsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProtectionPolicyApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProtectionPolicyApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPolicyApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListProtectionPolicyApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProtectionPolicyApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProtectionPolicyApplicationsResponse) GetBody() *ListProtectionPolicyApplicationsResponseBody {
	return s.Body
}

func (s *ListProtectionPolicyApplicationsResponse) SetHeaders(v map[string]*string) *ListProtectionPolicyApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListProtectionPolicyApplicationsResponse) SetStatusCode(v int32) *ListProtectionPolicyApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponse) SetBody(v *ListProtectionPolicyApplicationsResponseBody) *ListProtectionPolicyApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListProtectionPolicyApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
