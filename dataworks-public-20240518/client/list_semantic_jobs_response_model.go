// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSemanticJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSemanticJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListSemanticJobsResponseBody) *ListSemanticJobsResponse
	GetBody() *ListSemanticJobsResponseBody
}

type ListSemanticJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSemanticJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSemanticJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSemanticJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSemanticJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSemanticJobsResponse) GetBody() *ListSemanticJobsResponseBody {
	return s.Body
}

func (s *ListSemanticJobsResponse) SetHeaders(v map[string]*string) *ListSemanticJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSemanticJobsResponse) SetStatusCode(v int32) *ListSemanticJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSemanticJobsResponse) SetBody(v *ListSemanticJobsResponseBody) *ListSemanticJobsResponse {
	s.Body = v
	return s
}

func (s *ListSemanticJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
