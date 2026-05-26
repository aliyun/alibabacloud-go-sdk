// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicySetAttachedGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicySetAttachedGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicySetAttachedGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicySetAttachedGatewaysResponseBody) *ListPolicySetAttachedGatewaysResponse
	GetBody() *ListPolicySetAttachedGatewaysResponseBody
}

type ListPolicySetAttachedGatewaysResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicySetAttachedGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicySetAttachedGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicySetAttachedGatewaysResponse) GoString() string {
	return s.String()
}

func (s *ListPolicySetAttachedGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicySetAttachedGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicySetAttachedGatewaysResponse) GetBody() *ListPolicySetAttachedGatewaysResponseBody {
	return s.Body
}

func (s *ListPolicySetAttachedGatewaysResponse) SetHeaders(v map[string]*string) *ListPolicySetAttachedGatewaysResponse {
	s.Headers = v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponse) SetStatusCode(v int32) *ListPolicySetAttachedGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponse) SetBody(v *ListPolicySetAttachedGatewaysResponseBody) *ListPolicySetAttachedGatewaysResponse {
	s.Body = v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
