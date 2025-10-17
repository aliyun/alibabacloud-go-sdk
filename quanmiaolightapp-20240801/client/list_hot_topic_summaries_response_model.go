// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotTopicSummariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotTopicSummariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotTopicSummariesResponse
	GetStatusCode() *int32
	SetBody(v *ListHotTopicSummariesResponseBody) *ListHotTopicSummariesResponse
	GetBody() *ListHotTopicSummariesResponseBody
}

type ListHotTopicSummariesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotTopicSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotTopicSummariesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicSummariesResponse) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotTopicSummariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotTopicSummariesResponse) GetBody() *ListHotTopicSummariesResponseBody {
	return s.Body
}

func (s *ListHotTopicSummariesResponse) SetHeaders(v map[string]*string) *ListHotTopicSummariesResponse {
	s.Headers = v
	return s
}

func (s *ListHotTopicSummariesResponse) SetStatusCode(v int32) *ListHotTopicSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotTopicSummariesResponse) SetBody(v *ListHotTopicSummariesResponseBody) *ListHotTopicSummariesResponse {
	s.Body = v
	return s
}

func (s *ListHotTopicSummariesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
