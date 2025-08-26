// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodcastTaskSubmitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PodcastTaskSubmitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PodcastTaskSubmitResponse
	GetStatusCode() *int32
	SetBody(v *PodcastTaskSubmitResponseBody) *PodcastTaskSubmitResponse
	GetBody() *PodcastTaskSubmitResponseBody
}

type PodcastTaskSubmitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PodcastTaskSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PodcastTaskSubmitResponse) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskSubmitResponse) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PodcastTaskSubmitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PodcastTaskSubmitResponse) GetBody() *PodcastTaskSubmitResponseBody {
	return s.Body
}

func (s *PodcastTaskSubmitResponse) SetHeaders(v map[string]*string) *PodcastTaskSubmitResponse {
	s.Headers = v
	return s
}

func (s *PodcastTaskSubmitResponse) SetStatusCode(v int32) *PodcastTaskSubmitResponse {
	s.StatusCode = &v
	return s
}

func (s *PodcastTaskSubmitResponse) SetBody(v *PodcastTaskSubmitResponseBody) *PodcastTaskSubmitResponse {
	s.Body = v
	return s
}

func (s *PodcastTaskSubmitResponse) Validate() error {
	return dara.Validate(s)
}
