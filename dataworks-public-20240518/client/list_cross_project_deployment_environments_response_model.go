// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectDeploymentEnvironmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrossProjectDeploymentEnvironmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrossProjectDeploymentEnvironmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListCrossProjectDeploymentEnvironmentsResponseBody) *ListCrossProjectDeploymentEnvironmentsResponse
	GetBody() *ListCrossProjectDeploymentEnvironmentsResponseBody
}

type ListCrossProjectDeploymentEnvironmentsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrossProjectDeploymentEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrossProjectDeploymentEnvironmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentEnvironmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrossProjectDeploymentEnvironmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrossProjectDeploymentEnvironmentsResponse) GetBody() *ListCrossProjectDeploymentEnvironmentsResponseBody {
	return s.Body
}

func (s *ListCrossProjectDeploymentEnvironmentsResponse) SetHeaders(v map[string]*string) *ListCrossProjectDeploymentEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponse) SetStatusCode(v int32) *ListCrossProjectDeploymentEnvironmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponse) SetBody(v *ListCrossProjectDeploymentEnvironmentsResponseBody) *ListCrossProjectDeploymentEnvironmentsResponse {
	s.Body = v
	return s
}

func (s *ListCrossProjectDeploymentEnvironmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
