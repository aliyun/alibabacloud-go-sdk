// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeploymentJobCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeploymentJobCertResponse
	GetStatusCode() *int32
	SetBody(v *ListDeploymentJobCertResponseBody) *ListDeploymentJobCertResponse
	GetBody() *ListDeploymentJobCertResponseBody
}

type ListDeploymentJobCertResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentJobCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentJobCertResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobCertResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeploymentJobCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeploymentJobCertResponse) GetBody() *ListDeploymentJobCertResponseBody {
	return s.Body
}

func (s *ListDeploymentJobCertResponse) SetHeaders(v map[string]*string) *ListDeploymentJobCertResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentJobCertResponse) SetStatusCode(v int32) *ListDeploymentJobCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentJobCertResponse) SetBody(v *ListDeploymentJobCertResponseBody) *ListDeploymentJobCertResponse {
	s.Body = v
	return s
}

func (s *ListDeploymentJobCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
