// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAcceleratePoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterpriseAcceleratePoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterpriseAcceleratePoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListEnterpriseAcceleratePoliciesResponseBody) *ListEnterpriseAcceleratePoliciesResponse
	GetBody() *ListEnterpriseAcceleratePoliciesResponseBody
}

type ListEnterpriseAcceleratePoliciesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseAcceleratePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseAcceleratePoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAcceleratePoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAcceleratePoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterpriseAcceleratePoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterpriseAcceleratePoliciesResponse) GetBody() *ListEnterpriseAcceleratePoliciesResponseBody {
	return s.Body
}

func (s *ListEnterpriseAcceleratePoliciesResponse) SetHeaders(v map[string]*string) *ListEnterpriseAcceleratePoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponse) SetStatusCode(v int32) *ListEnterpriseAcceleratePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponse) SetBody(v *ListEnterpriseAcceleratePoliciesResponseBody) *ListEnterpriseAcceleratePoliciesResponse {
	s.Body = v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
