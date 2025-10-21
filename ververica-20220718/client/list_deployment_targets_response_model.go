// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeploymentTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeploymentTargetsResponse
	GetStatusCode() *int32
	SetBody(v *ListDeploymentTargetsResponseBody) *ListDeploymentTargetsResponse
	GetBody() *ListDeploymentTargetsResponseBody
}

type ListDeploymentTargetsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeploymentTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeploymentTargetsResponse) GetBody() *ListDeploymentTargetsResponseBody {
	return s.Body
}

func (s *ListDeploymentTargetsResponse) SetHeaders(v map[string]*string) *ListDeploymentTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentTargetsResponse) SetStatusCode(v int32) *ListDeploymentTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentTargetsResponse) SetBody(v *ListDeploymentTargetsResponseBody) *ListDeploymentTargetsResponse {
	s.Body = v
	return s
}

func (s *ListDeploymentTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
