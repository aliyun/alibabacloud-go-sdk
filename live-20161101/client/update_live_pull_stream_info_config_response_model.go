// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePullStreamInfoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLivePullStreamInfoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLivePullStreamInfoConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLivePullStreamInfoConfigResponseBody) *UpdateLivePullStreamInfoConfigResponse
	GetBody() *UpdateLivePullStreamInfoConfigResponseBody
}

type UpdateLivePullStreamInfoConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivePullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivePullStreamInfoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivePullStreamInfoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLivePullStreamInfoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLivePullStreamInfoConfigResponse) GetBody() *UpdateLivePullStreamInfoConfigResponseBody {
	return s.Body
}

func (s *UpdateLivePullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *UpdateLivePullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivePullStreamInfoConfigResponse) SetStatusCode(v int32) *UpdateLivePullStreamInfoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigResponse) SetBody(v *UpdateLivePullStreamInfoConfigResponseBody) *UpdateLivePullStreamInfoConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLivePullStreamInfoConfigResponse) Validate() error {
	return dara.Validate(s)
}
