// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkloadIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkloadIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkloadIdentityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkloadIdentityResponseBody) *UpdateWorkloadIdentityResponse
	GetBody() *UpdateWorkloadIdentityResponseBody
}

type UpdateWorkloadIdentityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkloadIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkloadIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkloadIdentityResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkloadIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkloadIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkloadIdentityResponse) GetBody() *UpdateWorkloadIdentityResponseBody {
	return s.Body
}

func (s *UpdateWorkloadIdentityResponse) SetHeaders(v map[string]*string) *UpdateWorkloadIdentityResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkloadIdentityResponse) SetStatusCode(v int32) *UpdateWorkloadIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkloadIdentityResponse) SetBody(v *UpdateWorkloadIdentityResponseBody) *UpdateWorkloadIdentityResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkloadIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
