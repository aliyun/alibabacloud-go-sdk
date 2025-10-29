// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamPreloadTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamPreloadTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamPreloadTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamPreloadTasksResponseBody) *DescribeLiveStreamPreloadTasksResponse
	GetBody() *DescribeLiveStreamPreloadTasksResponseBody
}

type DescribeLiveStreamPreloadTasksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamPreloadTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamPreloadTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPreloadTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPreloadTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamPreloadTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamPreloadTasksResponse) GetBody() *DescribeLiveStreamPreloadTasksResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamPreloadTasksResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamPreloadTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponse) SetStatusCode(v int32) *DescribeLiveStreamPreloadTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponse) SetBody(v *DescribeLiveStreamPreloadTasksResponseBody) *DescribeLiveStreamPreloadTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamPreloadTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
