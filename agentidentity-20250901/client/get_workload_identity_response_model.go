// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkloadIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkloadIdentityResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkloadIdentityResponseBody) *GetWorkloadIdentityResponse
	GetBody() *GetWorkloadIdentityResponseBody
}

type GetWorkloadIdentityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkloadIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkloadIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetWorkloadIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkloadIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkloadIdentityResponse) GetBody() *GetWorkloadIdentityResponseBody {
	return s.Body
}

func (s *GetWorkloadIdentityResponse) SetHeaders(v map[string]*string) *GetWorkloadIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetWorkloadIdentityResponse) SetStatusCode(v int32) *GetWorkloadIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkloadIdentityResponse) SetBody(v *GetWorkloadIdentityResponseBody) *GetWorkloadIdentityResponse {
	s.Body = v
	return s
}

func (s *GetWorkloadIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
