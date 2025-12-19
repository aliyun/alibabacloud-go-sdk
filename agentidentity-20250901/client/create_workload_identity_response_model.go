// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkloadIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkloadIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkloadIdentityResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkloadIdentityResponseBody) *CreateWorkloadIdentityResponse
	GetBody() *CreateWorkloadIdentityResponseBody
}

type CreateWorkloadIdentityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkloadIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkloadIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkloadIdentityResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkloadIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkloadIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkloadIdentityResponse) GetBody() *CreateWorkloadIdentityResponseBody {
	return s.Body
}

func (s *CreateWorkloadIdentityResponse) SetHeaders(v map[string]*string) *CreateWorkloadIdentityResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkloadIdentityResponse) SetStatusCode(v int32) *CreateWorkloadIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkloadIdentityResponse) SetBody(v *CreateWorkloadIdentityResponseBody) *CreateWorkloadIdentityResponse {
	s.Body = v
	return s
}

func (s *CreateWorkloadIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
