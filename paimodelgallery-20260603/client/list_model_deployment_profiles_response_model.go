// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelDeploymentProfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelDeploymentProfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelDeploymentProfilesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelDeploymentProfilesResponseBody) *ListModelDeploymentProfilesResponse
	GetBody() *ListModelDeploymentProfilesResponseBody
}

type ListModelDeploymentProfilesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelDeploymentProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelDeploymentProfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelDeploymentProfilesResponse) GoString() string {
	return s.String()
}

func (s *ListModelDeploymentProfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelDeploymentProfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelDeploymentProfilesResponse) GetBody() *ListModelDeploymentProfilesResponseBody {
	return s.Body
}

func (s *ListModelDeploymentProfilesResponse) SetHeaders(v map[string]*string) *ListModelDeploymentProfilesResponse {
	s.Headers = v
	return s
}

func (s *ListModelDeploymentProfilesResponse) SetStatusCode(v int32) *ListModelDeploymentProfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelDeploymentProfilesResponse) SetBody(v *ListModelDeploymentProfilesResponseBody) *ListModelDeploymentProfilesResponse {
	s.Body = v
	return s
}

func (s *ListModelDeploymentProfilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
