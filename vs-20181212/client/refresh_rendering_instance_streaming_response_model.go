// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshRenderingInstanceStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshRenderingInstanceStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshRenderingInstanceStreamingResponse
	GetStatusCode() *int32
	SetBody(v *RefreshRenderingInstanceStreamingResponseBody) *RefreshRenderingInstanceStreamingResponse
	GetBody() *RefreshRenderingInstanceStreamingResponseBody
}

type RefreshRenderingInstanceStreamingResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshRenderingInstanceStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshRenderingInstanceStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshRenderingInstanceStreamingResponse) GoString() string {
	return s.String()
}

func (s *RefreshRenderingInstanceStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshRenderingInstanceStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshRenderingInstanceStreamingResponse) GetBody() *RefreshRenderingInstanceStreamingResponseBody {
	return s.Body
}

func (s *RefreshRenderingInstanceStreamingResponse) SetHeaders(v map[string]*string) *RefreshRenderingInstanceStreamingResponse {
	s.Headers = v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponse) SetStatusCode(v int32) *RefreshRenderingInstanceStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponse) SetBody(v *RefreshRenderingInstanceStreamingResponseBody) *RefreshRenderingInstanceStreamingResponse {
	s.Body = v
	return s
}

func (s *RefreshRenderingInstanceStreamingResponse) Validate() error {
	return dara.Validate(s)
}
