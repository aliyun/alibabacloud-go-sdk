// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamPreloadTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveStreamPreloadTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveStreamPreloadTasksResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveStreamPreloadTasksResponseBody) *SetLiveStreamPreloadTasksResponse
	GetBody() *SetLiveStreamPreloadTasksResponseBody
}

type SetLiveStreamPreloadTasksResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveStreamPreloadTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveStreamPreloadTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamPreloadTasksResponse) GoString() string {
	return s.String()
}

func (s *SetLiveStreamPreloadTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveStreamPreloadTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveStreamPreloadTasksResponse) GetBody() *SetLiveStreamPreloadTasksResponseBody {
	return s.Body
}

func (s *SetLiveStreamPreloadTasksResponse) SetHeaders(v map[string]*string) *SetLiveStreamPreloadTasksResponse {
	s.Headers = v
	return s
}

func (s *SetLiveStreamPreloadTasksResponse) SetStatusCode(v int32) *SetLiveStreamPreloadTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponse) SetBody(v *SetLiveStreamPreloadTasksResponseBody) *SetLiveStreamPreloadTasksResponse {
	s.Body = v
	return s
}

func (s *SetLiveStreamPreloadTasksResponse) Validate() error {
	return dara.Validate(s)
}
