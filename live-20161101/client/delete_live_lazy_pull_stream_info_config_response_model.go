// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveLazyPullStreamInfoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveLazyPullStreamInfoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveLazyPullStreamInfoConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveLazyPullStreamInfoConfigResponseBody) *DeleteLiveLazyPullStreamInfoConfigResponse
	GetBody() *DeleteLiveLazyPullStreamInfoConfigResponseBody
}

type DeleteLiveLazyPullStreamInfoConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveLazyPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveLazyPullStreamInfoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveLazyPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveLazyPullStreamInfoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveLazyPullStreamInfoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveLazyPullStreamInfoConfigResponse) GetBody() *DeleteLiveLazyPullStreamInfoConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveLazyPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveLazyPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveLazyPullStreamInfoConfigResponse) SetStatusCode(v int32) *DeleteLiveLazyPullStreamInfoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveLazyPullStreamInfoConfigResponse) SetBody(v *DeleteLiveLazyPullStreamInfoConfigResponseBody) *DeleteLiveLazyPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveLazyPullStreamInfoConfigResponse) Validate() error {
	return dara.Validate(s)
}
