// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodcastTaskResultQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PodcastTaskResultQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PodcastTaskResultQueryResponse
	GetStatusCode() *int32
	SetBody(v *PodcastTaskResultQueryResponseBody) *PodcastTaskResultQueryResponse
	GetBody() *PodcastTaskResultQueryResponseBody
}

type PodcastTaskResultQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PodcastTaskResultQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PodcastTaskResultQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskResultQueryResponse) GoString() string {
	return s.String()
}

func (s *PodcastTaskResultQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PodcastTaskResultQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PodcastTaskResultQueryResponse) GetBody() *PodcastTaskResultQueryResponseBody {
	return s.Body
}

func (s *PodcastTaskResultQueryResponse) SetHeaders(v map[string]*string) *PodcastTaskResultQueryResponse {
	s.Headers = v
	return s
}

func (s *PodcastTaskResultQueryResponse) SetStatusCode(v int32) *PodcastTaskResultQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *PodcastTaskResultQueryResponse) SetBody(v *PodcastTaskResultQueryResponseBody) *PodcastTaskResultQueryResponse {
	s.Body = v
	return s
}

func (s *PodcastTaskResultQueryResponse) Validate() error {
	return dara.Validate(s)
}
