// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentDraftsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeploymentDraftsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeploymentDraftsResponse
	GetStatusCode() *int32
	SetBody(v *ListDeploymentDraftsResponseBody) *ListDeploymentDraftsResponse
	GetBody() *ListDeploymentDraftsResponseBody
}

type ListDeploymentDraftsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentDraftsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentDraftsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentDraftsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentDraftsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeploymentDraftsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeploymentDraftsResponse) GetBody() *ListDeploymentDraftsResponseBody {
	return s.Body
}

func (s *ListDeploymentDraftsResponse) SetHeaders(v map[string]*string) *ListDeploymentDraftsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentDraftsResponse) SetStatusCode(v int32) *ListDeploymentDraftsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentDraftsResponse) SetBody(v *ListDeploymentDraftsResponseBody) *ListDeploymentDraftsResponse {
	s.Body = v
	return s
}

func (s *ListDeploymentDraftsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
