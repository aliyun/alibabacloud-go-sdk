// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticJobRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSemanticJobRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSemanticJobRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListSemanticJobRunsResponseBody) *ListSemanticJobRunsResponse
	GetBody() *ListSemanticJobRunsResponseBody
}

type ListSemanticJobRunsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSemanticJobRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSemanticJobRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobRunsResponse) GoString() string {
	return s.String()
}

func (s *ListSemanticJobRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSemanticJobRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSemanticJobRunsResponse) GetBody() *ListSemanticJobRunsResponseBody {
	return s.Body
}

func (s *ListSemanticJobRunsResponse) SetHeaders(v map[string]*string) *ListSemanticJobRunsResponse {
	s.Headers = v
	return s
}

func (s *ListSemanticJobRunsResponse) SetStatusCode(v int32) *ListSemanticJobRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSemanticJobRunsResponse) SetBody(v *ListSemanticJobRunsResponseBody) *ListSemanticJobRunsResponse {
	s.Body = v
	return s
}

func (s *ListSemanticJobRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
