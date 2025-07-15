// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveLazyPullStreamInfoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveLazyPullStreamInfoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveLazyPullStreamInfoConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveLazyPullStreamInfoConfigResponseBody) *SetLiveLazyPullStreamInfoConfigResponse
	GetBody() *SetLiveLazyPullStreamInfoConfigResponseBody
}

type SetLiveLazyPullStreamInfoConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveLazyPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveLazyPullStreamInfoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveLazyPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveLazyPullStreamInfoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveLazyPullStreamInfoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveLazyPullStreamInfoConfigResponse) GetBody() *SetLiveLazyPullStreamInfoConfigResponseBody {
	return s.Body
}

func (s *SetLiveLazyPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *SetLiveLazyPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigResponse) SetStatusCode(v int32) *SetLiveLazyPullStreamInfoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigResponse) SetBody(v *SetLiveLazyPullStreamInfoConfigResponseBody) *SetLiveLazyPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigResponse) Validate() error {
	return dara.Validate(s)
}
