// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineGroupsResponseBody) *ListPipelineGroupsResponse
	GetBody() *ListPipelineGroupsResponseBody
}

type ListPipelineGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineGroupsResponse) GetBody() *ListPipelineGroupsResponseBody {
	return s.Body
}

func (s *ListPipelineGroupsResponse) SetHeaders(v map[string]*string) *ListPipelineGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineGroupsResponse) SetStatusCode(v int32) *ListPipelineGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineGroupsResponse) SetBody(v *ListPipelineGroupsResponseBody) *ListPipelineGroupsResponse {
	s.Body = v
	return s
}

func (s *ListPipelineGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
