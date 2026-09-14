// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectPipelineRunItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrossProjectPipelineRunItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrossProjectPipelineRunItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListCrossProjectPipelineRunItemsResponseBody) *ListCrossProjectPipelineRunItemsResponse
	GetBody() *ListCrossProjectPipelineRunItemsResponseBody
}

type ListCrossProjectPipelineRunItemsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrossProjectPipelineRunItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrossProjectPipelineRunItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunItemsResponse) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrossProjectPipelineRunItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrossProjectPipelineRunItemsResponse) GetBody() *ListCrossProjectPipelineRunItemsResponseBody {
	return s.Body
}

func (s *ListCrossProjectPipelineRunItemsResponse) SetHeaders(v map[string]*string) *ListCrossProjectPipelineRunItemsResponse {
	s.Headers = v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponse) SetStatusCode(v int32) *ListCrossProjectPipelineRunItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponse) SetBody(v *ListCrossProjectPipelineRunItemsResponseBody) *ListCrossProjectPipelineRunItemsResponse {
	s.Body = v
	return s
}

func (s *ListCrossProjectPipelineRunItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
