// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiModelSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiModelSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListAiModelSourcesResponseBody) *ListAiModelSourcesResponse
	GetBody() *ListAiModelSourcesResponseBody
}

type ListAiModelSourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiModelSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiModelSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAiModelSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiModelSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiModelSourcesResponse) GetBody() *ListAiModelSourcesResponseBody {
	return s.Body
}

func (s *ListAiModelSourcesResponse) SetHeaders(v map[string]*string) *ListAiModelSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAiModelSourcesResponse) SetStatusCode(v int32) *ListAiModelSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiModelSourcesResponse) SetBody(v *ListAiModelSourcesResponseBody) *ListAiModelSourcesResponse {
	s.Body = v
	return s
}

func (s *ListAiModelSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
