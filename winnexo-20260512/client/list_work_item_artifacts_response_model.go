// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemArtifactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkItemArtifactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkItemArtifactsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkItemArtifactsResponseBody) *ListWorkItemArtifactsResponse
	GetBody() *ListWorkItemArtifactsResponseBody
}

type ListWorkItemArtifactsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkItemArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkItemArtifactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemArtifactsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkItemArtifactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkItemArtifactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkItemArtifactsResponse) GetBody() *ListWorkItemArtifactsResponseBody {
	return s.Body
}

func (s *ListWorkItemArtifactsResponse) SetHeaders(v map[string]*string) *ListWorkItemArtifactsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkItemArtifactsResponse) SetStatusCode(v int32) *ListWorkItemArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkItemArtifactsResponse) SetBody(v *ListWorkItemArtifactsResponseBody) *ListWorkItemArtifactsResponse {
	s.Body = v
	return s
}

func (s *ListWorkItemArtifactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
