// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectDeploymentCandidatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrossProjectDeploymentCandidatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrossProjectDeploymentCandidatesResponse
	GetStatusCode() *int32
	SetBody(v *ListCrossProjectDeploymentCandidatesResponseBody) *ListCrossProjectDeploymentCandidatesResponse
	GetBody() *ListCrossProjectDeploymentCandidatesResponseBody
}

type ListCrossProjectDeploymentCandidatesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrossProjectDeploymentCandidatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrossProjectDeploymentCandidatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentCandidatesResponse) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentCandidatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrossProjectDeploymentCandidatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrossProjectDeploymentCandidatesResponse) GetBody() *ListCrossProjectDeploymentCandidatesResponseBody {
	return s.Body
}

func (s *ListCrossProjectDeploymentCandidatesResponse) SetHeaders(v map[string]*string) *ListCrossProjectDeploymentCandidatesResponse {
	s.Headers = v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponse) SetStatusCode(v int32) *ListCrossProjectDeploymentCandidatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponse) SetBody(v *ListCrossProjectDeploymentCandidatesResponseBody) *ListCrossProjectDeploymentCandidatesResponse {
	s.Body = v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
