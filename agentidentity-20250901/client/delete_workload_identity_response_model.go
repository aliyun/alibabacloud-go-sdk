// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkloadIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkloadIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkloadIdentityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkloadIdentityResponseBody) *DeleteWorkloadIdentityResponse
	GetBody() *DeleteWorkloadIdentityResponseBody
}

type DeleteWorkloadIdentityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkloadIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkloadIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkloadIdentityResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkloadIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkloadIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkloadIdentityResponse) GetBody() *DeleteWorkloadIdentityResponseBody {
	return s.Body
}

func (s *DeleteWorkloadIdentityResponse) SetHeaders(v map[string]*string) *DeleteWorkloadIdentityResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkloadIdentityResponse) SetStatusCode(v int32) *DeleteWorkloadIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkloadIdentityResponse) SetBody(v *DeleteWorkloadIdentityResponseBody) *DeleteWorkloadIdentityResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkloadIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
