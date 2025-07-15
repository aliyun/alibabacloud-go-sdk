// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePullStreamInfoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivePullStreamInfoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivePullStreamInfoConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivePullStreamInfoConfigResponseBody) *DeleteLivePullStreamInfoConfigResponse
	GetBody() *DeleteLivePullStreamInfoConfigResponseBody
}

type DeleteLivePullStreamInfoConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivePullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivePullStreamInfoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivePullStreamInfoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivePullStreamInfoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivePullStreamInfoConfigResponse) GetBody() *DeleteLivePullStreamInfoConfigResponseBody {
	return s.Body
}

func (s *DeleteLivePullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *DeleteLivePullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivePullStreamInfoConfigResponse) SetStatusCode(v int32) *DeleteLivePullStreamInfoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigResponse) SetBody(v *DeleteLivePullStreamInfoConfigResponseBody) *DeleteLivePullStreamInfoConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLivePullStreamInfoConfigResponse) Validate() error {
	return dara.Validate(s)
}
