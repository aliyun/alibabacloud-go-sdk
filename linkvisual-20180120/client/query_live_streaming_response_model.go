// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLiveStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLiveStreamingResponse
	GetStatusCode() *int32
	SetBody(v *QueryLiveStreamingResponseBody) *QueryLiveStreamingResponse
	GetBody() *QueryLiveStreamingResponseBody
}

type QueryLiveStreamingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveStreamingResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLiveStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLiveStreamingResponse) GetBody() *QueryLiveStreamingResponseBody {
	return s.Body
}

func (s *QueryLiveStreamingResponse) SetHeaders(v map[string]*string) *QueryLiveStreamingResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveStreamingResponse) SetStatusCode(v int32) *QueryLiveStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveStreamingResponse) SetBody(v *QueryLiveStreamingResponseBody) *QueryLiveStreamingResponse {
	s.Body = v
	return s
}

func (s *QueryLiveStreamingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
